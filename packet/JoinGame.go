package packet

import (
	"HoneyBEE/biome"
	"HoneyBEE/config"
	"HoneyBEE/nbt"
	"os"
)

var Overworld_Entry nbt.Compound
var Overworld_Caves_Entry nbt.Compound
var Nether_Entry nbt.Compound
var End_Entry nbt.Compound
var BiomeEntry nbt.List

func CreateEntries() {
	// Create Overworld entry
	Overworld_Entry = nbt.CreateCompoundTag("")
	Overworld_Entry.AddTag(nbt.CreateStringTag("name", "minecraft:overworld"))
	Overworld_Entry.AddTag(nbt.CreateIntTag("id", 0))
	Overworld_Entry.AddTag(CreateDimensionTypeRegistry(0, 1, 0, 1, 1, 1, 0, 0, 0, 1.0, 0, 256, 256, 12000, "minecraft:infiniburn_overworld", "minecraft:overworld")) // element
	Overworld_Entry.EndTag()
	// Create Overworld_caves entry
	Overworld_Caves_Entry = nbt.CreateCompoundTag("")
	Overworld_Caves_Entry.AddTag(nbt.CreateStringTag("name", "minecraft:overworld_caves"))
	Overworld_Caves_Entry.AddTag(nbt.CreateIntTag("id", 1))
	Overworld_Caves_Entry.AddTag(CreateDimensionTypeRegistry(0, 1, 0, 1, 1, 1, 0, 1, 0.0, 1.0, 0, 256, 256, 12000, "minecraft:infiniburn_overworld", "minecraft:overworld")) // element
	Overworld_Caves_Entry.EndTag()
	// Create Nether Entry
	Nether_Entry = nbt.CreateCompoundTag("")
	Nether_Entry.AddTag(nbt.CreateStringTag("name", "minecraft:the_nether"))
	Nether_Entry.AddTag(nbt.CreateIntTag("id", 2))
	Nether_Entry.AddTag(CreateDimensionTypeRegistry(1, 0, 1, 0, 0, 0, 1, 1, 0.1, 8, 0, 256, 128, 18000, "minecraft:infiniburn_nether", "minecraft:the_nether")) // element
	Nether_Entry.EndTag()
	// Create End entry
	End_Entry = nbt.CreateCompoundTag("")
	End_Entry.AddTag(nbt.CreateStringTag("name", "minecraft:the_end"))
	End_Entry.AddTag(nbt.CreateIntTag("id", 3))
	End_Entry.AddTag(CreateDimensionTypeRegistry(0, 0, 0, 0, 0, 1, 0, 0, 0, 1.0, 0, 256, 256, 6000, "minecraft:infiniburn_end", "minecraft:the_end")) // element
	End_Entry.EndTag()
	BiomeEntry = CreateBiomeRegistries()
	// TestBiomeList()
}

func (JG *JoinGame_CB) Encode(PW *PacketWriter) {
	PW.ResetData(0x26)
	// Create Encoders
	JGDimensionCodec := nbt.CreateNBTEncoder()
	JGDimension := nbt.CreateNBTEncoder()
	PW.WriteInt(JG.EntityID)
	PW.WriteBoolean(JG.IsHardcore)
	PW.WriteUByte(JG.Gamemode)
	PW.WriteByte(JG.PreviousGamemode)
	PW.WriteVarInt(JG.WorldCount)
	PW.WriteArrayIdentifier(JG.WorldNames)
	// Create Dimension type
	JGDimensionCodec.AddCompoundTag("minecraft:dimension_type")                      // Create dimension type
	JGDimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:dimension_type")) // Create string tag
	// Create list of values for overworld
	List := nbt.CreateListTag("value", nbt.TagCompound)
	List.AddTag(Overworld_Entry)                                // Add the entry to the list
	List.AddTag(Overworld_Caves_Entry)                          // Add the entry to the list
	List.AddTag(Nether_Entry)                                   // Add the entry to the list
	List.AddTag(End_Entry)                                      // Add the entry to the list
	JGDimensionCodec.AddTag(List)                               // Add the list to JGDC
	JGDimensionCodec.EndCompoundTag()                           // Add end tag to first compound tag in dimension_type - overworld
	JGDimensionCodec.AddCompoundTag("minecraft:worldgen/biome") // Create biome registry
	JGDimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:worldgen/biome"))
	// Blist.AddTag(CreateBiomeRegistry("minecraft:void")                 // Add compound to List
	JGDimensionCodec.AddTag(BiomeEntry) // add list to JGC
	JGDimensionCodec.EndCompoundTag()   //End biome registry
	JGDimensionCodec.EndCompoundTag()   // End root compound
	// TC3 := CreateDimensionTypeRegistry(1, 1, 0, 1, 1, 1, 0, 0, 0, 1.0, 0, 256, 256, 12780, "minecraft:infiniburn_overworld", "minecraft:overworld")
	JGDimension.AddMultipleTags([]interface{}{
		nbt.CreateByteTag("piglin_safe", 0), nbt.CreateByteTag("natural", 1),
		nbt.CreateFloatTag("ambient_light", 0), nbt.CreateLongTag("fixed_time", 12780),
		nbt.CreateStringTag("infiniburn", "minecraft:infiniburn_overworld"), nbt.CreateByteTag("respawn_anchor_works", 1),
		nbt.CreateByteTag("has_skylight", 1), nbt.CreateByteTag("bed_works", 1),
		nbt.CreateStringTag("effects", "minecraft:overworld"), nbt.CreateByteTag("has_raids", 1), nbt.CreateIntTag("min_y", 0),
		nbt.CreateIntTag("height", 256), nbt.CreateIntTag("logical_height", 256),
		nbt.CreateDoubleTag("coordinate_scale", 1.0), nbt.CreateByteTag("ultrawarm", 0),
		nbt.CreateByteTag("has_ceiling", 0)})
	JGDimension.EndCompoundTag() // End root compound
	JG.DimensionCodec = JGDimensionCodec.Encode()
	JG.Dimension = JGDimension.Encode()
	Log.Debug(JG.Dimension)
	PW.WriteArray(JG.DimensionCodec)
	PW.WriteArray(JG.Dimension) //JG.Dimension)
	PW.WriteIdentifier(JG.WorldName)
	PW.WriteLong(0)
	PW.WriteVarInt(20)
	PW.WriteVarInt(int32(config.GConfig.Performance.ViewDistance))
	// PW.WriteVarInt(int32(config.GConfig.Performance.SimulationDistance)) // 1.18
	PW.WriteBoolean(false)
	PW.WriteBoolean(true)
	PW.WriteBoolean(false)
	PW.WriteBoolean(false)
	// Log.Debug("CREATED JOIN GAME")
	// Debug - save results
	// f, err := os.Create("testing.nbt")
	// if err != nil {
	// 	panic(err)
	// }
	// f.Write(JGDimensionCodec.GetData())
	// f.Close()
	// utils.PrintHexFromBytes("DimensionCodec", JGDimensionCodec.GetData())
	// Log.Info("DimensionCodec", JGDimensionCodec.GetData())
}

func CreateDimensionTypeRegistry(piglin_safe, natural, respawn_anchor_works, has_skylight, bed_works, has_raids, ultrawarm, has_ceiling byte, ambient_light float32, coordinate_scale float64, min_y, height, logical_height int32, fixed_time int64, infiniburn, effects string) nbt.Compound {
	TC := nbt.CreateCompoundTag("element")
	TC.AddMultipleTags([]interface{}{
		nbt.CreateByteTag("piglin_safe", piglin_safe), nbt.CreateByteTag("natural", natural),
		nbt.CreateFloatTag("ambient_light", ambient_light), nbt.CreateLongTag("fixed_time", fixed_time),
		nbt.CreateStringTag("infiniburn", infiniburn), nbt.CreateByteTag("respawn_anchor_works", respawn_anchor_works),
		nbt.CreateByteTag("has_skylight", has_skylight), nbt.CreateByteTag("bed_works", bed_works),
		nbt.CreateStringTag("effects", effects), nbt.CreateByteTag("has_raids", has_raids), nbt.CreateIntTag("min_y", min_y),
		nbt.CreateIntTag("height", height), nbt.CreateIntTag("logical_height", logical_height),
		nbt.CreateDoubleTag("coordinate_scale", coordinate_scale), nbt.CreateByteTag("ultrawarm", ultrawarm),
		nbt.CreateByteTag("has_ceiling", has_ceiling)})
	TC.EndTag()
	return TC
}

func CreateBiomeProperties(depth, temperature, scale, downfall float32, sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour int32, precipitation, category, temperature_modifier, grass_colour_modifier, ambient_sound string) nbt.Compound {
	TC := nbt.CreateCompoundTag("element")
	TC.AddMultipleTags([]interface{}{
		nbt.CreateStringTag("precipitation", precipitation), nbt.CreateFloatTag("depth", depth),
		nbt.CreateFloatTag("temperature", temperature), nbt.CreateFloatTag("scale", scale),
		nbt.CreateFloatTag("downfall", downfall), nbt.CreateStringTag("category", category),
		CreateBiomeEffects(sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour, grass_colour_modifier)})
	TC.EndTag()
	return TC
}

func CreateBiomeEffects(sky_colour, water_fog_colour, fog_colour, water_colour, foliage_colour, grass_colour int32, grass_colour_modifier string) nbt.Compound {
	TC := nbt.CreateCompoundTag("effects")
	TC.AddMultipleTags([]interface{}{
		nbt.CreateIntTag("sky_color", sky_colour), nbt.CreateIntTag("water_fog_color", water_fog_colour),
		nbt.CreateIntTag("fog_color", fog_colour), nbt.CreateIntTag("water_color", water_colour),
		nbt.CreateIntTag("foliage_color", foliage_colour), nbt.CreateIntTag("grass_color", grass_colour),
		nbt.CreateStringTag("grass_color_modifier", grass_colour_modifier)})
	TC.EndTag()
	return TC
}

// func TestJoinGame() {
// 	// Create NBT Encoders
// 	JGDimensionCodec := nbt.CreateNBTEncoder()
// 	JGDimension := nbt.CreateNBTEncoder()
// 	// Create Dimension type
// 	JGDimensionCodec.AddCompoundTag("minecraft:dimension_type")                      // Create compound
// 	JGDimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:dimension_type")) // Create string tag
// 	// Create list of values for overworld
// 	List := nbt.CreateListTag("value", nbt.TagCompound)
// List.AddTag(Overworld_Entry)                                // Add the entry to the list
// List.AddTag(Overworld_Caves_Entry)                          // Add the entry to the list
// List.AddTag(Nether_Entry)                                   // Add the entry to the list
// List.AddTag(End_Entry)                                      // Add the entry to the list                                                                                                         // End TC
// 	List.AddTag(TC)                                                                                                        // Add the compound above to the list
// 	JGDimensionCodec.AddTag(List)                                                                                          // Add the list to JGDC
// 	JGDimensionCodec.EndCompoundTag()                                                                                      // Add end tag to first compound tag in dimension_type - overworld
// 	JGDimensionCodec.AddCompoundTag("minecraft:worldgen/biome")                                                            // Create biome registry
// 	JGDimensionCodec.AddTag(nbt.CreateStringTag("type", "minecraft:worldgen/biome"))
// 	Blist := nbt.CreateListTag("value", nbt.TagCompound) // Create list
// 	TC2 := nbt.CreateCompoundTag("")
// 	TC2.AddTag(nbt.CreateStringTag("name", "minecraft:plains"))
// 	TC2.AddTag(nbt.CreateIntTag("id", 0))
// 	TC2.AddTag(CreateBiomeProperties(1.0, 1.0, 1.0, 1.0, 8364543, 8364543, 8364543, 8364543, 8364543, 8364543, "none", "plains", "", "", ""))
// 	TC2.EndTag()                      // End TC2
// 	Blist.AddTag(TC2)                 // Add compound to List
// 	JGDimensionCodec.AddTag(Blist)    // add list to JGC
// 	JGDimensionCodec.EndCompoundTag() //End biome registry
// 	JGDimensionCodec.EndCompoundTag() // End root compound
// 	TC3 := CreateDimensionTypeRegistry(1, 1, 0, 1, 1, 1, 0, 0, 1.0, 1.0, 0, 256, 64, 12000, "", "minecraft:overworld")
// 	JGDimension.AddTag(TC3)      // Add compound
// 	JGDimension.EndCompoundTag() // End root compound
// 	JGDimension.Encode()         // Encode :D
// 	JGDimensionCodec.Encode()    // Encode :D
// 	Log.Debug("CREATED JOIN GAME")
// 	Log.Info("DimensionCodec", JGDimensionCodec.GetData())
// 	// Debug - save results
// 	/*
// 		f, err := os.Create("testing.nbt")
// 		if err != nil {
// 			panic(err)
// 		}
// 		f.Write(JGDimensionCodec.GetData())
// 		f.Close()
// 		utils.PrintHexFromBytes("DimensionCodec", JGDimensionCodec.GetData())
// 		Log.Info("DimensionCodec", JGDimensionCodec.GetData())
// 	*/
// }

var TestDimension = []byte{0x0a, 0x00, 0x00, 0x08, 0x00, 0x0a, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62, 0x75, 0x72, 0x6e, 0x00, 0x1e, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3a, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62, 0x75, 0x72, 0x6e, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x08, 0x00, 0x07, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x00, 0x13, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3a, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x01, 0x00, 0x09, 0x75, 0x6c, 0x74, 0x72, 0x61, 0x77, 0x61, 0x72, 0x6d, 0x00, 0x03, 0x00, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x00, 0x00, 0x01, 0x00, 0x03, 0x00, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x07, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x01, 0x03, 0x00, 0x05, 0x6d, 0x69, 0x6e, 0x5f, 0x79, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09, 0x62, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x01, 0x06, 0x00, 0x10, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x0b, 0x70, 0x69, 0x67, 0x6c, 0x69, 0x6e, 0x5f, 0x73, 0x61, 0x66, 0x65, 0x00, 0x01, 0x00, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x73, 0x6b, 0x79, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x01, 0x01, 0x00, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x65, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x00, 0x05, 0x00, 0x0d, 0x61, 0x6d, 0x62, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x61, 0x69, 0x64, 0x73, 0x01, 0x01, 0x00, 0x14, 0x72, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x00, 0x00, 0x13, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3a, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x76, 0x5c, 0x3c, 0xed, 0x0e, 0xd8, 0x59, 0xde, 0x14, 0x0b, 0x00, 0x01, 0x00, 0x00}

func DWrite() {
	f, err := os.Create("Workpls.nbt")
	if err != nil {
		panic(err)
	}
	f.Write(TestDimension)
	f.Close()
}

func CreateBiomeRegistries() nbt.List {
	BiomeList := nbt.CreateListTagWithCapacity("value", nbt.TagCompound, 80) // Create list
	// BiomeIDMap := biome.GetBiomeIDMap()
	for _, v := range biome.GetBiomeStructs() {
		TC := nbt.CreateCompoundTag("") // Biome entry
		TC.AddTag(nbt.CreateStringTag("name", v.Name))
		TC.AddTag(nbt.CreateIntTag("id", int32(v.ID)))
		TC.AddTag(CreateBiomeProperties(1.0, 1.0, 1.0, 1.0, 65535, 65535, 65535, 65535, 65535, 65535, "none", "plains", "", "", "")) //1703705
		TC.EndTag()
		BiomeList.AddTag(TC)
	}
	// Log.Debug(BiomeList)
	return BiomeList
}

func TestBiomeList() {
	TC := nbt.CreateNBTEncoder()
	TC.AddTag(BiomeEntry)
	TC.EndCompoundTag()
	f, err := os.Create("biomelist.nbt")
	if err != nil {
		panic(err)
	}
	f.Write(TC.Encode())
}
