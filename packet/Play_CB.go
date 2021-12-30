package packet

import (
	"HoneyBEE/jsonstruct"

	"github.com/google/uuid"
)

///
///ClientBound
///

type Identifier string //Change me
type NBT int64         //Change me too
type Position int64    //Change me too aswell
type Angle byte        //Change me too aswell aswell

//Play_0x00_CB - Spawn Entity
type (
	SpawnEntity_CB struct {
		EntityID   int32
		ObjectUUID uuid.UUID
		Type       int32
		X          float64
		Y          float64
		Z          float64
		Pitch      byte
		Yaw        byte
		Data       int
		VelocityX  int16
		VelocityY  int16
		VelocityZ  int16
	}

	//Play_0x01_CB - Spawn Experience Orb
	SpawnExperienceOrb_CB struct {
		EntityID int32
		X        float64
		Y        float64
		Z        float64
		Count    int16
	}

	//Play_0x02_CB - Spawn Living Entity
	SpawnLivingEntity_CB struct {
		EntityID   int32
		EntityUUID uuid.UUID
		Type       int32
		X          float64
		Y          float64
		Z          float64
		Yaw        byte
		Pitch      byte
		HeadPitch  byte
		VelocityX  int16
		VelocityY  int16
		VelocityZ  int16
	}

	//Play_0x03_CB - Spawn Painting
	SpawnPainting_CB struct {
		EntityID   int32
		EntityUUID uuid.UUID
		Motive     int32
		Location   int64
		Direction  byte
	}

	//Play_0x04_CB - Spawn Player
	SpawnPlayer_CB struct {
		EntityID   int32
		PlayerUUID uuid.UUID
		X          float64
		Y          float64
		Z          float64
		Yaw        byte
		Pitch      byte
	}

	//Play_0x05_CB -
	SculkVibrationSignal_CB struct {
		SourcePosition Position
		Destination    Identifier
		Destnation     interface{}
		ArrivalTicks   int32
	}

	//Play_0x06_CB - Entity Animation
	EntityAnimation_CB struct {
		EntityID  int32
		Animation byte
	}

	//Play_0x07_CB - Statistics
	Statistics_CB struct {
		Count      int32
		Statistics Statistic
	}

	//Play_0x08_CB - Acknowledge Player Digging
	AckowledgePlayerDigging_CB struct {
		Location   int64
		Block      int32
		Status     int32
		Successful bool
	}

	//Play0x09_CB - Block Break Animation
	BlockBreakAnimation_CB struct {
		EntityID     int32
		Location     int64
		DestroyStage byte
	}

	//Play_0x0A_CB - Block Entity Data
	BlockEntityData_CB struct {
		Location int64
		Action   byte
		NBT      []byte
	}

	//Play_0x0B_CB - Block Action
	BlockAction_CB struct {
		Location    int64
		ActionID    byte
		ActionParam byte
		BlockType   int32
	}

	//Play_0x0C_CB - Block Change
	BlockChange_CB struct {
		Location int64
		BlockID  int32
	}

	//Play_0x0D_CB - Boss Bar
	BossBar_CB struct {
		UUID   uuid.UUID
		Action int32 //Const
		//Finish ME
	}

	//Play_0x0E_CB - Server Difficulty
	ServerDifficulty struct {
		Difficulty     byte
		DifficultyLock bool
	}

	//Play_0x0F_CB - Chat MessageID
	ChatMessage_CB struct {
		Chat     *jsonstruct.ChatComponent
		Position byte
		Sender   uuid.UUID
	}

	//Play 0x10
	ClearTitle_CB struct {
		Reset bool
	}

	//Play_0x12_CB - Declare Commands
	DeclareCommands_CB struct {
		Count     int32
		Nodes     []NodeFormat
		RootIndex int32
	}

	//Play_0x13_CB - Close Window
	CloseWindow_CB struct {
		WindowID byte
	}

	//Play_0x14_CB - Window Items
	WindowItems_CB struct {
		WindowID byte
		Count    int16
		SlotData Slot
	}

	//Play_0x15_CB - Window Items
	WindowProperty_CB struct {
		WindowID byte
		Property int16
		Value    []Slot
	}

	//Play_0x16
	SetSlot_CB struct {
		WindowsID byte
		StateID   int32
		Slot      int16
		SlotData  Slot
	}

	//Play_0x17
	SetCooldown_CB struct {
		ItemID        int32
		CooldownTicks int32
	}

	//Play_0x18
	PluginMessage_CB struct {
		Channel Identifier
		Data    []byte
	}

	//Play_0x19
	NamedSoundEffect_CB struct {
		SoundName     Identifier
		SoundCategory int32
		EffectPosX    int32
		EffectPosY    int32
		EffectPosZ    int32
		Volume        float32
		Pitch         float32
	}

	//Play 0x1A
	Disconnect_CB struct {
		Reason jsonstruct.ChatComponent
	}

	//Play_0x1B_CB - Entity Status
	EntityStatus_CB struct {
		EntityID     int32
		EntityStatus byte
	}

	//Play_0x1C
	Explosion_CB struct {
		X             float32
		Y             float32
		Z             float32
		Strength      float32
		RecordCount   int
		Records       [][][]byte
		PlayerMotionX float32
		PlayerMotionY float32
		PlayerMotionZ float32
	}

	//Play_0x1D
	UnloadChunk_CB struct {
		ChunkX int
		ChunkZ int
	}

	//Play 0x1E
	ChangeGameState_CB struct {
		Reason byte
		Value  float32
	}

	//Play 0x1F
	OpenHorseWindow_CB struct {
		WindowID      byte
		NumberOfSlots int32
		EntityID      int32
	}
	//Play 0x20
	InitialiseWorldBorder_CB struct {
		X                      float64
		Z                      float64
		OldDiameter            float64
		NewDiameter            float64
		Speed                  int64
		PortalTeleportBoundary int32
		WarningBlocks          int32
		WarningTime            int32
	}

	//Play 0x21
	KeepAlive_CB struct {
		KeepAliveID int64
	}

	//Play 0x22
	ChunkData_CB struct {
		ChunkX              int32
		ChunkZ              int32
		BitMaskLength       int32
		PrimaryBitMask      []int64
		HeightMaps          []byte
		BiomeLength         int32
		Biomes              []int32
		Size                int
		Data                []byte
		NumberBlockEntities int
		BlockEntities       NBT
	}

	//Play 0x23
	Effect_CB struct {
		EffectID              int32
		Location              Position
		Data                  int32
		DisableRelativeVolume bool
	}

	//Play 0x24
	Particle_CB struct {
		ParticleID   int32
		LongDistance bool
		X            float64
		Y            float64
		Z            float64
		OffsetX      float32
		OffsetY      float32
		OffsetZ      float32
		ParticleData float32
		Particle     int32
		Data         interface{}
	}

	//Play 0x25
	UpdateLight_CB struct {
		ChunkX     int32
		ChunkZ     int32
		TrustEdges bool
		//
		SkyLightMaskLength int32
		SkyLightMask       []int64
		//
		BlockLightMaskLength int32
		BlockLightMask       []int64
		//
		EmptyBlockLightMaskLength int32
		EmptyBlockLightMask       []int64
		//
		SkyLightArrayCount int32
		SkyLightArrays     struct {
			Length        int32
			SkyLightArray [2048]byte //Use go-nibble
		}
		BlockLightArrayCount struct {
			BlockLightArrayCount int32
			BlockLightArrays     [2048]byte //1 array for each bit set to true in block light mask
		}
	}

	//Play 0x26
	JoinGame_CB struct {
		EntityID            int32
		IsHardcore          bool
		Gamemode            byte
		PreviousGamemode    int8
		WorldCount          int32
		WorldNames          []Identifier
		DimensionCodec      []byte
		Dimension           []byte
		WorldName           Identifier
		HashedSeed          int64
		MaxPlayers          int32
		ViewDistance        int32
		SimulationDistance  int32
		ReducedDebugInfo    bool
		EnableRespawnScreen bool
		IsDebug             bool
		IsFlat              bool
	}

	//Play 0x27
	MapData_CB struct {
		MapID            int32
		Scale            byte
		Locked           bool
		TrackingPosition bool
		IconCount        int32
		Icons            []Icon
		Columns          byte
		Rows             byte
		X                byte
		Z                byte
		Length           int32
		Data             []byte
	}

	//Play 0x28
	TradeList_CB struct {
		WindowID          int32
		Size              byte
		Trade             []Trades
		VillagerLevel     int32
		Experience        int32
		IsRegularVillager bool
		CanRestock        bool
	}

	//Play 0x29
	EntityPosition_CB struct {
		EnitityID int32
		DeltaX    int16
		DeltaY    int16
		DeltaZ    int16
		OnGround  bool
	}

	//Play 0x2A
	EntityPostionRotation_CB struct {
		EP    EntityPosition_CB
		Yaw   Angle
		Pitch Angle
	}

	//Play 0x2B
	EntityRotation_CB struct {
		EntityID int32
		Yaw      Angle
		Pitch    Angle
		OnGround bool
	}

	//Play 0x2C
	VehicleMove_CB struct {
		X     float64
		Y     float64
		Z     float64
		Yaw   float32
		Pitch float32
	}

	//Play 0x2D
	OpenBook_CB struct {
		Hand int32 //0: Main Hand, 1: Off Hand
	}

	//Play 0x2E
	OpenWindows_CB struct {
		WindowID    int32
		WindowType  int32
		WindowTitle jsonstruct.ChatComponent
	}
	//Play 0x2F
	OpenSignEditor_CB struct {
		Location Position
	}

	//Play 0x30
	Play_Ping_CB struct {
		ID int32
	}

	//Play 0x31
	CraftRecipeResponse_CB struct {
		WindowID byte
		Recipe   Identifier
	}

	//Play 0x32
	PlayerAbilities_CB struct {
		Flags       byte //0x01: Invulnerable / 0x02: Flying / 0x04: Allow Flying / 0x08: CreativeMode
		FlyingSpeed float32
		FOVModifier float32
	}

	//Play 0x33
	EndCombatEvent_CB struct {
		Duration int32
		EntityID int32
	}

	//Play 0x34 UNUSED
	EnterCombatEvent_CB struct{}

	//Play 0xx35
	DeathCombatEvent_CB struct {
		PlayerID int32
		EntityID int32
		Message  jsonstruct.ChatComponent
	}

	//Play 0x36
	PlayerInfo_CB struct {
		Action        int32
		NumberPlayers int32
		Player        struct {
			UUID     uuid.UUID
			ActionID byte
			Action   struct{}
		}
	}

	//Play 0x37
	FacePlayer_CB struct {
		FeetEyes       int32
		TargetX        float64
		TargetY        float64
		TargetZ        float64
		IsEntity       bool
		EntityID       int32
		EntityFeetEyes int32
	}

	//Play 0x38
	PlayerPositionLook struct {
		X               float64
		Y               float64
		Z               float64
		Yaw             float32
		Pitch           float32
		Flags           byte //0x01: X / 0x02: Y / 0x04: Z / 0x08: Y_ROT / 0x10: X_ROT If set value is relative not absolute
		TeleportID      int32
		DismountVehicle bool
	}

	//Play 0x39
	UnlockRecipes_CB struct {
		Action                   int32 //0: init / 1: add / 2: remove
		CraftingRecipeBookOpen   bool
		CraftingRecipeBookFilter bool
		//
		SmeltingRecipeBookOpen   bool
		SmeltingRecipeBookFilter bool
		//
		BlastFurnaceRecipeBookOpen   bool
		BlastFurnaceRecipeBookFilter bool
		//
		SmokerRecipeBookOpen   bool
		SmokerRecipeBookFilter bool
		//
		ArraySize1 int32
		RecipeID1  []Identifier
		ArraySize2 int32
		RecipeID2  []Identifier
	}

	//Play 0x3A
	DestroyEntity_CB struct {
		EntityID int32
	}

	//Play 0x3B
	RemoveEntityEffect_CB struct {
		EntityID int32
		EffectID byte
	}

	//Play 0x3C
	ResourcePackSend_CB struct {
		URL           string
		Hash          string
		Forced        bool
		PromptMessage jsonstruct.ChatComponent
	}

	//Play 0x3D
	Respawn_CB struct {
		Dimension    NBT
		WorldName    Identifier
		HashedSeed   int64
		Gamemode     byte
		Previous     byte
		Bebug        bool
		Flat         bool
		CopyMetadata bool
	}

	//Play 0x3E
	EntityHeadLook_CB struct {
		EntityID int32
		HeadYaw  Angle
	}

	//Play 0x3F
	MultiBlockChange_CB struct {
		ChunkSectionPosition int64
		AnalSomething        bool //suggested by Kacper
		BlocksArraySize      int32
		Blocks               []int64
	}

	//Play 0x40
	SelectAdvancementTab_CB struct {
		HasID              bool
		OptionalIdentifier string
	}

	//Play 0x41
	ActionBar_CB struct {
		ActionBarText jsonstruct.ChatComponent
	}

	//Play 0x42
	WorldBorderCenter_CB struct {
		X float64
		Z float64
	}

	//Play 0x43
	WorldBorderLerpSize_CB struct {
		OldDiameter float64
		NewDiameter float64
		Speed       int64
	}

	//Play 0x44
	WorldBorderSize_CB struct {
		Diameter float64
	}

	//Play 0x45
	WorldBorderWarningDelay_CB struct {
		WarningTime int32
	}

	//Play 0x46
	WorldBorderWarningReach_CB struct {
		WarningBlock int32
	}

	//Play 0x47
	Camera_CB struct {
		CameraID int32
	}

	//Play 0x48
	HeldItemChange_CB struct {
		Slot byte
	}

	//Play 0x49
	UpdateViewPosition_CB struct {
		ChunkX int32
		ChunkZ int32
	}

	//Play 0x4A
	UpdateViewDistance_CB struct {
		ViewDistance int32
	}

	//Play 0x4B
	SpawnPosition_CB struct {
		Location Position
		Angle    float32
	}

	//Play 0x4C
	DisplayScoreboard_CB struct {
		Position  byte
		ScoreName string
	}

	//Play 0x4D
	EntityMetadata_CB struct {
		EntityID int32
		Metadata MetadataFormat
	}

	//Play 0x4E
	AttachEntity_CB struct {
		AttachedEntityID int32
		HoldingEntityID  int32
	}

	//Play 0x4F
	EntityVelocity_CB struct {
		EntityID  int32
		VelocityX int16
		VelocityY int16
		VelocityZ int16
	}

	//Play 0x50
	EntityEquipment_CB struct {
		EntityID  int32
		Equipment struct {
			Slot []byte //0:MainHead, 1:OffHand, 2-5:ArmorSlot
			Item []Slot
		}
	}

	//Play 0x51
	SetExperience_CB struct {
		Experiencebar   float32
		Level           int32
		TotalExperience int32
	}

	//Play 0x52
	UpdateHealth_CB struct {
		Health         float32
		Food           int32
		FoodSaturation float32
	}

	//Play 0x53
	ScoreboardObjective_CB struct {
		ObjectiveName  string
		Mode           byte
		ObjectiveValue jsonstruct.ChatComponent
		Type           int32
	}

	//Play 0x54
	SetPassengers_CB struct {
		EntityID       int32
		PassengerCount int32
		Passengers     []int32
	}

	//0x55 moved to teams.go

	//Play 0x56
	UpdateScore_CB struct {
		EntityName   string
		Action       byte
		ObectiveName string
		Value        int32
	}

	//Play 0x57
	SetTitleSubTitle_CB struct {
		SubtitleText jsonstruct.ChatComponent
	}

	//Play 0x58
	TimeUpdate_CB struct {
		WorldAge  int64
		TimeOfDay int64
	}

	//Ply 0x59
	SetTitleText_CB struct {
		TitleText jsonstruct.ChatComponent
	}

	//Play 0x5A
	SetTitleTimes_CB struct {
		FadeIn  int32
		Stay    int32
		FadeOut int32
	}

	//Play 0x5B
	EntitySoundEffect_CB struct {
		SoundID       int32
		SoundCategory int32
		EntityID      int32
		Volume        float32
		Pitch         float32
	}

	//Play 0x5C
	SoundEffect_CB struct {
		SoundID         int32
		SoundCategory   int32
		EffectPositionX int32
		EffectPositionY int32
		EffectPositionZ int32
		Volume          float32
		Pitch           float32
	}

	//Play 0x5D
	StopSound_CB struct {
		Flags  byte
		Source int32
		Sound  Identifier
	}

	//Play 0x5E
	PlayerListHeaderFooter_CB struct {
		Header jsonstruct.ChatComponent
		Footer jsonstruct.ChatComponent
	}

	//Play 0x5F
	NBTQueryResponse_CB struct {
		TransactionID int32
		NBT           NBT
	}

	//Play 0x60
	CollectItem_CB struct {
		CollectedEntityID int32
		CollectorEntityID int32
		PickupItemCount   int32
	}

	//Play 0x61
	EntityTeleport_CB struct {
		EntityID int32
		X        float64
		Y        float64
		Z        float64
		Yaw      Angle
		Pitch    Angle
		OnGround bool
	}

	//Play 0x62
	Advancements_CB struct {
		ResetClear         bool
		MappingSize        int32
		AdvancementMapping struct {
			Key   []Identifier
			Value []Advancement
		}
		ListSize        int32
		Identifiers     []Identifier
		ProgressSize    int32
		ProgressMapping struct {
			Key   []Identifier
			Value []Advancement
		}
	}

	//Play 0x63
	EntityProperties_CB struct {
		EntityID           int32
		NumberOfProperties int32
		Properties         []Property
	}

	//Play 0x64
	EntityEffect_CB struct {
		EntityID  int32
		EffectID  byte
		Amplifier byte
		Duration  int32
		Flags     byte
	}

	DeclareRecipes_CB struct {
		NumRecipes int32
		Recipe     Recipes
	}

	//Play 0x66
	Tags_CB struct {
		Length int32
		Tags   TagsArray
	}
)
