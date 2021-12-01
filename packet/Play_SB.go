package packet

import "github.com/google/uuid"

///
///Serverbound
///

type (
	//Play 0x00
	TeleportConfirm_SB struct {
		TeleportID int32
	}

	//Play 0x01
	QueryBlockNBT_SB struct {
		TransactionID int32
		Location      Position
	}

	//Play 0x02
	SetDifficulty_SB struct {
		NewDifficulty byte
	}

	//Play 0x02
	ChatMessage_SB struct {
		Message string
	}

	//Play 0x04
	ClientStatus_SB struct {
		ActionID int32
	}

	ClientSettings_SB struct {
		Locale               string
		ViewDistance         byte
		ChatMode             int32
		ChatColours          bool
		DisplayedSkinParts   byte
		MainHand             int32
		DisableTextFiltering bool
	}

	TabComplete_SB struct {
		TransactionID int32
		Text          string
	}

	ClickWindowButton_SB struct {
		WindowsID byte
		ButtonID  byte
	}

	ClickWindow_SB struct {
		WindowID    byte
		Slot        int16
		Button      byte
		Mode        int32
		Length      int32
		Array       []Slots
		ClickedItem Slot
	}

	//Play 0x09
	CloseWindow_SB struct {
		WindowID byte
	}

	//Play 0x0A
	PluginMessage_SB struct {
		Channel Identifier
		Data    []byte
	}

	//Play 0x0B
	EditBook_SB struct {
		NewBook   Slot
		IsSigning bool
		Hand      int32
	}

	QueryEntityNBT_SB struct {
		TransactionID int32
		EntityID      int32
	}

	InteractEntity_SB struct {
		EntityID int32
		Type     int32
		TargetX  float32
		TargetY  float32
		TargetZ  float32
		Hand     int32
		Sneaking bool
	}

	//Play 0x0E
	GenerateStructure_SB struct {
		Location    Position
		Levels      int32
		KeepJigsaws bool
	}

	KeepAlive_SB struct {
		KeepAlive int64
	}

	LockDifficulty_SB struct {
		Locked bool
	}

	//Play 0x11
	PlayerPosition_SB struct {
		X        float64
		FeetY    float64
		Z        float64
		OnGround bool
	}

	//Play 0x12
	PlayerPositionRotation_SB struct {
		X        float64
		FeetY    float64
		Z        float64
		Yaw      float32
		Pitch    float32
		OnGround bool
	}

	//Play 0x13
	PlayerRotation_SB struct {
		Yaw     float32
		Pitch   float32
		OnGroud bool
	}

	PlayerMovement_SB struct {
		OnGround bool
	}

	VehicleMove_SB struct {
		X     float64
		Y     float64
		Z     float64
		Yaw   float32
		Pitch float32
	}

	SteerBoat_SB struct {
		LeftPaddleTurning  bool
		RightPaddleTurning bool
	}

	PickItem_SB struct {
		Slot int32
	}

	CraftRecipeRequest_SB struct {
		WindowID byte
		Recipe   Identifier
		MakeAll  bool
	}

	PlayerAbilities_SB struct {
		Flags byte
	}

	PlayerDigging_SB struct {
		Status   int32 //0:Start, 1:Cancel, 2:Finish, 3:dropstack, 4:Drop, 5:Shoot/EatF, 6: SwapItem
		Location Position
		Face     byte //0: -Y_B, 1:+Y_T, 2:-Z_N,3:+Z_S, 4:-X_W, 5:+X_E
	}

	EntityAction_SB struct {
		EntityID  int32
		ActionID  int32
		JumpBoost int32
	}

	SteerVehcile_SB struct {
		Sideways float32
		Forward  float32
		Flags    byte
	}

	Pong_SB struct {
		ID int32
	}

	SetRecipeBookState_SB struct {
		BookID       int32
		BookOpen     bool
		FilterActive bool
	}

	NameItem_SB struct {
		ItemName string
	}

	ResourcePackStatus_SB struct {
		Result int32
	}

	AdvancementTab_SB struct {
		Action int32
		TabID  Identifier
	}

	SelectTrade_SB struct {
		SelectedSlot int32
	}

	SetBeaconEffect_SB struct {
		PrimaryEffect   int32
		SecondaryEffect int32
	}

	HeldItemChange_SB struct {
		Slot int16
	}

	//Play 0x26
	UpdateCommandBlock_SB struct {
		Location Position
		Command  string
		Mode     int32
		Flags    byte
	}

	//Play 0x27
	UpdateCommandBlockMinecart_SB struct {
		EntityID    int32
		Command     string
		TrackOutput string
	}

	CreativeInventoryAction_SB struct {
		Slot        int16
		ClickedItem Slot
	}

	UpdateJigsawBlock_SB struct {
		Location   Position
		Name       Identifier
		Target     Identifier
		Pool       Identifier
		FinalState string
		JointType  string
	}

	//Play 0x2A
	UpdateStructureBlock_SB struct {
		Location  Position
		Action    int32
		Mode      int32
		Name      string
		OffsetX   byte
		OffsetY   byte
		OffsetZ   byte
		SizeX     byte
		SizeY     byte
		SizeZ     byte
		Mirror    int32
		Rotation  int32
		Metadata  string
		Integrity string
		Seed      int64
		Flags     byte
	}

	//Play 0x2B
	UpdateSign_SB struct {
		Location Position
		Line1    string
		Line2    string
		Line3    string
		Line4    string
	}

	Animation_SB struct {
		Hand int32 //0: main hand, 1: main hand
	}

	Spectate_SB struct {
		TargetPlayer uuid.UUID
	}

	PlayerBlockPlacement_SB struct {
		Hand            int32
		Location        Position
		Face            int32
		CursorPositionX float32
		CursorPositionY float32
		CursorPositionZ float32
		InsideBlock     bool
	}

	UseItem_SB struct {
		Hand int32 //0: main hand, 1: offhand
	}
)
