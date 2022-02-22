package biome

type Biome struct {
	Name string
	ID   int
	//Properties...
}

//Thanks to HoneySmoke I could retrieve the NBT sent and implement the ID's used for JoinGame worldgen/biome

// ID's of all biomes
const (
	The_Void_BiomeID = iota
	Plains_BiomeID
	Sunflower_Plains_BiomeID
	Snowy_Plains_BiomeID // NEW
	Ice_Spikes_BiomeID
	Desert_BiomeID
	Swamp_BiomeID
	Forest_BiomeID
	Flower_Forest_BiomeID
	Birch_Forest_BiomeID
	Dark_Forest_BiomeID
	Old_Growth_Birch_Forest_BiomeID // NEW
	Old_Growth_Pine_Taiga_BiomeID   // NEW
	Old_Growth_Spruce_Taiga         // NEW
	Taiga_BiomeID
	Snowy_Taiga_BiomeID
	Savanna_BiomeID
	Savanna_Plateau_BiomeID
	Windswept_Hills_BiomeID          // NEW
	Windswept_Gravelly_Hills_BiomeID // NEW
	Windswept_Forest_BiomeID         // NEW
	Windswept_Savanna_BiomeID        // NEW
	Jungle_BiomeID
	Sparse_Jungle_BiomeID // NEW
	Bamboo_Jungle_BiomeID
	Badlands_BiomeID
	Eroded_Badlands_BiomeID
	Wooded_Badlands_BiomeID // NEW
	Meadow_BiomeID          // NEW
	Grove_BiomeID           // NEW
	Snowy_Slopes_BiomeID    // NEW
	Frozen_Peaks_BiomeID    // NEW
	Jagged_Peaks_BiomeID    // NEW
	Stony_Peaks_BiomeID     // NEW
	River_BiomeID
	Frozen_River_BiomeID
	Beach_BiomeID
	Snowy_Beach_BiomeID
	Stony_Shore_BiomeID
	Warm_Ocean_BiomeID
	Lukewarm_Ocean_BiomeID
	Deep_Lukewarm_Ocean_BiomeID
	Ocean_BiomeID
	Deep_Ocean_BiomeID
	Cold_Ocean_BiomeID
	Deep_Cold_Ocean_BiomeID
	Frozen_Ocean_BiomeID
	Deep_Frozen_Ocean_BiomeID
	Mushroom_Fields_BiomeID
	Dripstone_Caves_BiomeID
	Lush_Caves_BiomeID
	Nether_Wastes_BiomeID
	Warped_Forest_BiomeID
	Crimson_Forest_BiomeID
	Soul_Sand_Valley_BiomeID
	Basalt_Deltas_BiomeID
	The_End_BiomeID
	End_Highlands_BiomeID
	End_Midlands_BiomeID
	Small_End_Islands_BiomeID
	End_Barrens_BiomeID
)

var (
	// biomeNameMap stores all object's of biomes, there name and ID and further on their peoperties
	biomeNameMap = map[string]Biome{
		"the_void":                 biomeStructs[The_Void_BiomeID],                 //0
		"plains":                   biomeStructs[Plains_BiomeID],                   //1
		"sunflower_plains":         biomeStructs[Sunflower_Plains_BiomeID],         //2
		"snowy_plains":             biomeStructs[Snowy_Plains_BiomeID],             //3
		"ice_spikes":               biomeStructs[Ice_Spikes_BiomeID],               //4
		"desert":                   biomeStructs[Desert_BiomeID],                   //5
		"swamp":                    biomeStructs[Swamp_BiomeID],                    //6
		"forest":                   biomeStructs[Forest_BiomeID],                   //7
		"flower_forest":            biomeStructs[Flower_Forest_BiomeID],            //8
		"birch_forest":             biomeStructs[Birch_Forest_BiomeID],             //9
		"dark_forest":              biomeStructs[Dark_Forest_BiomeID],              //10
		"old_growth_birch_forest":  biomeStructs[Old_Growth_Birch_Forest_BiomeID],  //11
		"old_growth_pine_taiga":    biomeStructs[Old_Growth_Pine_Taiga_BiomeID],    //12
		"old_growth_spruce":        biomeStructs[Old_Growth_Spruce_Taiga],          //13
		"taiga":                    biomeStructs[Taiga_BiomeID],                    //14
		"snowy_taiga":              biomeStructs[Snowy_Taiga_BiomeID],              //15
		"savanna":                  biomeStructs[Savanna_BiomeID],                  //16
		"savanna_plateu":           biomeStructs[Savanna_Plateau_BiomeID],          //17
		"windswept_hills":          biomeStructs[Windswept_Hills_BiomeID],          //18
		"windswept_gravelly_hills": biomeStructs[Windswept_Gravelly_Hills_BiomeID], //19
		"windswept_forest":         biomeStructs[Windswept_Forest_BiomeID],         //20
		"windswpet_savanna":        biomeStructs[Windswept_Savanna_BiomeID],        //21
		"jungle":                   biomeStructs[Jungle_BiomeID],                   //22
		"sparse_jungle":            biomeStructs[Sparse_Jungle_BiomeID],            //23
		"bamboo_jungle":            biomeStructs[Bamboo_Jungle_BiomeID],            //24
		"badlands":                 biomeStructs[Badlands_BiomeID],                 //25
		"eroded_badlands":          biomeStructs[Eroded_Badlands_BiomeID],          //26
		"wooded_badlands":          biomeStructs[Wooded_Badlands_BiomeID],          //27
		"meadow":                   biomeStructs[Meadow_BiomeID],                   //28
		"grove":                    biomeStructs[Grove_BiomeID],                    //29
		"snowy_slopes":             biomeStructs[Snowy_Slopes_BiomeID],             //30
		"frozen_peaks":             biomeStructs[Frozen_Peaks_BiomeID],             //31
		"jagged_peaks":             biomeStructs[Jagged_Peaks_BiomeID],             //32
		"stony_peaks":              biomeStructs[Stony_Peaks_BiomeID],              //33
		"river":                    biomeStructs[River_BiomeID],                    //34
		"frozen_river":             biomeStructs[Frozen_River_BiomeID],             //35
		"beach":                    biomeStructs[Beach_BiomeID],                    //36
		"snowy_beach":              biomeStructs[Snowy_Beach_BiomeID],              //37
		"stony_shore":              biomeStructs[Stony_Shore_BiomeID],              //38
		"warm_ocean":               biomeStructs[Warm_Ocean_BiomeID],               //39
		"lukewarm_ocean":           biomeStructs[Lukewarm_Ocean_BiomeID],           //40
		"deep_lukewarm_ocean":      biomeStructs[Deep_Lukewarm_Ocean_BiomeID],      //41
		"ocean":                    biomeStructs[Ocean_BiomeID],                    //42
		"deep_ocean":               biomeStructs[Deep_Ocean_BiomeID],               //43
		"cold_ocean":               biomeStructs[Cold_Ocean_BiomeID],               //44
		"deep_cold_ocean":          biomeStructs[Deep_Cold_Ocean_BiomeID],          //45
		"frozen_ocean":             biomeStructs[Frozen_Ocean_BiomeID],             //46
		"deep_frozen_ocean":        biomeStructs[Deep_Frozen_Ocean_BiomeID],        //47
		"mushroom_fields":          biomeStructs[Mushroom_Fields_BiomeID],          //48
		"dripstone_caves":          biomeStructs[Dripstone_Caves_BiomeID],          //49
		"lush_caves":               biomeStructs[Lush_Caves_BiomeID],               //50
		"nether_wastes":            biomeStructs[Nether_Wastes_BiomeID],            //51
		"warped_forest":            biomeStructs[Warped_Forest_BiomeID],            //52
		"crimson_forest":           biomeStructs[Crimson_Forest_BiomeID],           //53
		"soul_sand_valley":         biomeStructs[Soul_Sand_Valley_BiomeID],         //54
		"basalt_deltas":            biomeStructs[Basalt_Deltas_BiomeID],            //55
		"the_end":                  biomeStructs[The_End_BiomeID],                  //56
		"end_highlands":            biomeStructs[End_Highlands_BiomeID],            //57
		"end_midlands":             biomeStructs[End_Midlands_BiomeID],             //58
		"small_end_islands":        biomeStructs[Small_End_Islands_BiomeID],        //59
		"end_barrens":              biomeStructs[End_Barrens_BiomeID],              //60
	}
	// biomeIDMap stores all objects of biomes via their ID defined in the constants above
	biomeIDMap = map[int]Biome{
		The_Void_BiomeID:                 biomeStructs[The_Void_BiomeID],                 //0
		Plains_BiomeID:                   biomeStructs[Plains_BiomeID],                   //1
		Sunflower_Plains_BiomeID:         biomeStructs[Sunflower_Plains_BiomeID],         //2
		Snowy_Plains_BiomeID:             biomeStructs[Snowy_Plains_BiomeID],             //3
		Ice_Spikes_BiomeID:               biomeStructs[Ice_Spikes_BiomeID],               //4
		Desert_BiomeID:                   biomeStructs[Desert_BiomeID],                   //5
		Swamp_BiomeID:                    biomeStructs[Swamp_BiomeID],                    //6
		Forest_BiomeID:                   biomeStructs[Forest_BiomeID],                   //7
		Flower_Forest_BiomeID:            biomeStructs[Flower_Forest_BiomeID],            //8
		Birch_Forest_BiomeID:             biomeStructs[Birch_Forest_BiomeID],             //9
		Dark_Forest_BiomeID:              biomeStructs[Dark_Forest_BiomeID],              //10
		Old_Growth_Birch_Forest_BiomeID:  biomeStructs[Old_Growth_Birch_Forest_BiomeID],  //11
		Old_Growth_Pine_Taiga_BiomeID:    biomeStructs[Old_Growth_Pine_Taiga_BiomeID],    //12
		Old_Growth_Spruce_Taiga:          biomeStructs[Old_Growth_Spruce_Taiga],          //13
		Taiga_BiomeID:                    biomeStructs[Taiga_BiomeID],                    //14
		Snowy_Taiga_BiomeID:              biomeStructs[Snowy_Taiga_BiomeID],              //15
		Savanna_BiomeID:                  biomeStructs[Savanna_BiomeID],                  //16
		Savanna_Plateau_BiomeID:          biomeStructs[Savanna_Plateau_BiomeID],          //17
		Windswept_Hills_BiomeID:          biomeStructs[Windswept_Hills_BiomeID],          //18
		Windswept_Gravelly_Hills_BiomeID: biomeStructs[Windswept_Gravelly_Hills_BiomeID], //19
		Windswept_Forest_BiomeID:         biomeStructs[Windswept_Forest_BiomeID],         //20
		Windswept_Savanna_BiomeID:        biomeStructs[Windswept_Savanna_BiomeID],        //21
		Jungle_BiomeID:                   biomeStructs[Jungle_BiomeID],                   //22
		Sparse_Jungle_BiomeID:            biomeStructs[Sparse_Jungle_BiomeID],            //23
		Bamboo_Jungle_BiomeID:            biomeStructs[Bamboo_Jungle_BiomeID],            //24
		Badlands_BiomeID:                 biomeStructs[Badlands_BiomeID],                 //25
		Eroded_Badlands_BiomeID:          biomeStructs[Eroded_Badlands_BiomeID],          //26
		Wooded_Badlands_BiomeID:          biomeStructs[Wooded_Badlands_BiomeID],          //27
		Meadow_BiomeID:                   biomeStructs[Meadow_BiomeID],                   //28
		Grove_BiomeID:                    biomeStructs[Grove_BiomeID],                    //29
		Snowy_Slopes_BiomeID:             biomeStructs[Snowy_Slopes_BiomeID],             //30
		Frozen_Peaks_BiomeID:             biomeStructs[Frozen_Peaks_BiomeID],             //31
		Jagged_Peaks_BiomeID:             biomeStructs[Jagged_Peaks_BiomeID],             //32
		Stony_Peaks_BiomeID:              biomeStructs[Stony_Peaks_BiomeID],              //33
		River_BiomeID:                    biomeStructs[River_BiomeID],                    //34
		Frozen_River_BiomeID:             biomeStructs[Frozen_River_BiomeID],             //35
		Beach_BiomeID:                    biomeStructs[Beach_BiomeID],                    //36
		Snowy_Beach_BiomeID:              biomeStructs[Snowy_Beach_BiomeID],              //37
		Stony_Shore_BiomeID:              biomeStructs[Stony_Shore_BiomeID],              //38
		Warm_Ocean_BiomeID:               biomeStructs[Warm_Ocean_BiomeID],               //39
		Lukewarm_Ocean_BiomeID:           biomeStructs[Lukewarm_Ocean_BiomeID],           //40
		Deep_Lukewarm_Ocean_BiomeID:      biomeStructs[Deep_Lukewarm_Ocean_BiomeID],      //41
		Ocean_BiomeID:                    biomeStructs[Ocean_BiomeID],                    //42
		Deep_Ocean_BiomeID:               biomeStructs[Deep_Ocean_BiomeID],               //43
		Cold_Ocean_BiomeID:               biomeStructs[Cold_Ocean_BiomeID],               //44
		Deep_Cold_Ocean_BiomeID:          biomeStructs[Deep_Cold_Ocean_BiomeID],          //45
		Frozen_Ocean_BiomeID:             biomeStructs[Frozen_Ocean_BiomeID],             //46
		Deep_Frozen_Ocean_BiomeID:        biomeStructs[Deep_Frozen_Ocean_BiomeID],        //47
		Mushroom_Fields_BiomeID:          biomeStructs[Mushroom_Fields_BiomeID],          //48
		Dripstone_Caves_BiomeID:          biomeStructs[Dripstone_Caves_BiomeID],          //49
		Lush_Caves_BiomeID:               biomeStructs[Lush_Caves_BiomeID],               //50
		Nether_Wastes_BiomeID:            biomeStructs[Nether_Wastes_BiomeID],            //51
		Warped_Forest_BiomeID:            biomeStructs[Warped_Forest_BiomeID],            //52
		Crimson_Forest_BiomeID:           biomeStructs[Crimson_Forest_BiomeID],           //53
		Soul_Sand_Valley_BiomeID:         biomeStructs[Soul_Sand_Valley_BiomeID],         //54
		Basalt_Deltas_BiomeID:            biomeStructs[Basalt_Deltas_BiomeID],            //55
		The_End_BiomeID:                  biomeStructs[The_End_BiomeID],                  //56
		End_Highlands_BiomeID:            biomeStructs[End_Highlands_BiomeID],            //57
		End_Midlands_BiomeID:             biomeStructs[End_Midlands_BiomeID],             //58
		Small_End_Islands_BiomeID:        biomeStructs[Small_End_Islands_BiomeID],        //59
		End_Barrens_BiomeID:              biomeStructs[End_Barrens_BiomeID],              //60
	}
)

var biomeStructs = []Biome{
	The_Void_BiomeID:                 {"minecraft:the_void", The_Void_BiomeID}, //0
	Plains_BiomeID:                   {"minecraft:plains", Plains_BiomeID},
	Sunflower_Plains_BiomeID:         {"minecraft:sunflower_plains", Sunflower_Plains_BiomeID},
	Snowy_Plains_BiomeID:             {"minecraft:snowt_plains", Snowy_Plains_BiomeID},
	Ice_Spikes_BiomeID:               {"minecraft:ice_spikes", Ice_Spikes_BiomeID},
	Desert_BiomeID:                   {"minecraft:desert", Desert_BiomeID},
	Swamp_BiomeID:                    {"minecraft:swamp", Swamp_BiomeID},
	Forest_BiomeID:                   {"minecraft:forest", Forest_BiomeID},
	Flower_Forest_BiomeID:            {"minecraft:flower_forest", Flower_Forest_BiomeID},
	Birch_Forest_BiomeID:             {"minecraft:birch_forest", Birch_Forest_BiomeID},
	Dark_Forest_BiomeID:              {"minecraft:dark_forest", Dark_Forest_BiomeID},                           //10
	Old_Growth_Birch_Forest_BiomeID:  {"minecraft:old_growth_birch_forest", Old_Growth_Birch_Forest_BiomeID},   //11
	Old_Growth_Pine_Taiga_BiomeID:    {"minecraft:old_growth_pine_taiga", Old_Growth_Pine_Taiga_BiomeID},       //12
	Old_Growth_Spruce_Taiga:          {"minecraft:old_growth_spruce_taiga", Old_Growth_Spruce_Taiga},           //13
	Taiga_BiomeID:                    {"minecraft:taiga", Taiga_BiomeID},                                       //14
	Snowy_Taiga_BiomeID:              {"minecraft:snowy_taiga", Snowy_Taiga_BiomeID},                           //15
	Savanna_BiomeID:                  {"minecraft:savanna", Savanna_BiomeID},                                   //16
	Savanna_Plateau_BiomeID:          {"minecraft:savanna_plateu", Savanna_Plateau_BiomeID},                    //17
	Windswept_Hills_BiomeID:          {"minecraft:windswept_hill", Windswept_Hills_BiomeID},                    //18
	Windswept_Gravelly_Hills_BiomeID: {"minecraft:windswept_gravelly_hills", Windswept_Gravelly_Hills_BiomeID}, //19
	Windswept_Forest_BiomeID:         {"minecraft:windswept_forest", Windswept_Forest_BiomeID},                 //20
	Windswept_Savanna_BiomeID:        {"minecraft:windswept_savanna", Windswept_Savanna_BiomeID},               //21
	Jungle_BiomeID:                   {"minecraft:jungle", Jungle_BiomeID},                                     //22
	Sparse_Jungle_BiomeID:            {"minecraft:sparse_jungle", Sparse_Jungle_BiomeID},                       //23
	Bamboo_Jungle_BiomeID:            {"minecraft:bamboo_jungle", Bamboo_Jungle_BiomeID},                       //24
	Badlands_BiomeID:                 {"minecraft:badlands", Badlands_BiomeID},                                 //25
	Eroded_Badlands_BiomeID:          {"minecraft:eroded_badlands", Eroded_Badlands_BiomeID},                   //26
	Wooded_Badlands_BiomeID:          {"minecraft:wooded_badlands", Wooded_Badlands_BiomeID},                   //28
	Meadow_BiomeID:                   {"minecraft:meadow", Meadow_BiomeID},                                     //28
	Grove_BiomeID:                    {"minecraft:grove", Grove_BiomeID},                                       //29
	Snowy_Slopes_BiomeID:             {"minecraft:snowy_slopes", Snowy_Slopes_BiomeID},                         //30
	Frozen_Peaks_BiomeID:             {"minecraft:frozen_peaks", Frozen_Peaks_BiomeID},                         //32
	Jagged_Peaks_BiomeID:             {"minecraft:jagged_peaks", Jagged_Peaks_BiomeID},                         //32
	Stony_Peaks_BiomeID:              {"minecraft:stony_peaks", Stony_Peaks_BiomeID},                           //33
	River_BiomeID:                    {"minecraft:river", River_BiomeID},                                       //34
	Frozen_River_BiomeID:             {"minecraft:frozen_river", Frozen_River_BiomeID},                         //35
	Beach_BiomeID:                    {"minecraft:beach", Beach_BiomeID},                                       //36
	Snowy_Beach_BiomeID:              {"minecraft:snowy_beach", Snowy_Beach_BiomeID},                           //37
	Stony_Shore_BiomeID:              {"minecraft:stony_shore", Stony_Shore_BiomeID},                           //38
	Warm_Ocean_BiomeID:               {"minecraft:warm_ocean", Warm_Ocean_BiomeID},                             //39
	Lukewarm_Ocean_BiomeID:           {"minecraft:lukewarm_ocean", Lukewarm_Ocean_BiomeID},                     //40
	Deep_Lukewarm_Ocean_BiomeID:      {"minecraft:deep_lukewarm_ocean", Deep_Lukewarm_Ocean_BiomeID},           //41
	Ocean_BiomeID:                    {"minecraft:ocean", Ocean_BiomeID},                                       //42
	Deep_Ocean_BiomeID:               {"minecraft:deep_ocean", Deep_Ocean_BiomeID},                             //43
	Cold_Ocean_BiomeID:               {"minecraft:cold_ocean", Cold_Ocean_BiomeID},                             //44
	Deep_Cold_Ocean_BiomeID:          {"minecraft:deep_cold_ocean", Deep_Cold_Ocean_BiomeID},                   //45
	Frozen_Ocean_BiomeID:             {"minecraft:frozen_ocean", Frozen_Ocean_BiomeID},                         //46
	Deep_Frozen_Ocean_BiomeID:        {"minecraft:deep_frozen_ocean", Deep_Frozen_Ocean_BiomeID},               //47
	Mushroom_Fields_BiomeID:          {"minecraft:mushroom_fields", Mushroom_Fields_BiomeID},                   //48
	Dripstone_Caves_BiomeID:          {"minecraft:dripstone_caves", Dripstone_Caves_BiomeID},                   //49
	Lush_Caves_BiomeID:               {"minecraft:lush_caves", Lush_Caves_BiomeID},                             //50
	Nether_Wastes_BiomeID:            {"minecraft:nether_wastes", Nether_Wastes_BiomeID},                       //51
	Warped_Forest_BiomeID:            {"minecraft:warped_forest", Warped_Forest_BiomeID},                       //52
	Crimson_Forest_BiomeID:           {"minecraft:crimson_forest", Crimson_Forest_BiomeID},                     //53
	Soul_Sand_Valley_BiomeID:         {"minecraft:soul_sand_valley", Soul_Sand_Valley_BiomeID},                 //54
	Basalt_Deltas_BiomeID:            {"minecraft:basalt_deltas", Basalt_Deltas_BiomeID},                       //55
	The_End_BiomeID:                  {"minecraft:the_end", The_End_BiomeID},                                   //56
	End_Highlands_BiomeID:            {"minecraft:end_highlands", End_Highlands_BiomeID},                       //57
	End_Midlands_BiomeID:             {"minecraft:end_midlands", End_Midlands_BiomeID},                         //58
	Small_End_Islands_BiomeID:        {"minecraft:small_end_islands", Small_End_Islands_BiomeID},               //59
	End_Barrens_BiomeID:              {"minecraft:end_barrens", End_Barrens_BiomeID}}                           //60

func GetBiomeByID(ID int) Biome {
	return biomeIDMap[ID]
}

func GetBiomeByName(Name string) Biome {
	return biomeNameMap[Name]
}

func GetBiomeIDMap() map[int]Biome {
	return biomeIDMap
}

func GetBiomeNameMap() map[string]Biome {
	return biomeNameMap
}

func GetBiomeStructs() []Biome {
	return biomeStructs
}
