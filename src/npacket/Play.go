package npacket

import "jsonstruct"

///
///ClientBound
///

type UUID string
type Identifier int64 //Change me
type NBT int64        //Change me too
type Position int64   //Change me too aswell

//Play_0x00_CB - Spawn Entity
type (
	SpawnEntity_CB struct {
		EntityID   int32
		ObjectUUID UUID
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
		EntityUUID UUID
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
		EntityUUID UUID
		Motive     int32
		Location   int64
		Direction  byte
	}

	//Play_0x04_CB - Spawn Player
	SpawnPlayer_CB struct {
		EntityID   int32
		PlayerUUID UUID
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
		UUID   UUID
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
		Sender   UUID
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

	//Play_0x16_CB -
	SetSlot_CB struct {
		ItemID        int32
		CooldownTicks int32
	}

	//Play_0x17_CB -
	SetCooldown_CB struct {
		Channel Identifier
		Data    []byte
	}

	//Play_0x18_CB -
	PluginMessage_CB struct {
		SoundName     Identifier
		SoundCategory int32
		EffectPosX    int32
		EffectPosY    int32
		EffectPosZ    int32
		Volume        float32
		Pitch         float32
	}

	//Play_0x19_CB -
	NamedSoundEffect_CB struct {
		Reason jsonstruct.ChatComponent
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

	//Play_0x1C_CB -
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

	//Play_0x1D_CB -
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
		PrimaryBitMask      NBT
		HeightMaps          NBT
		BiomeLength         int32
		Biomes              []int32
		Size                int32
		Data                []byte
		NumberBlockEntities int32
		BlockEntities       NBT
	}

	EntityMovement struct {
	}
)
