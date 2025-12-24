package gameserver

import (
	"sort"
	"strconv"
	"time"

	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

func expandOwnedMultiPathAvatarIDs(ids []uint32, p *player.Player) []uint32 {
	seen := map[uint32]struct{}{}
	out := make([]uint32, 0, len(ids)+8)
	hasMarch := false
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
		if id == 1001 || id == 1224 {
			hasMarch = true
		}
	}

	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	if gender == player.GenderMan {
		for _, id := range []uint32{8001, 8003, 8005, 8007} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	} else {
		for _, id := range []uint32{8002, 8004, 8006, 8008} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}

	if hasMarch {
		for _, id := range []uint32{1001, 1224} {
			if _, ok := seen[id]; ok {
				continue
			}
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func (s *Server) ownedAvatarIDs(p *player.Player) []uint32 {
	if s.freesr != nil && len(s.freesr.Avatars) > 0 {
		ids := expandOwnedMultiPathAvatarIDs(s.freesr.AvatarIDs(), p)
		return filterMultiPathAvatarIDs(ids, p)
	}
	ids := expandOwnedMultiPathAvatarIDs(lineup.OwnedAvatarIDs(p), p)
	return filterMultiPathAvatarIDs(ids, p)
}

func (s *Server) ownedAvatarIDsRaw(p *player.Player) []uint32 {
	if s.freesr != nil && len(s.freesr.Avatars) > 0 {
		return expandOwnedMultiPathAvatarIDs(s.freesr.AvatarIDs(), p)
	}
	return expandOwnedMultiPathAvatarIDs(lineup.OwnedAvatarIDs(p), p)
}

func (s *Server) buildAvatarList(p *player.Player) []*pb.Avatar {
	ids := s.ownedAvatarIDs(p)
	now := uint64(time.Now().Unix())
	out := make([]*pb.Avatar, 0, len(ids))
	for _, id := range ids {
		av := &pb.Avatar{
			BaseAvatarId:      id,
			Level:             70,
			Promotion:         6,
			Exp:               0,
			Rank:              6,
			DressedSkinId:     0,
			EquipmentUniqueId: 0,
			FirstMetTimeStamp: now,
		}
		if s.freesr != nil {
			av.EquipmentUniqueId = s.freesrEquipmentUniqueIDForAvatar(id)
			av.EquipRelicList = s.freesrEquipRelicListForAvatar(id)
			if preset, ok := s.freesr.GetAvatar(id); ok {
				if preset.Level > 0 {
					av.Level = preset.Level
				} else if p.Level > 0 {
					av.Level = p.Level
				}
				av.Promotion = preset.Promotion
				av.Rank = preset.Data.Rank
				av.UnkEnhancedId = preset.EnhancedID
				if len(preset.Data.Skills) > 0 {
					trees := make([]*pb.AvatarSkillTree, 0, len(preset.Data.Skills))
					for k, v := range preset.Data.Skills {
						point, err := strconv.ParseUint(k, 10, 32)
						if err != nil {
							continue
						}
						trees = append(trees, &pb.AvatarSkillTree{PointId: uint32(point), Level: uint32(v)})
					}
					av.SkilltreeList = trees
				}
			}
		} else if p.Level > 0 {
			av.Level = p.Level
		}
		if p != nil && len(p.AvatarEnhancedIDs) > 0 {
			if v, ok := p.AvatarEnhancedIDs[id]; ok {
				av.UnkEnhancedId = v
			} else if v, ok := p.AvatarEnhancedIDs[growthAvatarID(id)]; ok {
				av.UnkEnhancedId = v
			}
		}
		out = append(out, av)
	}
	return out
}

func growthAvatarID(id uint32) uint32 {
	switch id {
	case 8001, 8003, 8005, 8007:
		return 8001
	case 8002, 8004, 8006, 8008:
		return 8002
	case 1001, 1224:
		return 1001
	default:
		return id
	}
}

func filterMultiPathAvatarIDs(ids []uint32, p *player.Player) []uint32 {
	if len(ids) == 0 {
		return ids
	}
	isTBMale := func(id uint32) bool { return id == 8001 || id == 8003 || id == 8005 || id == 8007 }
	isTBFemale := func(id uint32) bool { return id == 8002 || id == 8004 || id == 8006 || id == 8008 }
	isMarch7 := func(id uint32) bool { return id == 1001 || id == 1224 }

	var hasTBMale, hasTBFemale, hasMarch7 bool
	for _, id := range ids {
		switch {
		case isTBMale(id):
			hasTBMale = true
		case isTBFemale(id):
			hasTBFemale = true
		case isMarch7(id):
			hasMarch7 = true
		}
	}
	if !hasTBMale && !hasTBFemale && !hasMarch7 {
		return ids
	}

	tbGender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		tbGender = p.Gender
	}
	tbGrowth := uint32(8002)
	if tbGender == player.GenderMan {
		tbGrowth = 8001
	}

	getCurrent := func(growth uint32, fallback uint32) uint32 {
		if p != nil && p.CurrentMultiPath != nil {
			if v, ok := p.CurrentMultiPath[growth]; ok && v != 0 {
				return v
			}
		}
		return fallback
	}

	tbDefault := uint32(8006)
	if tbGrowth == 8001 {
		tbDefault = 8005
	}
	tbCurrent := getCurrent(tbGrowth, tbDefault)
	marchCurrent := getCurrent(1001, 1001)

	out := make([]uint32, 0, len(ids))
	for _, id := range ids {
		switch {
		case isTBMale(id) || isTBFemale(id):
			if id == tbCurrent {
				out = append(out, id)
			}
		case isMarch7(id):
			if id == marchCurrent {
				out = append(out, id)
			}
		default:
			out = append(out, id)
		}
	}

	if hasTBMale || hasTBFemale {
		found := false
		for _, id := range out {
			if isTBMale(id) || isTBFemale(id) {
				found = true
				break
			}
		}
		if !found {
			for _, id := range ids {
				if tbGrowth == 8001 && isTBMale(id) {
					out = append(out, id)
					break
				}
				if tbGrowth == 8002 && isTBFemale(id) {
					out = append(out, id)
					break
				}
			}
		}
	}

	return out
}

func (s *Server) buildMultiPathAvatarInfoList(p *player.Player) []*pb.MultiPathAvatarInfo {
	ids := s.ownedAvatarIDsRaw(p)
	hasMarch := false
	for _, id := range ids {
		if id == 1001 || id == 1224 {
			hasMarch = true
			break
		}
	}

	var want []uint32
	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	if gender == player.GenderMan {
		want = append(want, 8001, 8003, 8005, 8007)
	} else {
		want = append(want, 8002, 8004, 8006, 8008)
	}
	if hasMarch {
		want = append(want, 1001, 1224)
	}

	out := make([]*pb.MultiPathAvatarInfo, 0, len(want))
	for _, id := range want {
		out = append(out, &pb.MultiPathAvatarInfo{
			AvatarId:        pb.MultiPathAvatarType(id),
			Rank:            6,
			UnkEnhancedId:   0,
			DressedSkinId:   0,
			PathEquipmentId: 0,
		})
	}
	return out
}

func (s *Server) currentTrailblazerAvatarID(p *player.Player) uint32 {
	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	growth := uint32(8002)
	if gender == player.GenderMan {
		growth = 8001
	}
	if p != nil && p.CurrentMultiPath != nil {
		if v := p.CurrentMultiPath[growth]; v != 0 {
			return v
		}
	}
	if growth == 8001 {
		return 8005
	}
	return 8006
}

func (s *Server) multiPathInfosForGender(p *player.Player) []*pb.MultiPathAvatarInfo {
	ids := s.ownedAvatarIDsRaw(p)
	hasMarch := false
	for _, id := range ids {
		if id == 1001 || id == 1224 {
			hasMarch = true
			break
		}
	}
	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	want := make([]uint32, 0, 8)
	if gender == player.GenderMan {
		want = append(want, 8001, 8003, 8005, 8007)
	} else {
		want = append(want, 8002, 8004, 8006, 8008)
	}
	if hasMarch {
		want = append(want, 1001, 1224)
	}

	seen := map[uint32]struct{}{}
	for _, id := range want {
		seen[id] = struct{}{}
	}
	curTB := s.currentTrailblazerAvatarID(p)
	if curTB != 0 {
		if _, ok := seen[curTB]; !ok {
			want = append(want, curTB)
		}
	}
	if len(want) == 0 {
		return nil
	}
	out := make([]*pb.MultiPathAvatarInfo, 0, len(want))
	for _, id := range want {
		out = append(out, &pb.MultiPathAvatarInfo{
			AvatarId:        pb.MultiPathAvatarType(id),
			Rank:            6,
			UnkEnhancedId:   0,
			DressedSkinId:   0,
			PathEquipmentId: 0,
		})
	}
	return out
}
