package gameserver

import (
	"google.golang.org/protobuf/proto"

	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/packet"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) registerPlayerHandlers() {
	s.reg.Register(packet.GetBagCsReq, packet.Handler{
		Name:     packet.Name(packet.GetBagCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			ctx.Send(packet.GetBagScRsp, &pb.GetBagScRsp{
				Retcode:       0,
				EquipmentList: s.buildEquipmentList(p),
				RelicList:     s.buildRelicList(p),
			})
		},
	})

	s.reg.Register(packet.GetAvatarDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetAvatarDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)

			// Ensure defaults are consistent with freesr-data when available.
			_ = s.ensurePlayerDefaults(p)
			avatarList := s.buildAvatarList(p)
			curPaths := map[uint32]pb.MultiPathAvatarType{}
			if p != nil {
				curPaths[8001] = pb.MultiPathAvatarType(s.currentTrailblazerAvatarID(p))
				// March 7th is also multi-path (1001/1224) when present.
				if ids := s.ownedAvatarIDsRaw(p); len(ids) > 0 {
					hasMarch := false
					for _, id := range ids {
						if id == 1001 || id == 1224 {
							hasMarch = true
							break
						}
					}
					if hasMarch {
						march := uint32(1001)
						if p.CurrentMultiPath != nil {
							if v := p.CurrentMultiPath[1001]; v != 0 {
								march = v
							}
						}
						curPaths[1001] = pb.MultiPathAvatarType(march)
					}
				}
			}
			ctx.Send(packet.GetAvatarDataScRsp, &pb.GetAvatarDataScRsp{
				Retcode:                 0,
				IsGetAll:                true,
				AvatarList:              avatarList,
				CurAvatarPath:           curPaths,
				MultiPathAvatarInfoList: s.buildMultiPathAvatarInfoList(p),
			})
		},
	})

	sendSyncLineupNotify := func(ctx packet.Context, p *player.Player, lu *pb.LineupInfo) {
		ctx.Send(packet.SyncLineupNotify, &pb.SyncLineupNotify{
			// Keep empty to match client expectations for lineup changes.
			ReasonList: nil,
			Lineup:     lu,
		})

		// Some UI panels don't refresh on SyncLineupNotify alone; push lineup snapshots too.
		curIndex := uint32(0)
		if p != nil {
			curIndex = p.CurrentLineupIndex
		}
		ctx.Send(packet.GetCurLineupDataScRsp, &pb.GetCurLineupDataScRsp{Retcode: 0, Lineup: lu})
		ctx.Send(packet.GetAllLineupDataScRsp, &pb.GetAllLineupDataScRsp{Retcode: 0, CurIndex: curIndex, LineupList: lineup.AllToProto(p, s.scene.PlaneID)})
	}

	s.reg.Register(packet.GetCurLineupDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetCurLineupDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			ctx.Send(packet.GetCurLineupDataScRsp, &pb.GetCurLineupDataScRsp{
				Retcode: 0,
				Lineup:  lineup.ToProto(p, s.scene.PlaneID, p.CurrentLineupIndex),
			})
		},
	})

	s.reg.Register(packet.GetAllLineupDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetAllLineupDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			ctx.Send(packet.GetAllLineupDataScRsp, &pb.GetAllLineupDataScRsp{
				Retcode:    0,
				CurIndex:   p.CurrentLineupIndex,
				LineupList: lineup.AllToProto(p, s.scene.PlaneID),
			})
		},
	})

	s.reg.Register(packet.SetLineupNameCsReq, packet.Handler{
		Name:     packet.Name(packet.SetLineupNameCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetLineupNameCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.SetLineupNameScRsp, &pb.SetLineupNameScRsp{Retcode: 0, Index: req.Index, Name: req.Name})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.SetLineupNameScRsp, &pb.SetLineupNameScRsp{Retcode: 1, Index: req.Index, Name: req.Name})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[req.Index]
			if lu == nil {
				ctx.Send(packet.SetLineupNameScRsp, &pb.SetLineupNameScRsp{Retcode: 1, Index: req.Index, Name: req.Name})
				return
			}
			lu.Name = req.Name
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.SetLineupNameScRsp, &pb.SetLineupNameScRsp{Retcode: 0, Index: req.Index, Name: req.Name})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.SwitchLineupIndexCsReq, packet.Handler{
		Name:     packet.Name(packet.SwitchLineupIndexCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SwitchLineupIndexCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.SwitchLineupIndexScRsp, &pb.SwitchLineupIndexScRsp{Retcode: 0, Index: req.Index})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.SwitchLineupIndexScRsp, &pb.SwitchLineupIndexScRsp{Retcode: 1, Index: req.Index})
				return
			}
			p.CurrentLineupIndex = req.Index
			lineup.EnsureDefaults(p)
			_ = p.Lineups[req.Index]
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.SwitchLineupIndexScRsp, &pb.SwitchLineupIndexScRsp{Retcode: 0, Index: req.Index})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.ChangeLineupLeaderCsReq, packet.Handler{
		Name:     packet.Name(packet.ChangeLineupLeaderCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.ChangeLineupLeaderCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.ChangeLineupLeaderScRsp, &pb.ChangeLineupLeaderScRsp{Retcode: 0, Slot: req.Slot})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.ChangeLineupLeaderScRsp, &pb.ChangeLineupLeaderScRsp{Retcode: 1, Slot: req.Slot})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[p.CurrentLineupIndex]
			if lu == nil {
				ctx.Send(packet.ChangeLineupLeaderScRsp, &pb.ChangeLineupLeaderScRsp{Retcode: 1, Slot: req.Slot})
				return
			}
			if int(req.Slot) >= len(lu.AvatarIDs) {
				ctx.Send(packet.ChangeLineupLeaderScRsp, &pb.ChangeLineupLeaderScRsp{Retcode: 1, Slot: req.Slot})
				return
			}
			lu.LeaderSlot = req.Slot
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.ChangeLineupLeaderScRsp, &pb.ChangeLineupLeaderScRsp{Retcode: 0, Slot: req.Slot})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, p.CurrentLineupIndex))
		},
	})

	s.reg.Register(packet.ReplaceLineupCsReq, packet.Handler{
		Name:     packet.Name(packet.ReplaceLineupCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.ReplaceLineupCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.ReplaceLineupScRsp, &pb.ReplaceLineupScRsp{Retcode: 0})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.ReplaceLineupScRsp, &pb.ReplaceLineupScRsp{Retcode: 1})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[req.Index]
			if lu == nil {
				ctx.Send(packet.ReplaceLineupScRsp, &pb.ReplaceLineupScRsp{Retcode: 1})
				return
			}
			ids := make([]uint32, 0, len(req.LineupSlotList))
			for _, slot := range req.LineupSlotList {
				if slot == nil {
					continue
				}
				if slot.Id != 0 {
					ids = append(ids, slot.Id)
				}
			}
			if len(ids) == 0 {
				ids = append([]uint32(nil), lu.AvatarIDs...)
			}
			lu.AvatarIDs = ids
			lu.LeaderSlot = req.LeaderSlot
			if lu.LeaderSlot >= uint32(len(lu.AvatarIDs)) {
				lu.LeaderSlot = 0
			}
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.ReplaceLineupScRsp, &pb.ReplaceLineupScRsp{Retcode: 0})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.JoinLineupCsReq, packet.Handler{
		Name:     packet.Name(packet.JoinLineupCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.JoinLineupCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.JoinLineupScRsp, &pb.JoinLineupScRsp{Retcode: 0})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.JoinLineupScRsp, &pb.JoinLineupScRsp{Retcode: 1})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[req.Index]
			if lu == nil {
				ctx.Send(packet.JoinLineupScRsp, &pb.JoinLineupScRsp{Retcode: 1})
				return
			}
			slot := int(req.Slot)
			if slot < 0 {
				ctx.Send(packet.JoinLineupScRsp, &pb.JoinLineupScRsp{Retcode: 1})
				return
			}
			for len(lu.AvatarIDs) <= slot {
				lu.AvatarIDs = append(lu.AvatarIDs, 0)
			}
			lu.AvatarIDs[slot] = req.BaseAvatarId
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.JoinLineupScRsp, &pb.JoinLineupScRsp{Retcode: 0})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.QuitLineupCsReq, packet.Handler{
		Name:     packet.Name(packet.QuitLineupCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.QuitLineupCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.QuitLineupScRsp, &pb.QuitLineupScRsp{Retcode: 0, BaseAvatarId: req.BaseAvatarId, PlaneId: req.PlaneId, IsVirtual: req.IsVirtual})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.QuitLineupScRsp, &pb.QuitLineupScRsp{Retcode: 1})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[req.Index]
			if lu == nil {
				ctx.Send(packet.QuitLineupScRsp, &pb.QuitLineupScRsp{Retcode: 1})
				return
			}
			next := make([]uint32, 0, len(lu.AvatarIDs))
			for _, id := range lu.AvatarIDs {
				if id == 0 || id == req.BaseAvatarId {
					continue
				}
				next = append(next, id)
			}
			lu.AvatarIDs = next
			if lu.LeaderSlot >= uint32(len(lu.AvatarIDs)) {
				lu.LeaderSlot = 0
			}
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.QuitLineupScRsp, &pb.QuitLineupScRsp{Retcode: 0, BaseAvatarId: req.BaseAvatarId, PlaneId: req.PlaneId, IsVirtual: req.IsVirtual})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.SwapLineupCsReq, packet.Handler{
		Name:     packet.Name(packet.SwapLineupCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SwapLineupCsReq
			_ = proto.Unmarshal(payload, &req)
			if s.db == nil {
				ctx.Send(packet.SwapLineupScRsp, &pb.SwapLineupScRsp{Retcode: 0})
				return
			}
			p, ok := s.db.GetPlayerByUID(ctx.Session.UID)
			if !ok || p == nil {
				ctx.Send(packet.SwapLineupScRsp, &pb.SwapLineupScRsp{Retcode: 1})
				return
			}
			lineup.EnsureDefaults(p)
			lu := p.Lineups[req.Index]
			if lu == nil {
				ctx.Send(packet.SwapLineupScRsp, &pb.SwapLineupScRsp{Retcode: 1})
				return
			}
			a := int(req.SrcSlot)
			b := int(req.DstSlot)
			if a < 0 || b < 0 || a >= len(lu.AvatarIDs) || b >= len(lu.AvatarIDs) {
				ctx.Send(packet.SwapLineupScRsp, &pb.SwapLineupScRsp{Retcode: 1})
				return
			}
			lu.AvatarIDs[a], lu.AvatarIDs[b] = lu.AvatarIDs[b], lu.AvatarIDs[a]
			_ = s.db.SavePlayer(p)
			ctx.Send(packet.SwapLineupScRsp, &pb.SwapLineupScRsp{Retcode: 0})
			sendSyncLineupNotify(ctx, p, lineup.ToProto(p, s.scene.PlaneID, req.Index))
		},
	})

	s.reg.Register(packet.GetLineupAvatarDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetLineupAvatarDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
			_ = s.ensurePlayerDefaults(p)
			lineup.EnsureDefaults(p)
			ctx.Send(packet.GetLineupAvatarDataScRsp, &pb.GetLineupAvatarDataScRsp{
				Retcode:        0,
				AvatarDataList: s.buildLineupAvatarData(p),
			})
		},
	})

	s.reg.Register(packet.RecoverAllLineupCsReq, packet.Handler{
		Name:     packet.Name(packet.RecoverAllLineupCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			// TODO: hook to avatar hp/mp; for now ack success to avoid client waits.
			ctx.Send(packet.RecoverAllLineupScRsp, &pb.RecoverAllLineupScRsp{Retcode: 0})
		},
	})
}
