package scene

import (
	"fmt"

	"hyacine-server/internal/gameserver/data"
	pb "hyacine-server/internal/proto/gen"
)

type levelMapInfo struct {
	AreaList []struct {
		MinimapVolume struct {
			Sections []any `json:"Sections"`
		} `json:"MinimapVolume"`
	} `json:"AreaList"`
}

type floorAssets struct {
	UnlockTeleports []uint32
	LightenSections []uint32
	Groups          []*pb.MazeGroup
	Props           []*pb.MazePropState
	PropExtras      []*pb.MazePropExtraState
	Chests          []*pb.ChestInfo
	FloorSavedData  map[string]int32
}

func buildFloorAssets(gd *data.GameData, floorID uint32) (*floorAssets, error) {
	if gd == nil {
		return nil, fmt.Errorf("GameData is nil")
	}
	floor, err := gd.LevelOutputFloor(int(floorID))
	if err != nil {
		return nil, err
	}

	out := &floorAssets{
		FloorSavedData: map[string]int32{},
	}

	// Saved/custom values are polymorphic in 3.8.0 resources; keep nil for now.
	out.FloorSavedData = nil

	// Lighten sections from NavmapConfigPath (MapInfo_*.json).
	if floor.NavmapConfigPath != "" && gd.Loader() != nil {
		var mi levelMapInfo
		if err := gd.Loader().LoadConfigJSON(floor.NavmapConfigPath, &mi); err == nil {
			sectionCount := 0
			for _, a := range mi.AreaList {
				sectionCount += len(a.MinimapVolume.Sections)
			}
			if sectionCount > 0 {
				out.LightenSections = make([]uint32, 0, sectionCount)
				for i := 0; i < sectionCount; i++ {
					out.LightenSections = append(out.LightenSections, uint32(i))
				}
			}
		}
	}

	// Teleport list comes from TeleportConfig (not anchors).
	seenTeleport := map[uint32]struct{}{}
	seenGroup := map[uint32]struct{}{}
	seenProp := map[uint64]struct{}{}

	if gd != nil {
		if tc, err := gd.TeleportConfig(); err == nil && len(tc) > 0 {
			for _, row := range tc {
				if row == nil || row.ID == 0 || row.FloorID != floorID {
					continue
				}
				if _, ok := seenTeleport[row.ID]; ok {
					continue
				}
				seenTeleport[row.ID] = struct{}{}
				out.UnlockTeleports = append(out.UnlockTeleports, row.ID)
			}
		}
	}

	for _, gi := range floor.GroupInstanceList {
		if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
			continue
		}
		group, err := gd.LoadLevelOutputGroupByPath(gi.GroupPath)
		if err != nil || group == nil {
			continue
		}

		groupID := uint32(gi.ID)
		if _, ok := seenGroup[groupID]; !ok {
			seenGroup[groupID] = struct{}{}
			out.Groups = append(out.Groups, &pb.MazeGroup{GroupId: groupID})
		}

		// Use runtime prop instance IDs as MazePropState.config_id, matching SR-CasPS instId usage.
		for _, p := range group.PropList {
			if p.IsDelete || p.ID == 0 {
				continue
			}

			key := (uint64(groupID) << 32) | uint64(uint32(p.ID))
			if _, ok := seenProp[key]; ok {
				continue
			}
			seenProp[key] = struct{}{}

			// Keep consistent with SR-CasPS sample logs: default state=8 for interactive props.
			state := uint32(8)
			out.Props = append(out.Props, &pb.MazePropState{
				GroupId:  groupID,
				ConfigId: uint32(p.ID),
				State:    state,
			})
			out.PropExtras = append(out.PropExtras, &pb.MazePropExtraState{
				GroupId:  groupID,
				ConfigId: uint32(p.ID),
				State:    state,
			})
		}
	}

	if len(out.UnlockTeleports) == 0 && floor.StartAnchorID != 0 {
		out.UnlockTeleports = append(out.UnlockTeleports, uint32(floor.StartAnchorID))
	}

	// Chest summary list (client expects this list to exist for map UI; actual counts can be refined later).
	out.Chests = []*pb.ChestInfo{
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_NORMAL, OpenedNum: 1},
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_PUZZLE, OpenedNum: 1},
		{ChestType: pb.ChestType_MAP_INFO_CHEST_TYPE_CHALLENGE, OpenedNum: 1},
	}

	return out, nil
}

func BuildSceneMapInfo(gd *data.GameData, entryID, floorID uint32) (*pb.SceneMapInfo, error) {
	assets, err := buildFloorAssets(gd, floorID)
	if err != nil {
		return nil, err
	}
	return &pb.SceneMapInfo{
		Retcode:                0,
		EntryId:                entryID,
		CurMapEntryId:          entryID,
		FloorId:                floorID,
		UnlockTeleportList:     assets.UnlockTeleports,
		LightenSectionList:     assets.LightenSections,
		MazeGroupList:          assets.Groups,
		ChestList:              assets.Chests,
		MazePropList:           assets.Props,
		MazePropExtraStateList: assets.PropExtras,
		FloorSavedData:         assets.FloorSavedData,
	}, nil
}

func UnlockTeleportsForEntries(gd *data.GameData, entryIDs []uint32) ([]uint32, error) {
	if gd == nil {
		return nil, fmt.Errorf("GameData is nil")
	}

	entryToFloor := map[uint32]uint32{}
	if me, err := gd.MapEntrance(); err == nil {
		for _, row := range me {
			if row == nil || row.ID == 0 || row.FloorID == 0 {
				continue
			}
			entryToFloor[uint32(row.ID)] = uint32(row.FloorID)
		}
	}

	seen := map[uint32]struct{}{}
	var out []uint32

	tc, _ := gd.TeleportConfig()
	for _, entryID := range entryIDs {
		floorID := entryToFloor[entryID]
		if floorID == 0 {
			continue
		}
		for _, row := range tc {
			if row == nil || row.ID == 0 || row.FloorID != floorID {
				continue
			}
			if _, ok := seen[row.ID]; ok {
				continue
			}
			seen[row.ID] = struct{}{}
			out = append(out, row.ID)
		}
	}

	return out, nil
}
