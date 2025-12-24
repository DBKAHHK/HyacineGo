package gameserver

import (
	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/packet"
	"hyacine-server/internal/gameserver/player"
	"hyacine-server/internal/gameserver/scene"
	"log/slog"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

func (s *Server) registerSceneHandlers() {
	s.reg.Register(packet.GetEnteredSceneCsReq, packet.Handler{
		Name:     packet.Name(packet.GetEnteredSceneCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			entered := scene.BuildEnteredScenes(s.scene)
			if s.data != nil {
				if me, err := s.data.MapEntrance(); err == nil && len(me) > 0 {
					type key struct {
						plane uint32
						floor uint32
					}
					seen := map[key]struct{}{}
					list := make([]*pb.EnteredSceneInfo, 0, len(me))
					for _, row := range me {
						if row == nil || row.PlaneID == 0 || row.FloorID == 0 {
							continue
						}
						k := key{plane: uint32(row.PlaneID), floor: uint32(row.FloorID)}
						if _, ok := seen[k]; ok {
							continue
						}
						seen[k] = struct{}{}
						list = append(list, &pb.EnteredSceneInfo{PlaneId: k.plane, FloorId: k.floor})
					}
					if len(list) > 0 {
						entered = list
					}
				}
			}

			// Match SR-CasPS behavior: send the notify before the response.
			ctx.Send(packet.EnteredSceneChangeScNotify, &pb.EnteredSceneChangeScNotify{
				EnteredSceneInfoList: entered,
			})
			ctx.Send(packet.GetEnteredSceneScRsp, &pb.GetEnteredSceneScRsp{
				Retcode:              0,
				EnteredSceneInfoList: entered,
			})
		},
	})

	s.reg.Register(packet.GetCurSceneInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetCurSceneInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, _ []byte) {
			var p *player.Player
			if s.db != nil {
				if loaded, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok {
					p = loaded
				}
			}
			activeScene := s.scene
			if p != nil {
				if p.EntryID != 0 {
					activeScene.EntryID = p.EntryID
				}
				if p.PlaneID != 0 {
					activeScene.PlaneID = p.PlaneID
				}
				if p.FloorID != 0 {
					activeScene.FloorID = p.FloorID
				}
				// Best-effort world id correction if plane differs.
				if s.data != nil {
					if mp, err := s.data.MazePlane(); err == nil {
						if row := mp[activeScene.PlaneID]; row != nil && row.WorldID != 0 {
							activeScene.WorldID = uint32(row.WorldID)
						}
					}
				}
			}
			ctx.Send(packet.GetCurSceneInfoScRsp, &pb.GetCurSceneInfoScRsp{
				Retcode: 0,
				Scene:   scene.BuildSceneInfoForPlayerWithData(s.data, activeScene, ctx.Session.UID, p),
			})

			// In some client flows the game never sends EnterSceneCsReq automatically.
			// Push a single initial EnterSceneByServerScNotify after the first GetCurSceneInfo to trigger map loading.
			if ctx.Session != nil && !ctx.Session.HasEnteredWorld {
				ctx.Session.HasEnteredWorld = true
				ctx.Send(packet.SyncServerSceneChangeNotify, &pb.SyncServerSceneChangeNotify{})

				// Unlock all area maps (entry ids) early so the client can show the full map UI.
				if s.data != nil {
					if me, err := s.data.MapEntrance(); err == nil && len(me) > 0 {
						entryIDs := make([]uint32, 0, len(me))
						for _, row := range me {
							if row == nil || row.ID == 0 {
								continue
							}
							entryIDs = append(entryIDs, uint32(row.ID))
						}
						if len(entryIDs) > 0 {
							ctx.Send(packet.UnlockedAreaMapScNotify, &pb.UnlockedAreaMapScNotify{EntryIdList: entryIDs})
						}
					}
				}

				ctx.Send(packet.EnterSceneByServerScNotify, &pb.EnterSceneByServerScNotify{
					Reason: pb.EnterSceneReason_ENTER_SCENE_REASON_NONE,
					Lineup: buildLineupInfoForPlayer(ctx.Session.UID, p),
					Scene:  scene.BuildSceneInfoForPlayerWithData(s.data, activeScene, ctx.Session.UID, p),
				})
			}
		},
	})

	s.reg.Register(packet.EnterSceneCsReq, packet.Handler{
		Name:     packet.Name(packet.EnterSceneCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.EnterSceneCsReq{}
			if err := proto.Unmarshal(payload, req); err != nil {
				ctx.Send(packet.EnterSceneScRsp, &pb.EnterSceneScRsp{Retcode: 1})
				return
			}

			p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID)

			// Ack first.
			ctx.Send(packet.EnterSceneScRsp, &pb.EnterSceneScRsp{
				Retcode:         0,
				GameStoryLineId: req.GetGameStoryLineId(),
				IsCloseMap:      req.GetIsCloseMap(),
				ContentId:       req.GetContentId(),
				IsOverMap:       false,
			})

			// Send scene+lineup snapshot to actually enter the world.
			activeScene := s.scene
			entryID := req.GetEntryId()
			teleportID := req.GetTeleportId()

			// Some client flows provide only teleport_id, with entry_id left as 0.
			if teleportID != 0 {
				if tgt, ok := s.resolveTeleportTarget(teleportID); ok {
					if entryID == 0 && tgt.EntryID != 0 {
						entryID = tgt.EntryID
					}
					if tgt.PlaneID != 0 {
						activeScene.PlaneID = tgt.PlaneID
					}
					if tgt.FloorID != 0 {
						activeScene.FloorID = tgt.FloorID
					}
					// If client didn't provide a target position, use the teleport anchor position when available.
					if pos := req.GetBCPPMONGJGF(); pos != nil {
						p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
					} else if tgt.Pos != nil {
						p.X, p.Y, p.Z = tgt.Pos.GetX(), tgt.Pos.GetY(), tgt.Pos.GetZ()
					}
				}
			} else if pos := req.GetBCPPMONGJGF(); pos != nil {
				// Some requests only provide an explicit position.
				p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
			}

			// Resolve plane/floor from entry id when available.
			if s.data != nil && entryID != 0 {
				if me, err := s.data.MapEntrance(); err == nil {
					if row := me[entryID]; row != nil {
						if row.FloorID != 0 {
							activeScene.FloorID = uint32(row.FloorID)
						}
						if row.PlaneID != 0 {
							activeScene.PlaneID = uint32(row.PlaneID)
						}
					}
				}
			}

			activeScene.EntryID = entryID
			if activeScene.EntryID == 0 {
				if p.EntryID != 0 {
					activeScene.EntryID = p.EntryID
				} else {
					activeScene.EntryID = s.scene.EntryID
				}
			}
			if s.data != nil {
				if mp, err := s.data.MazePlane(); err == nil {
					if row := mp[activeScene.PlaneID]; row != nil && row.WorldID != 0 {
						activeScene.WorldID = uint32(row.WorldID)
					}
				}
			}

			// Persist selection so future GetCurSceneInfo uses it.
			if p != nil {
				p.EntryID = activeScene.EntryID
				p.PlaneID = activeScene.PlaneID
				p.FloorID = activeScene.FloorID
				if teleportID != 0 {
					p.TeleportID = teleportID
				}
				p.Touch()
				if s.db != nil {
					_ = s.db.SavePlayer(p)
				}
			}

			// Some clients require a scene-change notify before re-entering a scene (teleport UI).
			ctx.Send(packet.SyncServerSceneChangeNotify, &pb.SyncServerSceneChangeNotify{})
			ctx.Send(packet.EnterSceneByServerScNotify, &pb.EnterSceneByServerScNotify{
				Reason: pb.EnterSceneReason_ENTER_SCENE_REASON_NONE,
				Lineup: buildLineupInfoForPlayer(ctx.Session.UID, p),
				Scene:  scene.BuildSceneInfoForPlayerWithData(s.data, activeScene, ctx.Session.UID, p),
			})
		},
	})

	s.reg.Register(packet.SceneEntityMoveCsReq, packet.Handler{
		Name:     packet.Name(packet.SceneEntityMoveCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneEntityMoveCsReq{}
			_ = proto.Unmarshal(payload, req)

			// Best-effort persistence for the leader actor movement.
			if s.db != nil {
				if p, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && p != nil && len(req.GetEntityMotionList()) > 0 {
					for _, m := range req.GetEntityMotionList() {
						if m == nil || m.GetMotion() == nil || m.GetMotion().GetPos() == nil {
							continue
						}
						// We only persist motion updates for our actor entity (entity_id=1 in BuildSceneInfoForPlayer).
						if m.GetEntityId() != 1 {
							continue
						}
						pos := m.GetMotion().GetPos()
						p.X, p.Y, p.Z = pos.X, pos.Y, pos.Z
						p.Touch()
						_ = s.db.SavePlayer(p)
						break
					}
				}
			}

			ctx.Send(packet.SceneEntityMoveScRsp, &pb.SceneEntityMoveScRsp{
				Retcode:          0,
				EntityMotionList: req.GetEntityMotionList(),
				DownloadData:     nil,
			})
		},
	})

	s.reg.Register(packet.SceneEntityTeleportCsReq, packet.Handler{
		Name:     packet.Name(packet.SceneEntityTeleportCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.SceneEntityTeleportCsReq{}
			_ = proto.Unmarshal(payload, req)
			ctx.Send(packet.SceneEntityTeleportScRsp, &pb.SceneEntityTeleportScRsp{
				Retcode:      0,
				EntityMotion: req.GetEntityMotion(),
			})

			// Best-effort persistence + immediate move notify for map teleport.
			if p := s.getOrCreatePlayer(ctx.Session.UID, ctx.Session.AccountUID); p != nil {
				if em := req.GetEntityMotion(); em != nil && em.GetMotion() != nil && em.GetMotion().GetPos() != nil {
					pos := em.GetMotion().GetPos()
					p.X, p.Y, p.Z = pos.GetX(), pos.GetY(), pos.GetZ()
					if em.GetMapLayer() != 0 {
						p.MapLayer = em.GetMapLayer()
					}
					if req.GetEntryId() != 0 {
						p.EntryID = req.GetEntryId()
					}
					p.Touch()
					if s.db != nil {
						_ = s.db.SavePlayer(p)
					}

					entryID := req.GetEntryId()
					if entryID == 0 {
						entryID = p.EntryID
					}
					if entryID == 0 {
						entryID = s.scene.EntryID
					}
					ctx.Send(packet.SceneEntityMoveScNotify, &pb.SceneEntityMoveScNotify{
						Motion:   em.GetMotion(),
						EntryId:  entryID,
						EntityId: em.GetEntityId(),
					})
				}
			}
		},
	})

	s.reg.Register(packet.GetSceneMapInfoCsReq, packet.Handler{
		Name:     packet.Name(packet.GetSceneMapInfoCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetSceneMapInfoCsReq{}
			_ = proto.Unmarshal(payload, req)

			out := make([]*pb.SceneMapInfo, 0, len(req.GetEntryIdList())+len(req.GetFloorIdList())+1)

			// Try to infer floor ids from entry ids via MapEntrance table.
			entryToFloor := map[uint32]uint32{}
			if s.data != nil {
				if me, err := s.data.MapEntrance(); err == nil {
					for _, row := range me {
						if row == nil {
							continue
						}
						entryToFloor[uint32(row.ID)] = uint32(row.FloorID)
					}
				}
			}

			for _, entryID := range req.GetEntryIdList() {
				floorID := entryToFloor[entryID]
				if floorID == 0 {
					floorID = s.scene.FloorID
				}
				if s.data != nil {
					if info, err := scene.BuildSceneMapInfo(s.data, entryID, floorID); err == nil && info != nil {
						out = append(out, info)
						continue
					}
				}
				out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
			}
			for _, floorID := range req.GetFloorIdList() {
				entryID := s.scene.EntryID
				if s.data != nil {
					if info, err := scene.BuildSceneMapInfo(s.data, entryID, floorID); err == nil && info != nil {
						out = append(out, info)
						continue
					}
				}
				out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
			}

			if len(out) == 0 {
				entryID := s.scene.EntryID
				floorID := s.scene.FloorID
				if s.data != nil {
					if info, err := scene.BuildSceneMapInfo(s.data, entryID, floorID); err == nil && info != nil {
						out = append(out, info)
					} else {
						out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
					}
				} else {
					out = append(out, &pb.SceneMapInfo{Retcode: 0, EntryId: entryID, CurMapEntryId: entryID, FloorId: floorID})
				}
			}

			if len(out) > 0 {
				for _, info := range out {
					if info == nil {
						continue
					}
					slog.Debug("map info sent", "entry", info.EntryId, "floor", info.FloorId, "teleports", len(info.UnlockTeleportList))
				}
			} else {
				slog.Debug("map info sent empty", "entry_story_line", req.GetEntryStoryLineId(), "content", req.GetContentId())
			}

			ctx.Send(packet.GetSceneMapInfoScRsp, &pb.GetSceneMapInfoScRsp{
				Retcode:          0,
				SceneMapInfo:     out,
				EntryStoryLineId: req.GetEntryStoryLineId(),
				ContentId:        req.GetContentId(),
				Unk1:             req.GetUnk1(),
			})
		},
	})

	s.reg.Register(packet.GetUnlockTeleportCsReq, packet.Handler{
		Name:     packet.Name(packet.GetUnlockTeleportCsReq),
		MinState: packet.StateActive,
		Fn: func(ctx packet.Context, _ uint16, payload []byte) {
			req := &pb.GetUnlockTeleportCsReq{}
			_ = proto.Unmarshal(payload, req)

			var unlockList []uint32
			notifyEntryID := uint32(0)
			if s.data != nil {
				entries := req.GetEntryIdList()
				if len(entries) == 0 {
					// If client didn't provide entry list, unlock teleports for current entry.
					entries = []uint32{s.scene.EntryID}
					if ctx.Session != nil && ctx.Session.UID != 0 && s.db != nil {
						if p, ok := s.db.GetPlayerByUID(ctx.Session.UID); ok && p != nil && p.EntryID != 0 {
							entries = []uint32{p.EntryID}
						}
					}
				}
				if len(entries) > 0 {
					notifyEntryID = entries[0]
				}
				if out, err := scene.UnlockTeleportsForEntries(s.data, entries); err == nil {
					unlockList = out
				}
			}

			// Some clients expect individual unlock notifies for the map teleport UI.
			if notifyEntryID != 0 && len(unlockList) > 0 {
				for _, id := range unlockList {
					ctx.Send(packet.UnlockTeleportNotify, &pb.UnlockTeleportNotify{
						EntryId:    notifyEntryID,
						TeleportId: id,
					})
				}
			}
			ctx.Send(packet.GetUnlockTeleportScRsp, &pb.GetUnlockTeleportScRsp{
				Retcode:            0,
				UnlockTeleportList: unlockList,
			})
		},
	})
}

func buildLineupInfoForPlayer(uid uint32, p *player.Player) *pb.LineupInfo {
	if p == nil {
		p = player.New(uid, "")
	}
	return lineup.ToProto(p, 0, p.CurrentLineupIndex)
}
