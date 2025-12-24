package gameserver

import (
	"crypto/rand"
	"encoding/binary"
	"sort"

	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

func randomUint32() uint32 {
	var b [4]byte
	_, _ = rand.Read(b[:])
	return binary.LittleEndian.Uint32(b[:])
}

func (s *Server) freesrBattleConfig() *pb.SceneBattleInfo {
	p := s.getOrCreatePlayer(0, "")
	return s.buildSceneBattleInfo(p)
}

func (s *Server) buildSceneBattleInfo(p *player.Player) *pb.SceneBattleInfo {
	if s == nil {
		return &pb.SceneBattleInfo{}
	}
	if p == nil {
		p = player.New(0, "")
	}

	lineup.EnsureDefaults(p)

	stageID := uint32(0)
	roundLimit := uint32(30)
	worldLevel := p.WorldLevel
	if worldLevel == 0 {
		worldLevel = 6
	}

	var waves [][]*pb.SceneMonster
	var buffs []*pb.BattleBuff

	if s.freesr != nil && s.freesr.BattleConfig != nil {
		cfg := s.freesr.BattleConfig
		stageID = cfg.StageID
		if cfg.CycleCount > 0 {
			roundLimit = cfg.CycleCount
		}
		for wi, wave := range cfg.Monsters {
			_ = wi
			var monsters []*pb.SceneMonster
			for _, m := range wave {
				if m.MonsterID == 0 {
					continue
				}
				hp := uint32(10000)
				monsters = append(monsters, &pb.SceneMonster{
					MonsterId: m.MonsterID,
					MaxHp:     hp,
					CurHp:     hp,
				})
			}
			if len(monsters) > 0 {
				waves = append(waves, monsters)
			}
		}
		for _, b := range cfg.Blessings {
			if b.ID == 0 {
				continue
			}
			buffs = append(buffs, &pb.BattleBuff{
				Id:         b.ID,
				Level:      b.Level,
				OwnerIndex: 0,
				WaveFlag:   0,
			})
		}
	}

	// Battle avatars come from current lineup snapshot.
	var battleAvatars []*pb.BattleAvatar
	if lu := p.Lineups[p.CurrentLineupIndex]; lu != nil {
		for idx, id := range lu.AvatarIDs {
			if id == 0 {
				continue
			}
			battleAvatars = append(battleAvatars, s.buildBattleAvatar(p, id, uint32(idx), worldLevel))
		}
	}
	if len(battleAvatars) == 0 {
		// Fallback: use the current hero to avoid client crashes.
		hero := s.currentTrailblazerAvatarID(p)
		if hero == 0 {
			hero = 8006
		}
		battleAvatars = append(battleAvatars, s.buildBattleAvatar(p, hero, 0, worldLevel))
	}

	// Build monster waves list.
	monsterWaveList := make([]*pb.SceneMonsterWave, 0, len(waves))
	for i, monsters := range waves {
		level := uint32(1)
		for _, w := range waves {
			_ = w
		}
		monsterWaveList = append(monsterWaveList, &pb.SceneMonsterWave{
			MonsterParam:  &pb.SceneMonsterWaveParam{Level: level},
			MonsterList:   monsters,
			BattleStageId: stageID,
			BattleWaveId:  uint32(i + 1),
		})
	}

	// Keep deterministic ordering for buffs so the client stays stable.
	sort.Slice(buffs, func(i, j int) bool {
		if buffs[i].Id == buffs[j].Id {
			return buffs[i].Level < buffs[j].Level
		}
		return buffs[i].Id < buffs[j].Id
	})

	info := &pb.SceneBattleInfo{
		StageId:           stageID,
		BattleId:          stageID,
		MonsterWaveLength: uint32(len(monsterWaveList)),
		LogicRandomSeed:   randomUint32(),
		MonsterWaveList:   monsterWaveList,
		BattleAvatarList:  battleAvatars,
		RoundsLimit:       roundLimit,
		BuffList:          buffs,
		WorldLevel:        worldLevel,
	}
	return info
}

func (s *Server) buildBattleAvatar(p *player.Player, avatarID, index, worldLevel uint32) *pb.BattleAvatar {
	level := uint32(70)
	promotion := uint32(6)
	rank := uint32(6)
	enhancedID := uint32(0)
	curSp := uint32(0)
	maxSp := uint32(100)
	var skillTrees []*pb.AvatarSkillTree

	if s != nil && s.freesr != nil {
		if preset, ok := s.freesr.GetAvatar(avatarID); ok {
			if preset.Level > 0 {
				level = preset.Level
			}
			if preset.Promotion > 0 {
				promotion = preset.Promotion
			}
			if preset.Data.Rank > 0 {
				rank = preset.Data.Rank
			}
			enhancedID = preset.EnhancedID
			if preset.SpMax > 0 {
				maxSp = preset.SpMax
			}
			curSp = preset.SpValue
			if len(preset.Data.Skills) > 0 {
				trees := make([]*pb.AvatarSkillTree, 0, len(preset.Data.Skills))
				for k, v := range preset.Data.Skills {
					pointID := uint32(0)
					// skill ids are stored as strings in freesr preset
					// ignore parse errors (best-effort)
					if parsed, err := parseUint32(k); err == nil {
						pointID = parsed
					}
					if pointID == 0 {
						continue
					}
					trees = append(trees, &pb.AvatarSkillTree{PointId: pointID, Level: uint32(v)})
				}
				skillTrees = trees
			}
		}
	}

	// Basic HP baseline (not real battle sim yet).
	hp := uint32(10000)

	ba := &pb.BattleAvatar{
		AvatarType:    pb.AvatarType_AVATAR_FORMAL_TYPE,
		Id:            avatarID,
		Level:         level,
		Rank:          rank,
		Index:         index,
		SkilltreeList: skillTrees,
		EquipmentList: s.buildBattleEquipmentList(avatarID),
		Hp:            hp,
		Promotion:     promotion,
		RelicList:     s.buildBattleRelicList(avatarID),
		WorldLevel:    worldLevel,
		AssistUid:     0,
		SpBar:         &pb.SpBarInfo{CurSp: curSp, MaxSp: maxSp},
		EnhancedId:    enhancedID,
	}
	return ba
}

func (s *Server) buildBattleEquipmentList(avatarID uint32) []*pb.BattleEquipment {
	if s == nil || s.freesr == nil || avatarID == 0 {
		return nil
	}
	for _, lc := range s.freesr.Lightcones {
		if lc.EquipAvatar != avatarID || lc.ItemID == 0 {
			continue
		}
		return []*pb.BattleEquipment{{
			Id:        lc.ItemID,
			Level:     lc.Level,
			Promotion: lc.Promotion,
			Rank:      lc.Rank,
		}}
	}
	return nil
}

func (s *Server) buildBattleRelicList(avatarID uint32) []*pb.BattleRelic {
	if s == nil || s.freesr == nil || avatarID == 0 || len(s.freesr.Relics) == 0 {
		return nil
	}
	out := make([]*pb.BattleRelic, 0, 6)
	for _, r := range s.freesr.Relics {
		if r.EquipAvatar != avatarID || r.RelicID == 0 {
			continue
		}
		uid := freesrRelicUniqueID(r)
		sub := make([]*pb.RelicAffix, 0, len(r.SubAffixes))
		for _, sa := range r.SubAffixes {
			if sa.SubAffixID == 0 {
				continue
			}
			sub = append(sub, &pb.RelicAffix{AffixId: sa.SubAffixID, Cnt: sa.Count, Step: sa.Step})
		}
		out = append(out, &pb.BattleRelic{
			Id:          r.RelicID,
			Level:       r.Level,
			MainAffixId:  r.MainAffixID,
			SubAffixList: sub,
			UniqueId:     uid,
			SetId:        r.RelicSetID,
			Type:         relicSlotTypeFromTID(r.RelicID),
			Rarity:       5,
		})
	}
	if len(out) == 0 {
		return nil
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UniqueId < out[j].UniqueId })
	return out
}

