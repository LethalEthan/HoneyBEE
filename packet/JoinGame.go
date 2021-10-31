package packet

import (
	"HoneyBEE/nbt"
	"HoneyBEE/player"
)

func (JG *JoinGame_CB) Encode(playername string, GM byte) *PacketWriter {
	JG.EntityID = int32(player.AssignEID(playername))
	JG.IsHardcore = false
	JG.Gamemode = GM
	JG.PreviousGamemode = -1
	JG.WorldCount = 1
	JG.WorldNames = []Identifier{"minecraft:overworld"}
	JG.DimensionCodec = nbt.CreateNBTWriter("")
	JG.Dimension = nbt.CreateNBTWriter("")
	JG.WorldName = "minecraft:overworld"
	JG.HashedSeed = 0
	JG.MaxPlayers = 20
	JG.ViewDistance = 12
	JG.ReducedDebugInfo = false
	JG.EnableRespawnScreen = true
	JG.IsDebug = false
	JG.IsFlat = true
	PW := CreatePacketWriterWithCapacity(0x26, 1024)
	PW.WriteInt(JG.EntityID)
	PW.WriteBoolean(JG.IsHardcore)
	PW.WriteUnsignedByte(JG.Gamemode)
	PW.WriteByte(JG.PreviousGamemode)
	PW.WriteVarInt(JG.WorldCount)
	PW.WriteArrayIdentifier(JG.WorldNames)
	JG.DimensionCodec.AddCompoundTag("minecraft:dimension_type")
	JG.DimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:dimension_type"))
	List := nbt.CreateListTag("value", nbt.TagCompound)
	TC := nbt.CreateCompoundTag("", 8)
	TC.AddTag(nbt.CreateStringTag("name", "minecraft:overworld"))
	TC.AddTag(nbt.CreateIntTag("id", 0))
	TC.AddTag(CreateDimensionTypeRegistry(1, 1, 0, 1, 1, 1, 0, 0, 1.0, 1.0, 0, 256, 64, 12000, "", "minecraft:overworld"))
	TC.PreviousTag = JG.DimensionCodec.CurrentTag
	TC.EndTag()
	List.AddToList(TC)
	JG.DimensionCodec.AddTag(List)
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.AddCompoundTag("minecraft:worldgen/biome")
	JG.DimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:worldgen/biome"))
	Blist := nbt.CreateListTag("value", nbt.TagCompound)
	TC2 := nbt.CreateCompoundTag("", 8)
	TC2.AddTag(nbt.CreateStringTag("name", "minecraft:plains"))
	TC2.AddTag(nbt.CreateIntTag("id", 0))
	TC2.AddTag(CreateBiomeProperties(1.0, 1.0, 1.0, 1.0, 8364543, 8364543, 8364543, 8364543, 8364543, 8364543, "none", "plains", "", "", ""))
	TC2.PreviousTag = JG.DimensionCodec.CurrentTag
	TC2.EndTag()
	Blist.AddToList(TC2)
	JG.DimensionCodec.AddTag(Blist)
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.EndCompoundTag()
	JG.DimensionCodec.Encode()
	PW.WriteArray(JG.DimensionCodec.Data)
	TC3 := CreateDimensionTypeRegistry(1, 1, 0, 1, 1, 1, 0, 0, 1.0, 1.0, 0, 256, 64, 12000, "", "minecraft:overworld")
	JG.Dimension.AddTag(TC3)
	JG.Dimension.EndCompoundTag()
	JG.Dimension.Encode()
	PW.WriteArray(JG.DimensionCodec.Data)
	PW.WriteArray(JG.Dimension.Data)
	PW.WriteIdentifier(JG.WorldName)
	PW.WriteLong(0)
	PW.WriteVarInt(20)
	PW.WriteVarInt(12)
	PW.WriteBoolean(false)
	PW.WriteBoolean(true)
	PW.WriteBoolean(true)
	PW.WriteBoolean(true)
	Log.Debug("CREATED JOIN GAME")
	return PW
	// f, err := os.Create("testing.nbt")
	// if err != nil {
	// 	panic(err)
	// }
	// f.Write(JG.DimensionCodec.Data)
	// f.Close()
	// utils.PrintHexFromBytes("DimensionCodec", JG.DimensionCodec.Data)
	// Log.Info("DimensionCodec", JG.DimensionCodec.Data)
}

func CreateDimensionTypeRegistry(piglin_safe, natural, respawn_anchor_works, has_skylight, bed_works, has_raids, ultrawarm, has_ceiling byte, ambient_light, coordinate_scale float32, min_y, height, logical_height int32, fixed_time int64, infiniburn, effects string) nbt.TCompound {
	TC := nbt.CreateCompoundTag("element", 16)
	TC.AddMultipleTags([]interface{}{
		nbt.CreateByteTag("piglin_safe", piglin_safe), nbt.CreateByteTag("natural", natural),
		nbt.CreateFloatTag("ambient_light", ambient_light), nbt.CreateLongTag("fixed_time", fixed_time),
		nbt.CreateStringTag("infiniburn", infiniburn), nbt.CreateByteTag("respawn_anchor_works", respawn_anchor_works),
		nbt.CreateByteTag("has_skylight", has_skylight), nbt.CreateByteTag("bed_works", bed_works),
		nbt.CreateStringTag("effects", effects), nbt.CreateByteTag("has_raids", has_raids), nbt.CreateIntTag("min_y", min_y),
		nbt.CreateIntTag("height", height), nbt.CreateIntTag("logical_height", logical_height),
		nbt.CreateFloatTag("coordinate_scale", coordinate_scale), nbt.CreateByteTag("ultrawarm", ultrawarm),
		nbt.CreateByteTag("has_ceiling", has_ceiling)})
	TC.EndTag()
	return TC
}

func CreateBiomeProperties(depth, temperature, scale, downfall float32, sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour int32, precipitation, category, temperature_modifier, grass_colour_modifier, ambient_sound string) nbt.TCompound {
	TC := nbt.CreateCompoundTag("element", 8)
	TC.AddMultipleTags([]interface{}{
		nbt.CreateStringTag("precipitation", precipitation), nbt.CreateFloatTag("depth", depth),
		nbt.CreateFloatTag("temperature", temperature), nbt.CreateFloatTag("scale", scale),
		nbt.CreateFloatTag("downfall", downfall), nbt.CreateStringTag("category", category),
		CreateBiomeEffects(sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour, grass_colour_modifier)})
	TC.EndTag()
	return TC
}

func CreateBiomeEffects(sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour int32, grass_colour_modifier string) nbt.TCompound {
	TC := nbt.CreateCompoundTag("effects", 10)
	TC.AddMultipleTags([]interface{}{
		nbt.CreateIntTag("sky_color", sky_colour), nbt.CreateIntTag("water_fog_color", water_fog_colour),
		nbt.CreateIntTag("fog_color", fog_colour), nbt.CreateIntTag("water_color", water_colour),
		nbt.CreateIntTag("foliage_color", foliage_colour), nbt.CreateIntTag("grass_color", grass_colour),
		nbt.CreateStringTag("grass_color_modifier", grass_colour_modifier)})
	TC.EndTag()
	return TC
}
