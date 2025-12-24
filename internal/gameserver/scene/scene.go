package scene

import (
	"fmt"
	"time"

	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/lineup"
	"hyacine-server/internal/gameserver/player"
	pb "hyacine-server/internal/proto/gen"
)

// DefaultScene is a minimal scene snapshot for the client to finish loading.
type DefaultScene struct {
	EntryID uint32
	PlaneID uint32
	FloorID uint32
	WorldID uint32
}

func PickDefault(loader *data.GameData) (DefaultScene, error) {
	// Align with provided persistent default: plane=20313 -> floor=20313001, entry=2031301.
	entry := DefaultScene{
		EntryID: 2031301,
		PlaneID: 20313,
		FloorID: 20313001,
		WorldID: 401,
	}
	if loader == nil {
		return entry, nil
	}
	// Best-effort correction using resources.
	if mp, err := loader.MazePlane(); err == nil {
		if row := mp[entry.PlaneID]; row != nil {
			if row.WorldID != 0 {
				entry.WorldID = uint32(row.WorldID)
			}
			if row.StartFloorID != 0 {
				entry.FloorID = uint32(row.StartFloorID)
			}
		}
	}
	if me, err := loader.MapEntrance(); err == nil {
		// Try to find an entry matching plane+floor.
		for _, row := range me {
			if row == nil {
				continue
			}
			if uint32(row.PlaneID) == entry.PlaneID && uint32(row.FloorID) == entry.FloorID {
				entry.EntryID = uint32(row.ID)
				break
			}
		}
	}
	return entry, nil
}

func BuildSceneInfo(s DefaultScene, uid uint32) *pb.SceneInfo {
	return BuildSceneInfoForPlayer(s, uid, nil)
}

func BuildSceneInfoForPlayer(s DefaultScene, uid uint32, p *player.Player) *pb.SceneInfo {
	return BuildSceneInfoForPlayerWithData(nil, s, uid, p)
}

func BuildSceneInfoForPlayerWithData(gd *data.GameData, s DefaultScene, uid uint32, p *player.Player) *pb.SceneInfo {
	var (
		x int32
		y int32
		z int32
		mapLayer uint32
	)
	if p != nil {
		x, y, z = p.X, p.Y, p.Z
		mapLayer = p.MapLayer
	}

	leaderEntityID := uint32(1)
	leaderAvatarID := uint32(8001)
	var avatarIDs []uint32
	if p != nil {
		lineup.EnsureDefaults(p)
		if lu := p.Lineups[p.CurrentLineupIndex]; lu != nil {
			if int(lu.LeaderSlot) < len(lu.AvatarIDs) {
				if id := lu.AvatarIDs[lu.LeaderSlot]; id != 0 {
					leaderAvatarID = id
				}
			}
			avatarIDs = append([]uint32(nil), lu.AvatarIDs...)
			if len(avatarIDs) == 0 {
				avatarIDs = []uint32{leaderAvatarID}
			}
		}
	}

	sceneInfo := &pb.SceneInfo{
		GameModeType:   1,
		PlaneId:        s.PlaneID,
		WorldId:        s.WorldID,
		FloorId:        s.FloorID,
		EntryId:        s.EntryID,
		LeaderEntityId: leaderEntityID,
		DimensionId:    0,
	}

	// Actor entities for the current lineup (Java assigns entity ids to each avatar in scene).
	nextActorEntityID := uint32(1)
	for _, id := range avatarIDs {
		if id == 0 {
			continue
		}
		entityID := nextActorEntityID
		nextActorEntityID++
		if id == leaderAvatarID {
			leaderEntityID = entityID
			sceneInfo.LeaderEntityId = leaderEntityID
		}

		sceneInfo.EntityList = append(sceneInfo.EntityList, &pb.SceneEntityInfo{
			EntityId: entityID,
			GroupId:  0,
			InstId:   entityID,
			Motion: &pb.MotionInfo{
				Pos: &pb.Vector{X: x, Y: y, Z: z},
				Rot: &pb.Vector{X: 0, Y: 0, Z: 0},
			},
			Entity: &pb.SceneEntityInfo_Actor{
				Actor: &pb.SceneActorInfo{
					AvatarType:   pb.AvatarType_AVATAR_FORMAL_TYPE,
					BaseAvatarId: id,
					Uid:          uid,
					MapLayer:     mapLayer,
				},
			},
		})
	}
	// Fallback if lineup list was empty.
	if len(sceneInfo.EntityList) == 0 {
		sceneInfo.EntityList = append(sceneInfo.EntityList, &pb.SceneEntityInfo{
			EntityId: 1,
			GroupId:  0,
			InstId:   1,
			Motion: &pb.MotionInfo{
				Pos: &pb.Vector{X: x, Y: y, Z: z},
				Rot: &pb.Vector{X: 0, Y: 0, Z: 0},
			},
			Entity: &pb.SceneEntityInfo_Actor{
				Actor: &pb.SceneActorInfo{
					AvatarType:   pb.AvatarType_AVATAR_FORMAL_TYPE,
					BaseAvatarId: leaderAvatarID,
					Uid:          uid,
					MapLayer:     mapLayer,
				},
			},
		})
		sceneInfo.LeaderEntityId = 1
	}

	if gd != nil && s.FloorID != 0 {
		if assets, err := buildFloorAssets(gd, s.FloorID); err == nil && assets != nil {
			sceneInfo.LightenSectionList = assets.LightenSections
			sceneInfo.FloorSavedData = assets.FloorSavedData
		}

		ents, entityGroups, groupStates := BuildRuntimeEntities(gd, s.FloorID, p, time.Now())
		sceneInfo.EntityList = append(sceneInfo.EntityList, ents...)
		sceneInfo.EntityGroupList = entityGroups
		sceneInfo.GroupStateList = groupStates
	}

	return sceneInfo
}

func BuildEnteredScenes(s DefaultScene) []*pb.EnteredSceneInfo {
	return []*pb.EnteredSceneInfo{
		{PlaneId: s.PlaneID, FloorId: s.FloorID},
	}
}

func Validate(s DefaultScene) error {
	if s.PlaneID == 0 || s.FloorID == 0 || s.EntryID == 0 {
		return fmt.Errorf("invalid scene: entry=%d plane=%d floor=%d", s.EntryID, s.PlaneID, s.FloorID)
	}
	return nil
}
