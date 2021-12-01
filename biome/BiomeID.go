package biome

type Biome struct {
	Name string
	ID   int
	//Properties...
}

//Thanks to HoneySmoke I could retrieve the NBT sent and implement the ID's used for JoinGame worldgen/biome

// ID's of all biomes
const (
	Ocean_BiomeID = iota
	Plains_BiomeID
	Desert_BiomeID
	Mountains_BiomeID
	Forest_BiomeID
	Taiga_BiomeID
	Swamp_BiomeID
	River_BiomeID
	Nether_Wastes_BiomeID
	The_End_BiomeID
	Frozen_Ocean_BiomeID
	Frozen_River_BiomeID
	Snowy_Tundra_BiomeID
	Snowy_Mountains_BiomeID
	Mushroom_Fields_BiomeID
	Mushroom_Fields_Shore_BiomeID
	Beach_BiomeID
	Desert_Hills_BiomeID
	Wooded_Hills_BiomeID
	Taiga_Hills_BiomeID
	Mountain_Edge_BiomeID
	Jungle_BiomeID
	Jungle_Hills_BiomeID
	Jungle_Edge_BiomeID
	Deep_Ocean_BiomeID
	Stone_Shore_BiomeID
	Snowy_Beach_BiomeID
	Birch_Forest_BiomeID
	Birch_Forest_Hills_BiomeID
	Dark_Forest_BiomeID
	Snowy_Taiga_BiomeID
	Snowy_Taiga_Hills_BiomeID
	Giant_Tree_Taiga_BiomeID
	Giant_Tree_Taiga_Hills_BiomeID
	Wooded_Mountains_BiomeID
	Savanna_BiomeID
	Savanna_Plateau_BiomeID
	Badlands_BiomeID
	Wooded_Badlands_Plateau_BiomeID
	Badlands_Plateau_BiomeID
	Small_End_Islands_BiomeID
	End_Midlands_BiomeID
	End_Highlands_BiomeID
	End_Barrens_BiomeID
	Warm_Ocean_BiomeID
	Lukewarm_Ocean_BiomeID
	Cold_Ocean_BiomeID
	Deep_Warm_Ocean_BiomeID
	Deep_Lukewarm_Ocean_BiomeID
	Deep_Cold_Ocean_BiomeID
	Deep_Frozen_Ocean_BiomeID
	The_Void_BiomeID         = 127
	Sunflower_Plains_BiomeID = iota + 77
	Desert_Lakes_BiomeID
	Gravelly_Mountains_BiomeID
	Flower_Forest_BiomeID
	Taiga_Mountains_BiomeID
	Swamp_Hills_BiomeID
	Ice_Spikes_BiomeID                       = 140
	Modified_Jungle_BiomeID                  = 149
	Modified_Jungle_Edge_BiomeID             = 151
	Tall_Birch_Forest_BiomeID                = 155
	Tall_Birch_Hills_BiomeID                 = 156
	Dark_Forest_Hills_BiomeID                = 157
	Snowy_Taiga_Mountains_BiomeID            = 158
	Giant_Spruce_Taiga_BiomeID               = 160
	Modified_Gravelly_Mountains_BiomeID      = 162
	Shattered_Savanna_BiomeID                = 163
	Shattered_Savanna_Plateau_BiomeID        = 164
	Eroded_Badlands_BiomeID                  = 165
	Modified_Wooded_Badlands_Plateau_BiomeID = 166
	Modified_Badlands_Plateau_BiomeID        = 167
	Bamboo_Jungle_BiomeID                    = 168
	Bamboo_Jungle_Hills_BiomeID              = 169
	Soul_Sand_Valley_BiomeID                 = 170
	Crimson_Forest_BiomeID                   = 171
	Warped_Forest_BiomeID                    = 172
	Basalt_Deltas_BiomeID                    = 173
	Dripstone_Caves_BiomeID                  = 174
	Lush_Caves_BiomeID                       = 175
)

var (
	// biomeNameMap stores all object's of biomes, there name and ID and further on their peoperties
	biomeNameMap = map[string]Biome{
		"ocean":                           biomeStructs[0],  //{"minecraft:ocean", Ocean_BiomeID},
		"plains":                          biomeStructs[1],  //{"minecraft:plains", Plains_BiomeID},
		"desert":                          biomeStructs[2],  //{"minecraft:desert", Desert_BiomeID},
		"mountains":                       biomeStructs[3],  //{"minecraft:mountains", Mountains_BiomeID},
		"forest":                          biomeStructs[4],  //{"minecraft:forest", Forest_BiomeID},
		"taiga":                           biomeStructs[5],  //{"minecraft:taiga", Taiga_BiomeID},
		"swamp":                           biomeStructs[6],  //{"minecraft:swamp", Swamp_BiomeID},
		"river":                           biomeStructs[7],  //{"minecraft:river", River_BiomeID},
		"nether_wastes":                   biomeStructs[8],  //{"minecraft:nether_wastes", Nether_Wastes_BiomeID},
		"the_end":                         biomeStructs[9],  //{"minecraft:the_end", The_End_BiomeID},
		"frozen_ocean":                    biomeStructs[10], //{"minecraft:frozen_ocean", Frozen_Ocean_BiomeID},
		"frozen_river":                    biomeStructs[11], //{"minecraft:frozen_river", Frozen_River_BiomeID},
		"snowy_tundra":                    biomeStructs[12], //{"minecraft:snowy_tundra", Snowy_Tundra_BiomeID},
		"snowy_mountains":                 biomeStructs[13], //{"minecraft:snowy_mountains", Snowy_Mountains_BiomeID},
		"mushroom_fields":                 biomeStructs[14], //{"minecraft:mushroom_fields", Mushroom_Fields_BiomeID},
		"mushroom_fields_shore":           biomeStructs[15], //{"minecraft:mushroom_fields_shore", Mushroom_Fields_Shore_BiomeID},
		"beach":                           biomeStructs[16], //{"minecraft:beach", Beach_BiomeID},
		"desert_hills":                    biomeStructs[17], //{"minecraft:desert_hills", Desert_Hills_BiomeID},
		"wooded_hills":                    biomeStructs[18], //{"minecraft:wooded_hills", Wooded_Hills_BiomeID},
		"taiga_hills":                     biomeStructs[19], //{"minecraft:taiga_hills", Taiga_Hills_BiomeID},
		"mountain_edge":                   biomeStructs[20], //{"minecraft:mountain_edge", Mountain_Edge_BiomeID},
		"jungle":                          biomeStructs[21], //{"minecraft:jungle", Jungle_BiomeID},
		"jungle_hills":                    biomeStructs[22], //{"minecraft:jungle_hills", Jungle_Hills_BiomeID},
		"jungle_edge":                     biomeStructs[23], //{"minecraft:jungle_edge", Jungle_Edge_BiomeID},
		"deep_ocean":                      biomeStructs[24], //{"minecraft:deep_ocean", Deep_Ocean_BiomeID},
		"stone_shore":                     biomeStructs[25], //{"minecraft:stone_shore", Stone_Shore_BiomeID},
		"snowy_beach":                     biomeStructs[26], //{"minecraft:snowy_beach", Snowy_Beach_BiomeID},
		"birch_forest":                    biomeStructs[27], //{"minecraft:birch_forest", Birch_Forest_BiomeID},
		"birch_forest_hills":              biomeStructs[28], //{"minecraft:birch_forest_hills", Birch_Forest_Hills_BiomeID},
		"dark_forest":                     biomeStructs[29], //{"minecraft:dark_forest", Dark_Forest_BiomeID},
		"snowy_taiga":                     biomeStructs[30], //{"minecraft:snowy_taiga", Snowy_Taiga_BiomeID},
		"snowy_taiga_hills":               biomeStructs[31], //{"minecraft:snowy_taiga_hills", Snowy_Taiga_Hills_BiomeID},
		"giant_tree_taiga":                biomeStructs[32], //{"minecraft:giant_tree_taiga", Giant_Tree_Taiga_BiomeID},
		"giant_tree_taiga_hills":          biomeStructs[33], //{"minecraft:giant_tree_taiga_hills", Giant_Tree_Taiga_Hills_BiomeID},
		"wooded_mountains":                biomeStructs[34], //{"minecraft:wooded_mountains", Wooded_Mountains_BiomeID},
		"savanna":                         biomeStructs[35], //{"minecraft:savanna", Savanna_BiomeID},
		"savanna_plateu":                  biomeStructs[36], //{"minecraft:savanna_plateu", Savanna_Plateau_BiomeID},
		"badlands":                        biomeStructs[37], //{"minecraft:badlands", Badlands_BiomeID},
		"wooded_badlands_plateu":          biomeStructs[38], //{"minecraft:wooded_badlands_plateu", Wooded_Badlands_Plateau_BiomeID},
		"badlands_plateu":                 biomeStructs[39], //{"minecraft:badlands_plateu", Badlands_Plateau_BiomeID},
		"small_end_islands":               biomeStructs[40], //{"minecraft:small_end_islands", Small_End_Islands_BiomeID},
		"end_midlands":                    biomeStructs[41], //{"minecraft:end_midlands", End_Midlands_BiomeID},
		"end_highlands":                   biomeStructs[42], //{"minecraft:end_highlands", End_Highlands_BiomeID},
		"end_barrens":                     biomeStructs[43], //{"minecraft:end_barrens", End_Barrens_BiomeID},
		"warm_ocean":                      biomeStructs[44], //{"minecraft:warm_ocean", Warm_Ocean_BiomeID},
		"lukewarm_ocean":                  biomeStructs[45], //{"minecraft:lukewarm_ocean", Lukewarm_Ocean_BiomeID},
		"cold_ocean":                      biomeStructs[46], //{"minecraft:cold_ocean", Cold_Ocean_BiomeID},
		"deep_warm_ocean":                 biomeStructs[47], //{"minecraft:deep_warm_ocean", Deep_Warm_Ocean_BiomeID},
		"deep_lukewarm_ocean":             biomeStructs[48], //{"minecraft:deep_lukewarm_ocean", Deep_Lukewarm_Ocean_BiomeID},
		"deep_cold_ocean":                 biomeStructs[49], //{"minecraft:deep_cold_:ocean", Deep_Cold_Ocean_BiomeID},
		"deep_frozen_ocean":               biomeStructs[50], //{"minecraft:deep_frozen_ocean", Deep_Frozen_Ocean_BiomeID},
		"the_void":                        biomeStructs[51], //{"minecraft:the_void", The_Void_BiomeID},
		"sunflower_plains":                biomeStructs[52], //{"minecraft:sunflower_plains", Sunflower_Plains_BiomeID},
		"desert_lakes":                    biomeStructs[53], //{"minecraft:desert_lakes", Desert_BiomeID},
		"gravelly_mountains":              biomeStructs[54], //{"minecraft:gravelly_mountains", Gravelly_Mountains_BiomeID},
		"flower_forest":                   biomeStructs[55], //{"minecraft:flower_forest", Flower_Forest_BiomeID},
		"taiga_mountains":                 biomeStructs[56], //{"minecraft:taiga_mountains", Taiga_Mountains_BiomeID},
		"swamp_hills":                     biomeStructs[57], //{"minecraft:swamp_hills", Swamp_Hills_BiomeID},
		"ice_spikes":                      biomeStructs[58], //{"minecraft:ice_spikes", Ice_Spikes_BiomeID},
		"modified_jungle":                 biomeStructs[59], //{"minecraft:modified_junlge", Modified_Jungle_BiomeID},
		"modified_jungle_edge":            biomeStructs[60], //{"minecraft:modified_jungle_edge", Modified_Jungle_Edge_BiomeID},
		"tall_birch_forest":               biomeStructs[61], //{"minecraft:tall_birch_forest", Tall_Birch_Forest_BiomeID},
		"tall_birch_hills":                biomeStructs[62], //{"minecraft:tall_birch_hills", Tall_Birch_Hills_BiomeID},
		"dark_forest_hills":               biomeStructs[63], //{"minecraft:dark_forest_hills", Dark_Forest_Hills_BiomeID},
		"snowy_taiga_mountains":           biomeStructs[64], //{"minecraft:snowy_taiga_mountains", Snowy_Taiga_Mountains_BiomeID},
		"giant_spruce_taiga":              biomeStructs[65], //{"minecraft:giant_spruce_taiga", Giant_Spruce_Taiga_BiomeID},
		"modified_gravelly_mountains":     biomeStructs[66], //{"minecraft:modified_gavelly_mountains", Modified_Gravelly_Mountains_BiomeID},
		"shattered_savanna":               biomeStructs[67], //{"minecraft:shattered_savanna", Shattered_Savanna_BiomeID},
		"shattered_savanna_plateu":        biomeStructs[68], //{"minecraft:shattered_savanna_plateu", Shattered_Savanna_Plateau_BiomeID},
		"eroded_badlands":                 biomeStructs[69], //{"minecraft:eroded_badlands", Eroded_Badlands_BiomeID},
		"modified_wooded_badlands_plateu": biomeStructs[70], //{"minecraft:modified_wooded_badlands_plateu", Modified_Wooded_Badlands_Plateau_BiomeID},
		"modified_badlands_plateu":        biomeStructs[71], //{"minecraft:modified_badlands_plateu", Modified_Badlands_Plateau_BiomeID},
		"bamboo_jungle":                   biomeStructs[72], //{"minecraft:bamboo_jungle", Bamboo_Jungle_BiomeID},
		"bamboo_jungle_hills":             biomeStructs[73], //{"minecraft:bamboo_jungle_hills", Bamboo_Jungle_Hills_BiomeID},
		"soul_sand_valley":                biomeStructs[74], //{"minecraft:soul_sand_valley", Soul_Sand_Valley_BiomeID},
		"crimson_forest":                  biomeStructs[75], //{"minecraft:crimson_forest", Crimson_Forest_BiomeID},
		"warped_forest":                   biomeStructs[76], //{"minecraft:warped_forest", Warped_Forest_BiomeID},
		"basalt_deltas":                   biomeStructs[77], //{"minecraft:basalt_deltas", Basalt_Deltas_BiomeID},
		"dripstone_caves":                 biomeStructs[78], //{"minecraft:dripstone_caves", Dripstone_Caves_BiomeID},
		"lush_caves":                      biomeStructs[79], //{"minecraft:lush_caves", Lush_Caves_BiomeID},
	}
	// biomeIDMap stores all objects of biomes via their ID defined in the constants above
	biomeIDMap = map[int]Biome{
		Ocean_BiomeID:                            biomeStructs[0],  //{"minecraft:ocean", Ocean_BiomeID},
		Plains_BiomeID:                           biomeStructs[1],  //{"minecraft:plains", Plains_BiomeID},
		Desert_BiomeID:                           biomeStructs[2],  //{"minecraft:desert", Desert_BiomeID},
		Mountains_BiomeID:                        biomeStructs[3],  //{"minecraft:mountains", Mountains_BiomeID},
		Forest_BiomeID:                           biomeStructs[4],  //{"minecraft:forest", Forest_BiomeID},
		Taiga_BiomeID:                            biomeStructs[5],  //{"minecraft:taiga", Taiga_BiomeID},
		Swamp_BiomeID:                            biomeStructs[6],  //{"minecraft:swamp", Swamp_BiomeID},
		River_BiomeID:                            biomeStructs[7],  //{"minecraft:river", River_BiomeID},
		Nether_Wastes_BiomeID:                    biomeStructs[8],  //{"minecraft:nether_wastes", Nether_Wastes_BiomeID},
		The_End_BiomeID:                          biomeStructs[9],  //{"minecraft:the_end", The_End_BiomeID},
		Frozen_Ocean_BiomeID:                     biomeStructs[10], //{"minecraft:frozen_ocean", Frozen_Ocean_BiomeID},
		Frozen_River_BiomeID:                     biomeStructs[11], //{"minecraft:frozen_river", Frozen_River_BiomeID},
		Snowy_Tundra_BiomeID:                     biomeStructs[12], //{"minecraft:snowy_tundra", Snowy_Tundra_BiomeID},
		Snowy_Mountains_BiomeID:                  biomeStructs[13], //{"minecraft:snowy_mountains", Snowy_Mountains_BiomeID},
		Mushroom_Fields_BiomeID:                  biomeStructs[14], //{"minecraft:mushroom_fields", Mushroom_Fields_BiomeID},
		Mushroom_Fields_Shore_BiomeID:            biomeStructs[15], //{"minecraft:mushroom_fields_shore", Mushroom_Fields_Shore_BiomeID},
		Beach_BiomeID:                            biomeStructs[16], //{"minecraft:beach", Beach_BiomeID},
		Desert_Hills_BiomeID:                     biomeStructs[17], //{"minecraft:desert_hills", Desert_Hills_BiomeID},
		Wooded_Hills_BiomeID:                     biomeStructs[18], //{"minecraft:wooded_hills", Wooded_Hills_BiomeID},
		Taiga_Hills_BiomeID:                      biomeStructs[19], //{"minecraft:taiga_hills", Taiga_Hills_BiomeID},
		Mountain_Edge_BiomeID:                    biomeStructs[20], //{"minecraft:mountain_edge", Mountain_Edge_BiomeID},
		Jungle_BiomeID:                           biomeStructs[21], //{"minecraft:jungle", Jungle_BiomeID},
		Jungle_Hills_BiomeID:                     biomeStructs[22], //{"minecraft:jungle_hills", Jungle_Hills_BiomeID},
		Jungle_Edge_BiomeID:                      biomeStructs[23], //{"minecraft:jungle_edge", Jungle_Edge_BiomeID},
		Deep_Ocean_BiomeID:                       biomeStructs[24], //{"minecraft:deep_ocean", Deep_Ocean_BiomeID},
		Stone_Shore_BiomeID:                      biomeStructs[25], //{"minecraft:stone_shore", Stone_Shore_BiomeID},
		Snowy_Beach_BiomeID:                      biomeStructs[26], //{"minecraft:snowy_beach", Snowy_Beach_BiomeID},
		Birch_Forest_BiomeID:                     biomeStructs[27], //{"minecraft:birch_forest", Birch_Forest_BiomeID},
		Birch_Forest_Hills_BiomeID:               biomeStructs[28], //{"minecraft:birch_forest_hills", Birch_Forest_Hills_BiomeID},
		Dark_Forest_BiomeID:                      biomeStructs[29], //{"minecraft:dark_forest", Dark_Forest_BiomeID},
		Snowy_Taiga_BiomeID:                      biomeStructs[30], //{"minecraft:snowy_taiga", Snowy_Taiga_BiomeID},
		Snowy_Taiga_Hills_BiomeID:                biomeStructs[31], //{"minecraft:snowy_taiga_hills", Snowy_Taiga_Hills_BiomeID},
		Giant_Tree_Taiga_BiomeID:                 biomeStructs[32], //{"minecraft:giant_tree_taiga", Giant_Tree_Taiga_BiomeID},
		Giant_Tree_Taiga_Hills_BiomeID:           biomeStructs[33], //{"minecraft:giant_tree_taiga_hills", Giant_Tree_Taiga_Hills_BiomeID},
		Wooded_Mountains_BiomeID:                 biomeStructs[34], //{"minecraft:wooded_mountains", Wooded_Mountains_BiomeID},
		Savanna_BiomeID:                          biomeStructs[35], //{"minecraft:savanna", Savanna_BiomeID},
		Savanna_Plateau_BiomeID:                  biomeStructs[36], //{"minecraft:savanna_plateu", Savanna_Plateau_BiomeID},
		Badlands_BiomeID:                         biomeStructs[37], //{"minecraft:badlands", Badlands_BiomeID},
		Wooded_Badlands_Plateau_BiomeID:          biomeStructs[38], //{"minecraft:wooded_badlands_plateu", Wooded_Badlands_Plateau_BiomeID},
		Badlands_Plateau_BiomeID:                 biomeStructs[39], //{"minecraft:badlands_plateu", Badlands_Plateau_BiomeID},
		Small_End_Islands_BiomeID:                biomeStructs[40], //{"minecraft:small_end_islands", Small_End_Islands_BiomeID},
		End_Midlands_BiomeID:                     biomeStructs[41], //{"minecraft:end_midlands", End_Midlands_BiomeID},
		End_Highlands_BiomeID:                    biomeStructs[42], //{"minecraft:end_highlands", End_Highlands_BiomeID},
		End_Barrens_BiomeID:                      biomeStructs[43], //{"minecraft:end_barrens", End_Barrens_BiomeID},
		Warm_Ocean_BiomeID:                       biomeStructs[44], //{"minecraft:warm_ocean", Warm_Ocean_BiomeID},
		Lukewarm_Ocean_BiomeID:                   biomeStructs[45], //{"minecraft:lukewarm_ocean", Lukewarm_Ocean_BiomeID},
		Cold_Ocean_BiomeID:                       biomeStructs[46], //{"minecraft:cold_ocean", Cold_Ocean_BiomeID},
		Deep_Warm_Ocean_BiomeID:                  biomeStructs[47], //{"minecraft:deep_warm_ocean", Deep_Warm_Ocean_BiomeID},
		Deep_Lukewarm_Ocean_BiomeID:              biomeStructs[48], //{"minecraft:deep_lukewarm_ocean", Deep_Lukewarm_Ocean_BiomeID},
		Deep_Cold_Ocean_BiomeID:                  biomeStructs[49], //{"minecraft:deep_cold_:ocean", Deep_Cold_Ocean_BiomeID},
		Deep_Frozen_Ocean_BiomeID:                biomeStructs[50], //{"minecraft:deep_frozen_ocean", Deep_Frozen_Ocean_BiomeID},
		The_Void_BiomeID:                         biomeStructs[51], //{"minecraft:the_void", The_Void_BiomeID},
		Sunflower_Plains_BiomeID:                 biomeStructs[52], //{"minecraft:sunflower_plains", Sunflower_Plains_BiomeID},
		Desert_Lakes_BiomeID:                     biomeStructs[53], //{"minecraft:desert_lakes", Desert_BiomeID},
		Gravelly_Mountains_BiomeID:               biomeStructs[54], //{"minecraft:gravelly_mountains", Gravelly_Mountains_BiomeID},
		Flower_Forest_BiomeID:                    biomeStructs[55], //{"minecraft:flower_forest", Flower_Forest_BiomeID},
		Taiga_Mountains_BiomeID:                  biomeStructs[56], //{"minecraft:taiga_mountains", Taiga_Mountains_BiomeID},
		Swamp_Hills_BiomeID:                      biomeStructs[57], //{"minecraft:swamp_hills", Swamp_Hills_BiomeID},
		Ice_Spikes_BiomeID:                       biomeStructs[58], //{"minecraft:ice_spikes", Ice_Spikes_BiomeID},
		Modified_Jungle_BiomeID:                  biomeStructs[59], //{"minecraft:modified_junlge", Modified_Jungle_BiomeID},
		Modified_Jungle_Edge_BiomeID:             biomeStructs[60], //{"minecraft:modified_jungle_edge", Modified_Jungle_Edge_BiomeID},
		Tall_Birch_Forest_BiomeID:                biomeStructs[61], //{"minecraft:tall_birch_forest", Tall_Birch_Forest_BiomeID},
		Tall_Birch_Hills_BiomeID:                 biomeStructs[62], //{"minecraft:tall_birch_hills", Tall_Birch_Hills_BiomeID},
		Dark_Forest_Hills_BiomeID:                biomeStructs[63], //{"minecraft:dark_forest_hills", Dark_Forest_Hills_BiomeID},
		Snowy_Taiga_Mountains_BiomeID:            biomeStructs[64], //{"minecraft:snowy_taiga_mountains", Snowy_Taiga_Mountains_BiomeID},
		Giant_Spruce_Taiga_BiomeID:               biomeStructs[65], //{"minecraft:giant_spruce_taiga", Giant_Spruce_Taiga_BiomeID},
		Modified_Gravelly_Mountains_BiomeID:      biomeStructs[66], //{"minecraft:modified_gavelly_mountains", Modified_Gravelly_Mountains_BiomeID},
		Shattered_Savanna_BiomeID:                biomeStructs[67], //{"minecraft:shattered_savanna", Shattered_Savanna_BiomeID},
		Shattered_Savanna_Plateau_BiomeID:        biomeStructs[68], //{"minecraft:shattered_savanna_plateu", Shattered_Savanna_Plateau_BiomeID},
		Eroded_Badlands_BiomeID:                  biomeStructs[69], //{"minecraft:eroded_badlands", Eroded_Badlands_BiomeID},
		Modified_Wooded_Badlands_Plateau_BiomeID: biomeStructs[70], //{"minecraft:modified_wooded_badlands_plateu", Modified_Wooded_Badlands_Plateau_BiomeID},
		Modified_Badlands_Plateau_BiomeID:        biomeStructs[71], //{"minecraft:modified_badlands_plateu", Modified_Badlands_Plateau_BiomeID},
		Bamboo_Jungle_BiomeID:                    biomeStructs[72], //{"minecraft:bamboo_jungle", Bamboo_Jungle_BiomeID},
		Bamboo_Jungle_Hills_BiomeID:              biomeStructs[73], //{"minecraft:bamboo_jungle_hills", Bamboo_Jungle_Hills_BiomeID},
		Soul_Sand_Valley_BiomeID:                 biomeStructs[74], //{"minecraft:soul_sand_valley", Soul_Sand_Valley_BiomeID},
		Crimson_Forest_BiomeID:                   biomeStructs[75], //{"minecraft:crimson_forest", Crimson_Forest_BiomeID},
		Warped_Forest_BiomeID:                    biomeStructs[76], //{"minecraft:warped_forest", Warped_Forest_BiomeID},
		Basalt_Deltas_BiomeID:                    biomeStructs[77], //{"minecraft:basalt_deltas", Basalt_Deltas_BiomeID},
		Dripstone_Caves_BiomeID:                  biomeStructs[78], //{"minecraft:dripstone_caves", Dripstone_Caves_BiomeID},
		Lush_Caves_BiomeID:                       biomeStructs[79], //{"minecraft:lush_caves", Lush_Caves_BiomeID},
	}
)

var biomeStructs = []Biome{{"minecraft:ocean", Ocean_BiomeID},
	{"minecraft:plains", Plains_BiomeID},
	{"minecraft:desert", Desert_BiomeID},
	{"minecraft:mountains", Mountains_BiomeID},
	{"minecraft:forest", Forest_BiomeID},
	{"minecraft:taiga", Taiga_BiomeID},
	{"minecraft:swamp", Swamp_BiomeID},
	{"minecraft:river", River_BiomeID},
	{"minecraft:nether_wastes", Nether_Wastes_BiomeID},
	{"minecraft:the_end", The_End_BiomeID},
	{"minecraft:frozen_ocean", Frozen_Ocean_BiomeID},
	{"minecraft:frozen_river", Frozen_River_BiomeID},
	{"minecraft:snowy_tundra", Snowy_Tundra_BiomeID},
	{"minecraft:snowy_mountains", Snowy_Mountains_BiomeID},
	{"minecraft:mushroom_fields", Mushroom_Fields_BiomeID},
	{"minecraft:mushroom_fields_shore", Mushroom_Fields_Shore_BiomeID},
	{"minecraft:beach", Beach_BiomeID},
	{"minecraft:desert_hills", Desert_Hills_BiomeID},
	{"minecraft:wooded_hills", Wooded_Hills_BiomeID},
	{"minecraft:taiga_hills", Taiga_Hills_BiomeID},
	{"minecraft:mountain_edge", Mountain_Edge_BiomeID},
	{"minecraft:jungle", Jungle_BiomeID},
	{"minecraft:jungle_hills", Jungle_Hills_BiomeID},
	{"minecraft:jungle_edge", Jungle_Edge_BiomeID},
	{"minecraft:deep_ocean", Deep_Ocean_BiomeID},
	{"minecraft:stone_shore", Stone_Shore_BiomeID},
	{"minecraft:snowy_beach", Snowy_Beach_BiomeID},
	{"minecraft:birch_forest", Birch_Forest_BiomeID},
	{"minecraft:birch_forest_hills", Birch_Forest_Hills_BiomeID},
	{"minecraft:dark_forest", Dark_Forest_BiomeID},
	{"minecraft:snowy_taiga", Snowy_Taiga_BiomeID},
	{"minecraft:snowy_taiga_hills", Snowy_Taiga_Hills_BiomeID},
	{"minecraft:giant_tree_taiga", Giant_Tree_Taiga_BiomeID},
	{"minecraft:giant_tree_taiga_hills", Giant_Tree_Taiga_Hills_BiomeID},
	{"minecraft:wooded_mountains", Wooded_Mountains_BiomeID},
	{"minecraft:savanna", Savanna_BiomeID},
	{"minecraft:savanna_plateu", Savanna_Plateau_BiomeID},
	{"minecraft:badlands", Badlands_BiomeID},
	{"minecraft:wooded_badlands_plateu", Wooded_Badlands_Plateau_BiomeID},
	{"minecraft:badlands_plateu", Badlands_Plateau_BiomeID},
	{"minecraft:small_end_islands", Small_End_Islands_BiomeID},
	{"minecraft:end_midlands", End_Midlands_BiomeID},
	{"minecraft:end_highlands", End_Highlands_BiomeID},
	{"minecraft:end_barrens", End_Barrens_BiomeID},
	{"minecraft:warm_ocean", Warm_Ocean_BiomeID},
	{"minecraft:lukewarm_ocean", Lukewarm_Ocean_BiomeID},
	{"minecraft:cold_ocean", Cold_Ocean_BiomeID},
	{"minecraft:deep_warm_ocean", Deep_Warm_Ocean_BiomeID},
	{"minecraft:deep_lukewarm_ocean", Deep_Lukewarm_Ocean_BiomeID},
	{"minecraft:deep_cold_ocean", Deep_Cold_Ocean_BiomeID},
	{"minecraft:deep_frozen_ocean", Deep_Frozen_Ocean_BiomeID},
	{"minecraft:the_void", The_Void_BiomeID},
	{"minecraft:sunflower_plains", Sunflower_Plains_BiomeID},
	{"minecraft:desert_lakes", Desert_BiomeID},
	{"minecraft:gravelly_mountains", Gravelly_Mountains_BiomeID},
	{"minecraft:flower_forest", Flower_Forest_BiomeID},
	{"minecraft:taiga_mountains", Taiga_Mountains_BiomeID},
	{"minecraft:swamp_hills", Swamp_Hills_BiomeID},
	{"minecraft:ice_spikes", Ice_Spikes_BiomeID},
	{"minecraft:modified_junlge", Modified_Jungle_BiomeID},
	{"minecraft:modified_jungle_edge", Modified_Jungle_Edge_BiomeID},
	{"minecraft:tall_birch_forest", Tall_Birch_Forest_BiomeID},
	{"minecraft:tall_birch_hills", Tall_Birch_Hills_BiomeID},
	{"minecraft:dark_forest_hills", Dark_Forest_Hills_BiomeID},
	{"minecraft:snowy_taiga_mountains", Snowy_Taiga_Mountains_BiomeID},
	{"minecraft:giant_spruce_taiga", Giant_Spruce_Taiga_BiomeID},
	{"minecraft:modified_gavelly_mountains", Modified_Gravelly_Mountains_BiomeID},
	{"minecraft:shattered_savanna", Shattered_Savanna_BiomeID},
	{"minecraft:shattered_savanna_plateu", Shattered_Savanna_Plateau_BiomeID},
	{"minecraft:eroded_badlands", Eroded_Badlands_BiomeID},
	{"minecraft:modified_wooded_badlands_plateu", Modified_Wooded_Badlands_Plateau_BiomeID},
	{"minecraft:modified_badlands_plateu", Modified_Badlands_Plateau_BiomeID},
	{"minecraft:bamboo_jungle", Bamboo_Jungle_BiomeID},
	{"minecraft:bamboo_jungle_hills", Bamboo_Jungle_Hills_BiomeID},
	{"minecraft:soul_sand_valley", Soul_Sand_Valley_BiomeID},
	{"minecraft:crimson_forest", Crimson_Forest_BiomeID},
	{"minecraft:warped_forest", Warped_Forest_BiomeID},
	{"minecraft:basalt_deltas", Basalt_Deltas_BiomeID},
	{"minecraft:dripstone_caves", Dripstone_Caves_BiomeID},
	{"minecraft:lush_caves", Lush_Caves_BiomeID}}

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
