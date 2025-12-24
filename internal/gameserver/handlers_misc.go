package gameserver

import (
	"time"

	"google.golang.org/protobuf/proto"

	"hyacine-server/internal/gameserver/packet"
	pb "hyacine-server/internal/proto/gen"
)

func (s *Server) registerMiscHandlers() {
	s.reg.Register(packet.SetClientPausedCsReq, packet.Handler{
		Name:     packet.Name(packet.SetClientPausedCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SetClientPausedCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.SetClientPausedScRsp, &pb.SetClientPausedScRsp{
				Retcode: 0,
				Paused:  req.GetPaused(),
			})

			// 3.8.0src sends ClientDownloadDataScNotify immediately after SetClientPausedScRsp.
			if s != nil && s.opts.SendClientDownloadData {
				ctx.Send(packet.ClientDownloadDataScNotify, &pb.ClientDownloadDataScNotify{
					DownloadData: &pb.ClientDownloadData{
						Data:    s.getClientDownloadData(),
						Version: 81,
						Time:    time.Now().Unix(),
					},
				})
			}
		},
	})

	s.reg.Register(packet.GetArchiveDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetArchiveDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetArchiveDataScRsp, &pb.GetArchiveDataScRsp{
				Retcode:     0,
				ArchiveData: &pb.ArchiveData{},
			})
		},
	})

	s.reg.Register(packet.GetUpdatedArchiveDataCsReq, packet.Handler{
		Name:     packet.Name(packet.GetUpdatedArchiveDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.GetUpdatedArchiveDataScRsp, &pb.GetUpdatedArchiveDataScRsp{
				Retcode:     0,
				ArchiveData: &pb.ArchiveData{},
			})
		},
	})

	s.reg.Register(packet.ContentPackageGetDataCsReq, packet.Handler{
		Name:     packet.Name(packet.ContentPackageGetDataCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			ctx.Send(packet.ContentPackageGetDataScRsp, &pb.ContentPackageGetDataScRsp{
				Retcode: 0,
				Data: &pb.ContentPackageData{
					CurContentId:       0,
					ContentPackageList: nil,
				},
			})
		},
	})
}
