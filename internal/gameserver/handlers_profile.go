package gameserver

import (
	"google.golang.org/protobuf/proto"

	"hyacine-server/internal/gameserver/packet"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) registerProfileHandlers() {
	getPlayer := func(uid uint32) *player.Player {
		if s.db == nil {
			return nil
		}
		p, ok := s.db.GetPlayerByUID(uid)
		if !ok || p == nil {
			return nil
		}
		return p
	}

	savePlayer := func(p *player.Player) {
		if p == nil || s.db == nil {
			return
		}
		p.Touch()
		_ = s.db.SavePlayer(p)
	}

	s.reg.Register(packet.GetPhoneDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetPhoneDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := getPlayer(ctx.Session.UID)
			if p == nil {
				ctx.Send(packet.GetPhoneDataScRsp, &pb.GetPhoneDataScRsp{Retcode: 0, CurPhoneCase: 254001})
				return
			}
			ctx.Send(packet.GetPhoneDataScRsp, &pb.GetPhoneDataScRsp{
				Retcode:          0,
				CurChatBubble:    p.CurChatBubble,
				CurPhoneTheme:    p.CurPhoneTheme,
				CurPhoneCase:     p.CurPhoneCase,
				OwnedChatBubbles: p.OwnedChatBubbles,
				OwnedPhoneThemes: p.OwnedPhoneThemes,
				OwnedPhoneCases:  p.OwnedPhoneCases,
			})
		},
	})

	s.reg.Register(packet.SelectPhoneThemeCsReq, packet.Handler{
		Name:     packet.Name(packet.SelectPhoneThemeCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SelectPhoneThemeCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.CurPhoneTheme = req.GetThemeId()
				savePlayer(p)
			}
			ctx.Send(packet.SelectPhoneThemeScRsp, &pb.SelectPhoneThemeScRsp{Retcode: 0, CurPhoneTheme: req.GetThemeId()})
		},
	})

	s.reg.Register(packet.SelectChatBubbleCsReq, packet.Handler{
		Name:     packet.Name(packet.SelectChatBubbleCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SelectChatBubbleCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.CurChatBubble = req.GetBubbleId()
				savePlayer(p)
			}
			ctx.Send(packet.SelectChatBubbleScRsp, &pb.SelectChatBubbleScRsp{Retcode: 0, CurChatBubble: req.GetBubbleId()})
		},
	})

	s.reg.Register(packet.SelectPhoneCaseCsReq, packet.Handler{
		Name:     packet.Name(packet.SelectPhoneCaseCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SelectPhoneCaseCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.CurPhoneCase = req.GetPhoneCaseId()
				savePlayer(p)
			}
			ctx.Send(packet.SelectPhoneCaseScRsp, &pb.SelectPhoneCaseScRsp{Retcode: 0, CurPhoneCase: req.GetPhoneCaseId()})
		},
	})

	s.reg.Register(packet.GetPlayerBoardDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetPlayerBoardDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := getPlayer(ctx.Session.UID)
			if p == nil {
				ctx.Send(packet.GetPlayerBoardDataScRsp, &pb.GetPlayerBoardDataScRsp{Retcode: 0})
				return
			}
			unlockedHeadIcons := make([]*pb.HeadIconData, 0, len(p.UnlockedHeadIconIDs))
			for _, id := range p.UnlockedHeadIconIDs {
				unlockedHeadIcons = append(unlockedHeadIcons, &pb.HeadIconData{Id: id})
			}
			displayList := make([]*pb.DisplayAvatarData, 0, len(p.DisplayAvatars))
			for _, d := range p.DisplayAvatars {
				displayList = append(displayList, &pb.DisplayAvatarData{AvatarId: d.AvatarID, Pos: d.Pos})
			}
			ctx.Send(packet.GetPlayerBoardDataScRsp, &pb.GetPlayerBoardDataScRsp{
				Retcode:               0,
				CurrentPersonalCardId: p.CurrentPersonalCardID,
				UnlockedPersonalCardList: p.UnlockedPersonalCardIDs,
				CurrentHeadIconId:     p.CurrentHeadIconID,
				UnlockedHeadIconList:  unlockedHeadIcons,
				Signature:             p.Signature,
				HeadFrameInfo: &pb.HeadFrameInfo{
					HeadFrameExpireTime: p.HeadFrameExpireTime,
					HeadFrameItemId:     p.HeadFrameItemID,
				},
				AssistAvatarIdList: p.AssistAvatarIDs,
				DisplayAvatarVec: &pb.DisplayAvatarVec{
					IsDisplay:         true,
					DisplayAvatarList: displayList,
				},
			})
		},
	})

	s.reg.Register(packet.SetSignatureCsReq, packet.Handler{
		Name:     packet.Name(packet.SetSignatureCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetSignatureCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.Signature = req.GetSignature()
				savePlayer(p)
			}
			ctx.Send(packet.SetSignatureScRsp, &pb.SetSignatureScRsp{Retcode: 0, Signature: req.GetSignature()})
		},
	})

	s.reg.Register(packet.SetGameplayBirthdayCsReq, packet.Handler{
		Name:     packet.Name(packet.SetGameplayBirthdayCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetGameplayBirthdayCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.GameplayBirthday = req.GetBirthday()
				savePlayer(p)
			}
			ctx.Send(packet.SetGameplayBirthdayScRsp, &pb.SetGameplayBirthdayScRsp{Retcode: 0, Birthday: req.GetBirthday()})
		},
	})

	s.reg.Register(packet.SetHeadIconCsReq, packet.Handler{
		Name:     packet.Name(packet.SetHeadIconCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetHeadIconCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.CurrentHeadIconID = req.GetId()
				// Keep unlock list conservative: only ensure the selected icon exists.
				found := false
				for _, id := range p.UnlockedHeadIconIDs {
					if id == req.GetId() {
						found = true
						break
					}
				}
				if !found && req.GetId() != 0 {
					p.UnlockedHeadIconIDs = append(p.UnlockedHeadIconIDs, req.GetId())
				}
				savePlayer(p)
			}
			ctx.Send(packet.SetHeadIconScRsp, &pb.SetHeadIconScRsp{Retcode: 0, CurrentHeadIconId: req.GetId()})
		},
	})

	s.reg.Register(packet.SetPersonalCardCsReq, packet.Handler{
		Name:     packet.Name(packet.SetPersonalCardCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetPersonalCardCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				p.CurrentPersonalCardID = req.GetId()
				found := false
				for _, id := range p.UnlockedPersonalCardIDs {
					if id == req.GetId() {
						found = true
						break
					}
				}
				if !found && req.GetId() != 0 {
					p.UnlockedPersonalCardIDs = append(p.UnlockedPersonalCardIDs, req.GetId())
				}
				savePlayer(p)
			}
			ctx.Send(packet.SetPersonalCardScRsp, &pb.SetPersonalCardScRsp{Retcode: 0, CurrentPersonalCardId: req.GetId()})
		},
	})

	s.reg.Register(packet.SetAssistAvatarCsReq, packet.Handler{
		Name:     packet.Name(packet.SetAssistAvatarCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetAssistAvatarCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				ids := req.GetAvatarIdList()
				if ids == nil {
					ids = []uint32{}
				}
				p.AssistAvatarIDs = ids
				savePlayer(p)
			}
			ctx.Send(packet.SetAssistAvatarScRsp, &pb.SetAssistAvatarScRsp{
				Retcode:      0,
				AvatarId:     req.GetAvatarId(),
				AvatarIdList: req.GetAvatarIdList(),
			})
		},
	})

	s.reg.Register(packet.SetDisplayAvatarCsReq, packet.Handler{
		Name:     packet.Name(packet.SetDisplayAvatarCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetDisplayAvatarCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(ctx.Session.UID)
			if p != nil {
				list := make([]player.DisplayAvatar, 0, len(req.GetDisplayAvatarList()))
				for _, it := range req.GetDisplayAvatarList() {
					if it == nil {
						continue
					}
					list = append(list, player.DisplayAvatar{AvatarID: it.GetAvatarId(), Pos: it.GetPos()})
				}
				p.DisplayAvatars = list
				savePlayer(p)
			}
			ctx.Send(packet.SetDisplayAvatarScRsp, &pb.SetDisplayAvatarScRsp{Retcode: 0, DisplayAvatarList: req.GetDisplayAvatarList()})
		},
	})

	s.reg.Register(packet.UpdatePlayerSettingCsReq, packet.Handler{
		Name:     packet.Name(packet.UpdatePlayerSettingCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.UpdatePlayerSettingCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.UpdatePlayerSettingScRsp, &pb.UpdatePlayerSettingScRsp{Retcode: 0, PlayerSetting: req.GetPlayerSetting()})
		},
	})

	s.reg.Register(packet.GetPlayerDetailInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetPlayerDetailInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetPlayerDetailInfoCsReq{}
			_ = proto.Unmarshal(payload, req)
			p := getPlayer(req.GetUid())
			if p == nil {
				p = getPlayer(ctx.Session.UID)
			}
			if p == nil {
				ctx.Send(packet.GetPlayerDetailInfoScRsp, &pb.GetPlayerDetailInfoScRsp{Retcode: 0})
				return
			}
			ctx.Send(packet.GetPlayerDetailInfoScRsp, &pb.GetPlayerDetailInfoScRsp{
				Retcode: 0,
				DetailInfo: &pb.PlayerDetailInfo{
					Uid:        p.UID,
					Level:      p.Level,
					WorldLevel: p.WorldLevel,
					HeadIcon:   p.CurrentHeadIconID,
					HeadFrameInfo: &pb.HeadFrameInfo{
						HeadFrameExpireTime: p.HeadFrameExpireTime,
						HeadFrameItemId:     p.HeadFrameItemID,
					},
				},
			})
		},
	})
}

