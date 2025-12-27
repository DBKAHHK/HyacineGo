package gameserver

import (
	"hyacine-server/internal/gameserver/packet"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func (s *Server) registerItemHandlers() {
	s.reg.Register(packet.GetMarkItemListCsReq, packet.Handler{
		Name:     packet.Name(packet.GetMarkItemListCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			// Required by the map UI; return empty list for a clean PS baseline.
			ctx.Send(packet.GetMarkItemListScRsp, &pb.GetMarkItemListScRsp{Retcode: 0})
		},
	})

	s.reg.Register(packet.MarkItemCsReq, packet.Handler{
		Name:     packet.Name(packet.MarkItemCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			var req pb.MarkItemCsReq
			_ = proto.Unmarshal(payload, &req)
			ctx.Send(packet.MarkItemScRsp, &pb.MarkItemScRsp{
				Retcode:     0,
				JBNEDFJJDOM: req.JBNEDFJJDOM,
				ItemId:      req.ItemId,
			})
		},
	})
}
