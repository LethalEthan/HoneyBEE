package npacket

import (
	"HoneyGO/nbt"
	"HoneyGO/player"

	"github.com/panjf2000/gnet"
)

func (JG *JoinGame_CB) Encode(UUID string, playername string, GM byte, Conn gnet.Conn) {
	JG = &JoinGame_CB{
		EntityID:            int32(player.AssignEID(playername)),
		IsHardcore:          false,
		Gamemode:            GM,
		PreviousGamemode:    -1,
		WorldCount:          1,
		WorldNames:          []Identifier{"minecraft:overworld"},
		DimensionCodec:      nbt.CreateNBTWriter(""),
		Dimension:           nbt.CreateNBTWriter(""),
		WorldName:           "minecraft:overworld",
		HashedSeed:          0,
		MaxPlayers:          20,
		ViewDistance:        12,
		ReducedDebugInfo:    false,
		EnableRespawnScreen: true,
		IsDebug:             false,
		IsFlat:              true,
	}
	PW := CreatePacketWriter(0x26)
	PW.WriteInt(JG.EntityID)
	PW.WriteBoolean(JG.IsHardcore)
	PW.WriteUnsignedByte(JG.Gamemode)
	PW.WriteByte(JG.PreviousGamemode)
	PW.WriteVarInt(JG.WorldCount)
	PW.WriteArrayIdentifier(JG.WorldNames)
	JG.DimensionCodec.CreateCompoundTag("minecraft:dimension_type")
	//JG.DimensionCodec.

}
