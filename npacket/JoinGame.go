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
	JG.DimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:dimension_type"))
	List := nbt.CreateListTag("value", nbt.TagCompound)
	TC := nbt.CreateCompoundTagObject("", 8)
	TC.AddTag("", nbt.CreateStringTag("name", "minecraft:overworld"))
	TC.AddTag("", nbt.CreateIntTag("id", 0))
	TC.AddTag("", CreateDimensionTypeRegistry(1, 1, 0, 1, 1, 1, 0, 0, 1.0, 1.0, 0, 256, 64, 12000, "", "minecraft:overworld"))
	List.AddToList(TC)
	JG.DimensionCodec.AddTag(List)
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.CreateCompoundTag("minecraft:worldgen/biome")
	JG.DimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:worldgen/biome"))
	Blist := nbt.CreateListTag("value", nbt.TagCompound)
	_ = Blist
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.Encode()
	//JG.DimensionCodec.
}

func CreateDimensionTypeRegistry(piglin_safe, natural, respawn_anchor_works, has_skylight, bed_works, has_raids, ultrawarm, has_ceiling byte, ambient_light, coordinate_scale float32, min_y, height, logical_height int32, fixed_time int64, infiniburn, effects string) nbt.TCompound {
	TC := nbt.CreateCompoundTagObject("element", 16)
	TC.AddMultipleTags([]interface{}{
		nbt.CreateByteTag("piglin_safe", piglin_safe), nbt.CreateByteTag("natural", natural),
		nbt.CreateFloatTag("ambient_light", ambient_light), nbt.CreateLongTag("fixed_time", fixed_time),
		nbt.CreateStringTag("infiniburn", infiniburn), nbt.CreateByteTag("respawn_anchor_works", respawn_anchor_works),
		nbt.CreateByteTag("has_skylight", has_skylight), nbt.CreateByteTag("bed_works", bed_works),
		nbt.CreateStringTag("effects", effects), nbt.CreateByteTag("has_raids", has_raids), nbt.CreateIntTag("min_y", min_y),
		nbt.CreateIntTag("height", height), nbt.CreateIntTag("logical_height", logical_height),
		nbt.CreateFloatTag("coordinate_scale", coordinate_scale), nbt.CreateByteTag("ultrawarm", ultrawarm),
		nbt.CreateByteTag("has_ceiling", has_ceiling), nbt.TEnd{}})
	return TC
}
