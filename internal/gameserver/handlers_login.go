package gameserver

import (
	"strings"

	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/packet"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func (s *Server) registerLoginHandlers() {
	pushGenderInfo := func(ctx packet.Context, p *player.Player) {
		ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{
			Retcode:               0,
			CurAvatarPath:         pb.MultiPathAvatarType(s.currentTrailblazerAvatarID(p)),
			CurAvatarPathInfoList: s.multiPathInfosForGender(p),
		})
	}
	pushLineupInfo := func(ctx packet.Context, p *player.Player) {
		if p == nil {
			return
		}
		lu := lineup.ToProto(p, s.scene.PlaneID, p.CurrentLineupIndex)
		ctx.Send(packet.SyncLineupNotify, &pb.SyncLineupNotify{
			ReasonList: nil,
			Lineup:     lu,
		})
		ctx.Send(packet.GetCurLineupDataScRsp, &pb.GetCurLineupDataScRsp{Retcode: 0, Lineup: lu})
		ctx.Send(packet.GetAllLineupDataScRsp, &pb.GetAllLineupDataScRsp{
			Retcode:    0,
			CurIndex:   p.CurrentLineupIndex,
			LineupList: lineup.AllToProto(p, s.scene.PlaneID),
		})
	}

	s.reg.Register(packet.GetSecretKeyInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetSecretKeyInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetSecretKeyInfoScRsp, &pb.GetSecretKeyInfoScRsp{Retcode: 0})
		},
	})

	s.reg.Register(packet.PlayerHeartBeatCsReq, packet.Handler{
		Name: packet.Name(packet.PlayerHeartBeatCsReq),
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.PlayerHeartBeatCsReq
			_ = proto.Unmarshal(payload, &req)
			nowMs := uint64(ctx.Now().UnixMilli())
			ctx.Send(packet.PlayerHeartBeatScRsp, &pb.PlayerHeartBeatScRsp{
				ClientTimeMs: req.ClientTimeMs,
				ServerTimeMs: nowMs,
				Retcode:      0,
			})
		},
	})

	s.reg.Register(packet.GetBasicInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetBasicInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.db != nil {
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
			}
			ctx.Send(packet.GetBasicInfoScRsp, &pb.GetBasicInfoScRsp{
				Retcode:                 0,
				Gender:                  p.Gender,
				PlayerSettingInfo:       &pb.PlayerSettingInfo{DisplayDiary: true},
				NextRecoverTime:         p.NextRecoverTimeMs,
				GameplayBirthday:        p.GameplayBirthday,
				IsGenderSet:             p.Gender != 0,
				LastSetNicknameTime:     p.LastSetNicknameTimeMs,
				CurDay:                  1,
				ExchangeTimes:           0,
				WeekCocoonFinishedCount: 0,
			})
			pushGenderInfo(ctx, p)
		},
	})

	s.reg.Register(packet.SetPlayerInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.SetPlayerInfoCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetPlayerInfoCsReq
			_ = proto.Unmarshal(payload, &req)

			if s.db == nil {
				ctx.Send(packet.SetPlayerInfoScRsp, &pb.SetPlayerInfoScRsp{Retcode: 0, IsModify: req.IsModify})
				return
			}

			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
				p = loaded
			}
			if req.Nickname != "" {
				p.Nickname = req.Nickname
				p.LastSetNicknameTimeMs = ctx.Now().Unix()
			}
			genderChanged := false
			if req.Gender != pb.Gender_GenderNone {
				genderChanged = true
				p.Gender = uint32(req.Gender)
				if p.CurrentMultiPath == nil {
					p.CurrentMultiPath = map[uint32]uint32{}
				}
				if p.Gender == player.GenderMan {
					if p.CurrentMultiPath[8001] == 0 {
						p.CurrentMultiPath[8001] = 8005
					}
				} else if p.Gender == player.GenderWoman {
					if p.CurrentMultiPath[8002] == 0 {
						p.CurrentMultiPath[8002] = 8006
					}
				}
				_ = s.ensurePlayerDefaults(p)
			}
			p.Touch()
			_ = s.db.SavePlayer(p)

			ctx.Send(packet.SetPlayerInfoScRsp, &pb.SetPlayerInfoScRsp{
				Retcode:               0,
				IsModify:              req.IsModify,
				CHADHDHLDFJ:           ctx.Now().Unix(),
				CurAvatarPathInfoList: s.multiPathInfosForGender(p),
				CurAvatarPath:         pb.MultiPathAvatarType(s.currentTrailblazerAvatarID(p)),
			})

			// Refresh avatar list for the client when gender changes.
			if genderChanged {
				ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
					AvatarSync:              &pb.AvatarSync{AvatarList: s.buildAvatarList(p)},
					MultiPathAvatarInfoList: s.buildMultiPathAvatarInfoList(p),
					SyncStatus:              &pb.SyncStatus{},
				})
				pushLineupInfo(ctx, p)
			}
		},
	})

	s.reg.Register(packet.SetNicknameCsReq, packet.Handler{
		Name:     packet.Name(packet.SetNicknameCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetNicknameCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db != nil {
				p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if req.Nickname != "" {
					p.Nickname = req.Nickname
					p.LastSetNicknameTimeMs = ctx.Now().Unix()
					p.Touch()
					_ = s.db.SavePlayer(p)
				}
			}
			ctx.Send(packet.SetNicknameScRsp, &pb.SetNicknameScRsp{Retcode: 0, IsModify: req.IsModify, CHADHDHLDFJ: ctx.Now().Unix()})
		},
	})

	s.reg.Register(packet.SetGenderCsReq, packet.Handler{
		Name:     packet.Name(packet.SetGenderCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetGenderCsReq
			_ = proto.Unmarshal(payload, &req)
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.db != nil {
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
			}

			if req.Gender == pb.Gender_GenderNone {
				ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{Retcode: 0, CurAvatarPath: pb.MultiPathAvatarType_MultiPathAvatarTypeNone})
				return
			}

			p.Gender = uint32(req.Gender)
			if p.CurrentMultiPath == nil {
				p.CurrentMultiPath = map[uint32]uint32{}
			}
			if p.Gender == player.GenderMan {
				if p.CurrentMultiPath[8001] == 0 {
					p.CurrentMultiPath[8001] = 8005
				}
			} else if p.Gender == player.GenderWoman {
				if p.CurrentMultiPath[8002] == 0 {
					p.CurrentMultiPath[8002] = 8006
				}
			}
			_ = s.ensurePlayerDefaults(p)
			p.Touch()
			if s.db != nil {
				_ = s.db.SavePlayer(p)
			}

			ctx.Send(packet.SetGenderScRsp, &pb.SetGenderScRsp{
				Retcode:               0,
				CurAvatarPath:         pb.MultiPathAvatarType(s.currentTrailblazerAvatarID(p)),
				CurAvatarPathInfoList: s.multiPathInfosForGender(p),
			})
			ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
				AvatarSync:              &pb.AvatarSync{AvatarList: s.buildAvatarList(p)},
				MultiPathAvatarInfoList: s.buildMultiPathAvatarInfoList(p),
				SyncStatus:              &pb.SyncStatus{},
			})
			pushLineupInfo(ctx, p)
		},
	})

	s.reg.Register(packet.PlayerForceSyncGameStateFinishCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerForceSyncGameStateFinishCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.PlayerForceSyncGameStateFinishScRsp, &pb.PlayerForceSyncGameStateFinishScRsp{Retcode: 0})
		},
	})

	s.reg.Register(packet.SyncClientResVersionCsReq, packet.Handler{
		Name:     packet.Name(packet.SyncClientResVersionCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SyncClientResVersionCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.SyncClientResVersionScRsp, &pb.SyncClientResVersionScRsp{
				ClientResVersion: req.ClientResVersion,
				Retcode:          0,
			})
		},
	})

	s.reg.Register(packet.GetVideoVersionKeyCsReq, packet.Handler{
		Name:     packet.Name(packet.GetVideoVersionKeyCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetVideoVersionKeyScRsp, &pb.GetVideoVersionKeyScRsp{Retcode: 0})
		},
	})

	s.reg.Register(packet.GetAuthkeyCsReq, packet.Handler{
		Name:     packet.Name(packet.GetAuthkeyCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.GetAuthkeyCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.GetAuthkeyScRsp, &pb.GetAuthkeyScRsp{
				AuthkeyVer: req.AuthkeyVer,
				AuthAppid:  req.AuthAppid,
				SignType:   req.SignType,
				Authkey:    "dummy_authkey",
				Retcode:    0,
			})
		},
	})

	s.reg.Register(packet.PlayerGetTokenCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerGetTokenCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			if ctx.Session.State != packet.StateWaitingForToken {
				return
			}
			var req pb.PlayerGetTokenCsReq
			if err := proto.Unmarshal(payload, &req); err != nil {
				return
			}

			s.mu.RLock()
			db := s.db
			s.mu.RUnlock()
			if db == nil {
				return
			}
			acc, ok := db.GetAccountByUID(req.AccountUid)
			if !ok || acc == nil || acc.ComboToken == "" || acc.ComboToken != req.Token {
				return
			}
			if acc.ReservedPlayerUID == 0 {
				acc.ReservedPlayerUID = int(s.nextUID.Add(1))
				_ = db.SaveAccounts()
			}

			ctx.Session.AccountUID = acc.UID
			ctx.Session.UID = uint32(acc.ReservedPlayerUID)
			ctx.Session.State = packet.StateWaitingForLogin

			if p, _, err := db.EnsurePlayerForAccount(acc.UID, ctx.Session.UID); err == nil && p != nil {
				_ = db.SavePlayer(p)
			}

			ctx.Send(packet.PlayerGetTokenScRsp, &pb.PlayerGetTokenScRsp{
				Uid:       ctx.Session.UID,
				BlackInfo: &pb.BlackInfo{},
			})
		},
	})

	s.reg.Register(packet.PlayerLoginCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerLoginCsReq),
		MinState: packet.StateWaitingForLogin,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			if ctx.Session.State != packet.StateWaitingForLogin {
				return
			}
			nowMs := uint64(ctx.Now().UnixMilli())
			basic := &pb.PlayerBasicInfo{Nickname: "HyacineLover", Level: 70, WorldLevel: 6, Stamina: 300}
			stamina := uint32(300)
			reserve := uint32(0)
			nextRecover := int64(0)
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			if s.db != nil {
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
					basic.Nickname = loaded.Nickname
					basic.Level = loaded.Level
					basic.Exp = loaded.Exp
					basic.WorldLevel = loaded.WorldLevel
					basic.Stamina = loaded.Stamina
					stamina = loaded.Stamina
					reserve = loaded.ReserveStamina
					nextRecover = loaded.NextRecoverTimeMs
				}
			}

			ctx.Send(packet.PlayerLoginScRsp, &pb.PlayerLoginScRsp{
				BasicInfo:         basic,
				CurTimezone:       8,
				ServerTimestampMs: nowMs,
				Stamina:           stamina,
			})
			ctx.Send(packet.StaminaInfoScNotify, &pb.StaminaInfoScNotify{
				NextRecoverTime: nextRecover,
				Stamina:         stamina,
				ReserveStamina:  reserve,
			})

			// Sync player/avatars/lineup-related state after login.
			_ = s.ensurePlayerDefaults(p)
			ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
				BasicInfo: basic,
				BasicModuleSync: &pb.BasicModuleSync{
					Stamina: stamina,
				},
				EquipmentList: s.buildEquipmentList(p),
				RelicList:     s.buildRelicList(p),
				AvatarSync: &pb.AvatarSync{
					AvatarList: s.buildAvatarList(p),
				},
				MultiPathAvatarInfoList: s.buildMultiPathAvatarInfoList(p),
				PlayerboardModuleSync: &pb.PlayerBoardModuleSync{
					Signature:            "HyacinePS",
					HeadFrameInfo:        &pb.HeadFrameInfo{},
					UnlockedHeadIconList: s.unlockedHeadIconList(),
				},
				PlayerOutfitData: &pb.PlayerOutfitInfo{
					PlayerOutfitList: s.unlockedOutfitIDs(),
				},
				SyncStatus: &pb.SyncStatus{},
			})

			s.mu.RLock()
			extra := append([]uint16(nil), s.opts.LoginExtraSendCmdIDs...)
			s.mu.RUnlock()
			for _, id := range extra {
				ctx.SendEmpty(id)
			}

			ctx.Session.State = packet.StateActive
		},
	})

	s.reg.Register(packet.PlayerLoginFinishCsReq, packet.Handler{
		Name:     packet.Name(packet.PlayerLoginFinishCsReq),
		MinState: packet.StateWaitingForToken,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.PlayerLoginFinishScRsp, &pb.PlayerLoginFinishScRsp{Retcode: 0})
		},
	})

	// Keep a small startup hint in logs if a config is missing.
	if s.db == nil || strings.TrimSpace(s.opts.AccountPath) == "" {
		// keep silent; logger setup may not be ready yet.
	}
}
