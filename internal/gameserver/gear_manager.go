package gameserver

import (
	"sort"

	"hyacine-server/internal/gameserver/freesrdata"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

const (
	freesrEquipmentUIDBase uint32 = 2_000_000_000
	freesrRelicUIDBase     uint32 = 3_000_000_000
)

func freesrEquipmentUniqueID(lc freesrdata.LightconePreset) uint32 {
	if lc.InternalUID != 0 {
		return lc.InternalUID
	}
	if lc.EquipAvatar == 0 || lc.ItemID == 0 {
		return 0
	}
	// Keep stable and within uint32. Most lightcone item_ids fit in <100000.
	return freesrEquipmentUIDBase + lc.EquipAvatar*100_000 + (lc.ItemID % 100_000)
}

func freesrRelicUniqueID(r freesrdata.RelicPreset) uint32 {
	if r.InternalUID != 0 {
		return r.InternalUID
	}
	if r.EquipAvatar == 0 || r.RelicID == 0 {
		return 0
	}
	return freesrRelicUIDBase + r.EquipAvatar*100_000 + (r.RelicID % 100_000)
}

func relicSlotTypeFromTID(tid uint32) uint32 {
	// Empirically FreeSR relic ids end with 1..6 for the slot.
	t := tid % 10
	if t >= 1 && t <= 6 {
		return t
	}
	return 0
}

func (s *Server) buildEquipmentList(_ *player.Player) []*pb.Equipment {
	if s == nil || s.freesr == nil || len(s.freesr.Lightcones) == 0 {
		return nil
	}

	out := make([]*pb.Equipment, 0, len(s.freesr.Lightcones))
	seen := map[uint32]struct{}{}

	for _, lc := range s.freesr.Lightcones {
		if lc.ItemID == 0 {
			continue
		}
		uid := freesrEquipmentUniqueID(lc)
		if uid == 0 {
			continue
		}
		if _, dup := seen[uid]; dup {
			continue
		}
		seen[uid] = struct{}{}
		out = append(out, &pb.Equipment{
			DressAvatarId: lc.EquipAvatar,
			Rank:          lc.Rank,
			Tid:           lc.ItemID,
			Promotion:     lc.Promotion,
			Exp:           0,
			UniqueId:      uid,
			Level:         lc.Level,
			IsProtected:   false,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].UniqueId < out[j].UniqueId })
	return out
}

func (s *Server) buildRelicList(_ *player.Player) []*pb.Relic {
	if s == nil || s.freesr == nil || len(s.freesr.Relics) == 0 {
		return nil
	}

	out := make([]*pb.Relic, 0, len(s.freesr.Relics))
	seen := map[uint32]struct{}{}

	for _, r := range s.freesr.Relics {
		if r.RelicID == 0 {
			continue
		}
		uid := freesrRelicUniqueID(r)
		if uid == 0 {
			continue
		}
		if _, dup := seen[uid]; dup {
			continue
		}
		seen[uid] = struct{}{}

		sub := make([]*pb.RelicAffix, 0, len(r.SubAffixes))
		for _, sa := range r.SubAffixes {
			if sa.SubAffixID == 0 {
				continue
			}
			sub = append(sub, &pb.RelicAffix{
				AffixId: sa.SubAffixID,
				Cnt:     sa.Count,
				Step:    sa.Step,
			})
		}

		out = append(out, &pb.Relic{
			MainAffixId:         r.MainAffixID,
			Exp:                 0,
			Level:               r.Level,
			ReforgeBlockSubAffixId: 0,
			IsProtected:         false,
			ReforgeSubAffixList: nil,
			IsDiscarded:         false,
			SubAffixList:        sub,
			DressAvatarId:       r.EquipAvatar,
			UniqueId:            uid,
			Tid:                 r.RelicID,
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].UniqueId < out[j].UniqueId })
	return out
}

func (s *Server) freesrEquipmentUniqueIDForAvatar(avatarID uint32) uint32 {
	if s == nil || s.freesr == nil || avatarID == 0 {
		return 0
	}
	for _, lc := range s.freesr.Lightcones {
		if lc.EquipAvatar != avatarID {
			continue
		}
		uid := freesrEquipmentUniqueID(lc)
		if uid != 0 {
			return uid
		}
	}
	return 0
}

func (s *Server) freesrEquipRelicListForAvatar(avatarID uint32) []*pb.EquipRelic {
	if s == nil || s.freesr == nil || avatarID == 0 || len(s.freesr.Relics) == 0 {
		return nil
	}
	byType := map[uint32]uint32{}
	for _, r := range s.freesr.Relics {
		if r.EquipAvatar != avatarID || r.RelicID == 0 {
			continue
		}
		uid := freesrRelicUniqueID(r)
		if uid == 0 {
			continue
		}
		t := relicSlotTypeFromTID(r.RelicID)
		if t == 0 {
			continue
		}
		byType[t] = uid
	}
	if len(byType) == 0 {
		return nil
	}
	types := make([]uint32, 0, len(byType))
	for t := range byType {
		types = append(types, t)
	}
	sort.Slice(types, func(i, j int) bool { return types[i] < types[j] })
	out := make([]*pb.EquipRelic, 0, len(types))
	for _, t := range types {
		out = append(out, &pb.EquipRelic{RelicUniqueId: byType[t], Type: t})
	}
	return out
}
