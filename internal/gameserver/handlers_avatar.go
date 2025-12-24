package gameserver

import (
	"hyacine-server/internal/gameserver/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func (s *Server) registerAvatarHandlers() {
	s.reg.Register(packet.SetMultipleAvatarPathsCsReq, packet.Handler{
		Name:     packet.Name(packet.SetMultipleAvatarPathsCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetMultipleAvatarPathsCsReq
			_ = proto.Unmarshal(payload, &req)

			if s.db != nil {
				p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if p.CurrentMultiPath == nil {
					p.CurrentMultiPath = map[uint32]uint32{}
				}
				for _, t := range req.AvatarIdList {
					id := uint32(t)
					if id == 0 {
						continue
					}
					g := growthAvatarID(id)
					p.CurrentMultiPath[g] = id
				}
				p.Touch()
				_ = s.db.SavePlayer(p)

				// Push an updated sync so the avatar list doesn't show duplicates and reflects the chosen path.
				ctx.Send(packet.PlayerSyncScNotify, &pb.PlayerSyncScNotify{
					AvatarSync:              &pb.AvatarSync{AvatarList: s.buildAvatarList(p)},
					MultiPathAvatarInfoList: s.buildMultiPathAvatarInfoList(p),
					SyncStatus:              &pb.SyncStatus{},
				})
			}

			ctx.Send(packet.SetMultipleAvatarPathsScRsp, &pb.SetMultipleAvatarPathsScRsp{Retcode: 0})
		},
	})

	s.reg.Register(packet.SetAvatarEnhancedIdCsReq, packet.Handler{
		Name:     packet.Name(packet.SetAvatarEnhancedIdCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.SetAvatarEnhancedIdCsReq
			_ = proto.Unmarshal(payload, &req)

			if s.db != nil {
				p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && loaded != nil {
					p = loaded
				}
				if p.AvatarEnhancedIDs == nil {
					p.AvatarEnhancedIDs = map[uint32]uint32{}
				}
				if req.AvatarId != 0 {
					growth := growthAvatarID(req.AvatarId)
					// Store both the path-variant id and the growth id.
					p.AvatarEnhancedIDs[req.AvatarId] = req.EnhancedId
					p.AvatarEnhancedIDs[growth] = req.EnhancedId
					p.Touch()
					_ = s.db.SavePlayer(p)
				}
			}

			// This response is required for the in-game enhanced-id popup to close.
			ctx.Send(packet.SetAvatarEnhancedIdScRsp, &pb.SetAvatarEnhancedIdScRsp{
				Retcode:        0,
				GrowthAvatarId: growthAvatarID(req.AvatarId),
				UnkEnhancedId:  req.EnhancedId,
			})
		},
	})
}
