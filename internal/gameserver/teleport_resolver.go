package gameserver

import (
	"math"

	pb "hyacine-server/internal/proto/gen"
)

type teleportTarget struct {
	PlaneID uint32
	FloorID uint32
	EntryID uint32
	Pos     *pb.Vector
}

func coordFromFloat(v float32) int32 {
	return int32(math.Round(float64(v * 1000)))
}

func (s *Server) resolveTeleportTarget(teleportID uint32) (teleportTarget, bool) {
	var out teleportTarget
	if s == nil || s.data == nil || teleportID == 0 {
		return out, false
	}

	tc, err := s.data.TeleportConfig()
	if err != nil || len(tc) == 0 {
		return out, false
	}
	row := tc[teleportID]
	if row == nil || row.ID == 0 {
		return out, false
	}
	out.PlaneID = row.PlaneID
	out.FloorID = row.FloorID

	if me, err := s.data.MapEntrance(); err == nil && len(me) > 0 {
		for _, r := range me {
			if r == nil || r.ID == 0 {
				continue
			}
			if r.PlaneID == out.PlaneID && r.FloorID == out.FloorID {
				out.EntryID = r.ID
				break
			}
		}
	}

	if out.FloorID == 0 || row.GroupID == 0 || row.ConfigID == 0 {
		return out, true
	}

	floor, err := s.data.LevelOutputFloor(int(out.FloorID))
	if err != nil || floor == nil {
		return out, true
	}
	groupPath := ""
	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.ID == 0 || gi.GroupPath == "" {
			continue
		}
		if uint32(gi.ID) == row.GroupID {
			groupPath = gi.GroupPath
			break
		}
	}
	if groupPath == "" {
		return out, true
	}

	group, err := s.data.LoadLevelOutputGroupByPath(groupPath)
	if err != nil || group == nil {
		return out, true
	}
	for _, a := range group.AnchorList {
		if a.IsDelete || a.ID == 0 {
			continue
		}
		if uint32(a.ID) == row.ConfigID {
			out.Pos = &pb.Vector{X: coordFromFloat(a.PosX), Y: coordFromFloat(a.PosY), Z: coordFromFloat(a.PosZ)}
			break
		}
	}

	return out, true
}

