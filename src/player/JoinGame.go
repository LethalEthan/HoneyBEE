package Player

import (
	"packet"
)

const (
	Nether    = -1
	Overworld = 0
	End       = 1
)

var Dimension = Overworld

//GameJoin - Structure of the JoinGame packet
type GameJoin struct {
	EntityID            int32  //Players EntityID
	GameMode            uint8  //0: Survival, 1: Creative, 2: Adventure, 3: Spectator. Bit 3 (0x8) is the hardcore flag.
	Dimension           int    //See connstants above
	HashedSeed          int64  //First 8 bytes of the SHA-256 hash of world seed
	MaxPlayers          uint8  //Used to be used but according to wiki.vg it's no longer used
	LevelType           string //Max 16 length: default, flat, largeBiomes etc
	ViewDistance        VarInt //RenderDistance (2-32)
	ReducedDebugInfo    bool
	EnableRespawnScreen bool //Set false when doImmediateRespawn gamerule is true
}

//Temporary
func CreateGameJoin(*GameJoin) *GameJoin {
	GJ := new(GameJoin)
	GJ.EntityID = 0
	GJ.GameMode = 1
	GJ.Dimension = 0
	GJ.HashedSeed = 0
	GJ.MaxPlayers = 20
	GJ.LevelType = "default"
	GJ.ViewDistance = 16
	GJ.ReducedDebugInfo = false
	GJ.EnableRespawnScreen = true
	writer := Packet.CreatePacketWriter(0x26)
	writer.WriteInt(GJ.EntityID)
	return GJ
}
