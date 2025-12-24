package gameserver

import (
	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/player"
)

func heroAvatarID(p *player.Player) uint32 {
	gender := uint32(player.GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	growth := uint32(8002)
	if gender == player.GenderMan {
		growth = 8001
	}
	if p != nil && p.CurrentMultiPath != nil {
		if hero := p.CurrentMultiPath[growth]; hero != 0 {
			return hero
		}
	}
	return 0
}

func ensureLineupHero(lu *player.LineupData, hero uint32) bool {
	if lu == nil {
		return false
	}
	if hero == 0 {
		return false
	}
	for _, id := range lu.AvatarIDs {
		if id == hero {
			return false
		}
	}
	lu.AvatarIDs = append([]uint32{hero}, lu.AvatarIDs...)
	if lu.LeaderSlot >= uint32(len(lu.AvatarIDs)) {
		lu.LeaderSlot = 0
	}
	return true
}

func ensureDisplayHero(p *player.Player, hero uint32) bool {
	if p == nil {
		return false
	}
	if hero == 0 {
		return false
	}
	for _, da := range p.DisplayAvatars {
		if da.AvatarID == hero {
			return false
		}
	}
	p.DisplayAvatars = append([]player.DisplayAvatar{{AvatarID: hero, Pos: 1}}, p.DisplayAvatars...)
	return true
}

func isTrailblazerAvatarID(id uint32) bool {
	switch id {
	case 8001, 8002, 8003, 8004, 8005, 8006, 8007, 8008:
		return true
	default:
		return false
	}
}

func normalizeTrailblazerInLineup(lu *player.LineupData, current uint32) bool {
	if lu == nil || current == 0 {
		return false
	}
	hasTB := false
	next := make([]uint32, 0, len(lu.AvatarIDs)+1)
	for _, id := range lu.AvatarIDs {
		if id == 0 {
			continue
		}
		if isTrailblazerAvatarID(id) {
			hasTB = true
			continue
		}
		next = append(next, id)
	}
	if !hasTB {
		return false
	}
	next = append([]uint32{current}, next...)
	same := len(next) == len(lu.AvatarIDs)
	if same {
		for i := range next {
			if next[i] != lu.AvatarIDs[i] {
				same = false
				break
			}
		}
	}
	if same {
		return false
	}
	lu.AvatarIDs = next
	if lu.LeaderSlot >= uint32(len(lu.AvatarIDs)) {
		lu.LeaderSlot = 0
	}
	return true
}

func normalizeTrailblazerInDisplay(p *player.Player, current uint32) bool {
	if p == nil || current == 0 || len(p.DisplayAvatars) == 0 {
		return false
	}
	mutated := false
	seenCurrent := false
	out := make([]player.DisplayAvatar, 0, len(p.DisplayAvatars))
	for _, da := range p.DisplayAvatars {
		if !isTrailblazerAvatarID(da.AvatarID) {
			out = append(out, da)
			continue
		}
		if da.AvatarID != current {
			mutated = true
		}
		if seenCurrent {
			mutated = true
			continue
		}
		da.AvatarID = current
		seenCurrent = true
		out = append(out, da)
	}
	if !mutated {
		return false
	}
	p.DisplayAvatars = out
	return true
}

// ensurePlayerDefaults now only guards against nil maps/slices and preserves
// whatever values already exist in the player file/template. It no longer
// tries to override gender, multi-path choices, or lineup contents beyond
// ensuring the current hero is represented.
func (s *Server) ensurePlayerDefaults(p *player.Player) bool {
	if p == nil {
		return false
	}

	mutated := false

	if len(p.Lineups) == 0 {
		p.Lineups = map[uint32]*player.LineupData{}
		mutated = true
	}
	lineup.EnsureDefaults(p)

	if p.AvatarEnhancedIDs == nil {
		p.AvatarEnhancedIDs = map[uint32]uint32{}
		mutated = true
	}

	if p.CurrentMultiPath == nil {
		p.CurrentMultiPath = map[uint32]uint32{}
		mutated = true
	}

	hero := s.currentTrailblazerAvatarID(p)
	if hero != 0 {
		for _, lu := range p.Lineups {
			if normalizeTrailblazerInLineup(lu, hero) {
				mutated = true
			}
		}
		if normalizeTrailblazerInDisplay(p, hero) {
			mutated = true
		}
	}
	if lu, ok := p.Lineups[0]; ok {
		if ensureLineupHero(lu, hero) {
			mutated = true
		}
	}
	if ensureDisplayHero(p, hero) {
		mutated = true
	}

	if mutatedSpawn := s.ensurePlayerSpawn(p); mutatedSpawn {
		mutated = true
	}

	return mutated
}
