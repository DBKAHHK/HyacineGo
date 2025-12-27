package gameserver

import (
	"hyacine-server/internal/gameserver/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func (s *Server) registerPropHandlers() {
	s.reg.Register(packet.InteractPropCsReq, packet.Handler{
		Name:     packet.Name(packet.InteractPropCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.InteractPropCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.InteractPropScRsp, &pb.InteractPropScRsp{
				Retcode:      0,
				PropEntityId: req.PropEntityId,
				PropState:    8,
			})
		},
	})

	s.reg.Register(packet.UpdateGroupPropertyCsReq, packet.Handler{
		Name:     packet.Name(packet.UpdateGroupPropertyCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.UpdateGroupPropertyCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.UpdateGroupPropertyScRsp, &pb.UpdateGroupPropertyScRsp{
				Retcode:     0,
				GroupId:     req.GroupId,
				FloorId:     req.FloorId,
				DimensionId: req.DimensionId,
				DELEKFMGGCM: req.DELEKFMGGCM,
			})
		},
	})

	s.reg.Register(packet.ChangePropTimelineInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.ChangePropTimelineInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.ChangePropTimelineInfoCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.ChangePropTimelineInfoScRsp, &pb.ChangePropTimelineInfoScRsp{
				Retcode:      0,
				PropEntityId: req.PropEntityId,
			})
		},
	})
}
