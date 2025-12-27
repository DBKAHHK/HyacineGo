package player

import "time"

const (
	GenderNone  uint32 = 0
	GenderMan   uint32 = 1
	GenderWoman uint32 = 2
)

type Player struct {
	UID                   uint32 `json:"uid"`
	AccountUID            string `json:"account_uid"`
	Nickname              string `json:"nickname"`
	Gender                uint32 `json:"gender"`
	LastSetNicknameTimeMs int64  `json:"last_set_nickname_time_ms"`

	Level      uint32 `json:"level"`
	Exp        uint32 `json:"exp"`
	WorldLevel uint32 `json:"world_level"`

	Stamina           uint32 `json:"stamina"`
	ReserveStamina    uint32 `json:"reserve_stamina"`
	NextRecoverTimeMs int64  `json:"next_recover_time_ms"`

	// Profile / personalization.
	Signature           string   `json:"signature"`
	CurrentHeadIconID   uint32   `json:"current_head_icon_id"`
	UnlockedHeadIconIDs []uint32 `json:"unlocked_head_icon_ids"`

	HeadFrameExpireTime int64  `json:"head_frame_expire_time"`
	HeadFrameItemID     uint32 `json:"head_frame_item_id"`

	CurrentPersonalCardID   uint32   `json:"current_personal_card_id"`
	UnlockedPersonalCardIDs []uint32 `json:"unlocked_personal_card_ids"`

	AssistAvatarIDs  []uint32        `json:"assist_avatar_ids"`
	DisplayAvatars   []DisplayAvatar `json:"display_avatars"`
	GameplayBirthday uint32          `json:"gameplay_birthday"`

	// Phone cosmetics.
	CurChatBubble    uint32   `json:"cur_chat_bubble"`
	CurPhoneTheme    uint32   `json:"cur_phone_theme"`
	CurPhoneCase     uint32   `json:"cur_phone_case"`
	OwnedChatBubbles []uint32 `json:"owned_chat_bubbles"`
	OwnedPhoneThemes []uint32 `json:"owned_phone_themes"`
	OwnedPhoneCases  []uint32 `json:"owned_phone_cases"`

	// Minimal persistent location snapshot (for Scene spawn and map loading).
	EntryID       uint32 `json:"entry_id"`
	PlaneID       uint32 `json:"plane_id"`
	FloorID       uint32 `json:"floor_id"`
	TeleportID    uint32 `json:"teleport_id"`
	X             int32  `json:"x"`
	Y             int32  `json:"y"`
	Z             int32  `json:"z"`
	CalyxX        int32  `json:"calyx_x"`
	CalyxY        int32  `json:"calyx_y"`
	CalyxZ        int32  `json:"calyx_z"`
	MapLayer      uint32 `json:"map_layer"`
	CalyxGroupID  uint32 `json:"calyx_group_id"`
	CalyxInstID   uint32 `json:"calyx_inst_id"`
	CalyxEntityID uint32 `json:"calyx_entity_id"`
	CalyxPropID   uint32 `json:"calyx_prop_id"`

	// Lineup (team) snapshot - minimal persistent implementation.
	CurrentLineupIndex uint32                 `json:"current_lineup_index"`
	Lineups            map[uint32]*LineupData `json:"lineups"`

	// Avatar overrides.
	AvatarEnhancedIDs map[uint32]uint32 `json:"avatar_enhanced_ids"`
	// Current selected multi-path avatar per growth avatar id (e.g. 8001->8006 for Trailblazer, 1001->1224 for March 7th).
	CurrentMultiPath map[uint32]uint32 `json:"current_multi_path"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LineupData struct {
	Index      uint32   `json:"index"`
	Name       string   `json:"name"`
	LeaderSlot uint32   `json:"leader_slot"`
	AvatarIDs  []uint32 `json:"avatar_ids"`
	Mp         uint32   `json:"mp"`
	MaxMp      uint32   `json:"max_mp"`
}

type DisplayAvatar struct {
	AvatarID uint32 `json:"avatar_id"`
	Pos      uint32 `json:"pos"`
}

func New(uid uint32, accountUID string) *Player {
	now := time.Now()
	return &Player{
		UID:        uid,
		AccountUID: accountUID,

		// Keep New() minimal: defaults should come from `configs/defaults/player.json`
		// (via database template cloning) or be filled by ensurePlayerDefaults().
		Lineups:           map[uint32]*LineupData{},
		AvatarEnhancedIDs: map[uint32]uint32{},
		CurrentMultiPath:  map[uint32]uint32{},

		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p *Player) Touch() {
	p.UpdatedAt = time.Now()
}
