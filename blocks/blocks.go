package blocks

import logging "github.com/op/go-logging"

var log = logging.MustGetLogger("HoneyGO")

//BlockData - Not complete
var (
	Air = Block{ID: 0, NonSolid: true}
	//Stone
	Stone             = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 0, DataValueMap: map[byte]string{0: "Stone", 1: "Granite", 2: "Polished Granite", 3: "Diorite", 4: "Polished Diorite", 5: "Andesite", 6: "Polished Andesite"}}
	Granite           = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 1, DataValueMap: Stone.DataValueMap}
	Polished_Granite  = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 2, DataValueMap: Stone.DataValueMap}
	Diorite           = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 3, DataValueMap: Stone.DataValueMap}
	Polished_Diorite  = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 4, DataValueMap: Stone.DataValueMap}
	Andesite          = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 5, DataValueMap: Stone.DataValueMap}
	Polished_Andesite = Block{ID: 1, LightFiltered: 15, Sound: "break.stone", DataValue: 6, DataValueMap: Stone.DataValueMap}
	//dirt/grass/cobble
	Grass_Block = Block{ID: 2, NonSolid: true, LightFiltered: 15, Sound: "break.grass"}
	Dirt        = Block{ID: 3, LightFiltered: 1, Sound: "break.dirt", DataValue: 0, DataValueMap: map[byte]string{0: "dirt", 1: "coarse_dirt"}}
	Coarse_Dirt = Block{ID: 3, LightFiltered: 1, Sound: "break.dirt", DataValue: 1, DataValueMap: Dirt.DataValueMap}
	Cobblestone = Block{ID: 4, LightFiltered: 15, Sound: "break.stone"}
	//
	OakPlanks     = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 0, DataValueMap: map[byte]string{0: "oak_planks", 1: "spruce_planks", 2: "birch_planks", 3: "jungle_planks", 4: "acacia_planks", 5: "dark_oak_planks"}}
	SprucePlanks  = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 1, DataValueMap: OakPlanks.DataValueMap}
	BirchPlanks   = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 2, DataValueMap: OakPlanks.DataValueMap}
	JunglePlanks  = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 3, DataValueMap: OakPlanks.DataValueMap}
	AcaciaPlanks  = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 4, DataValueMap: OakPlanks.DataValueMap}
	DarkOakPlanks = Block{ID: 5, LightFiltered: 15, Sound: "break.wood", DataValue: 5, DataValueMap: OakPlanks.DataValueMap}
	//Sapling
	OakSapling = Block{ID: 6, NonSolid: true, Sound: "break.grass"}
	Bedrock    = Block{ID: 7, LightFiltered: 15, Sound: "break.stone"}
	//Liquid
	Water      = Block{ID: 8, NonSolid: true, LightFiltered: 2}
	WaterStill = Block{ID: 9, NonSolid: true, LightFiltered: 2}
	Lava       = Block{ID: 10, NonSolid: true, LightEmitted: 15}
	LavaStill  = Block{ID: 11, NonSolid: true, LightEmitted: 15}
	//
	Sand                = Block{ID: 12, LightFiltered: 15, Sound: "break.sand"}
	Gravel              = Block{ID: 13, LightFiltered: 15, Sound: "break.gravel"}
	GoldOre             = Block{ID: 14, LightFiltered: 15, Sound: "break.stone"}
	IronOre             = Block{ID: 15, LightFiltered: 15, Sound: "break.stone"}
	CoalOre             = Block{ID: 16, LightFiltered: 15, Sound: "break.stone"}
	Log                 = Block{ID: 17, LightFiltered: 15, Sound: "break.wood"}
	Leaves              = Block{ID: 18, LightFiltered: 1, Sound: "break.grass"}
	Sponge              = Block{ID: 19, LightFiltered: 15, Sound: "break.grass"}
	Glass               = Block{ID: 20, Sound: "step.stone"}
	LapisLazuliOre      = Block{ID: 21, LightFiltered: 15, Sound: "break.stone"}
	LapisLazuliBlock    = Block{ID: 22, LightFiltered: 15, Sound: "break.stone"}
	Dispenser           = Block{ID: 23, LightFiltered: 15, Sound: "break.stone"}
	Sandstone           = Block{ID: 24, LightFiltered: 15, Sound: "break.stone"}
	NoteBlock           = Block{ID: 25, LightFiltered: 15, Sound: "break.stone"}
	Bed                 = Block{ID: 26}
	PoweredRail         = Block{ID: 27, NonSolid: true, Sound: "step.stone"}
	DetectorRail        = Block{ID: 28, NonSolid: true, Sound: "step.stone"}
	StickyPiston        = Block{ID: 29, Sound: "break.stone"}
	Cobweb              = Block{ID: 30, NonSolid: true, Sound: "break.stone"}
	TallGrass           = Block{ID: 31, NonSolid: true, Sound: "break.grass"}
	DeadBush            = Block{ID: 32, NonSolid: true, Sound: "break.grass"}
	Piston              = Block{ID: 33, Sound: "break.stone"}
	PistonHead          = Block{ID: 34}
	Wool                = Block{ID: 35, LightFiltered: 15, Sound: "break.cloth"}
	MovedByPiston       = Block{ID: 36}
	Dandelion           = Block{ID: 37, NonSolid: true, Sound: "break.grass"}
	Flower              = Block{ID: 38, NonSolid: true, Sound: "break.grass"}
	BrownMushroom       = Block{ID: 39, NonSolid: true, Sound: "break.grass"}
	RedMushroom         = Block{ID: 40, NonSolid: true, Sound: "break.grass"}
	GoldBlock           = Block{ID: 41, LightFiltered: 15, Sound: "break.stone"}
	IronBlock           = Block{ID: 42, LightFiltered: 15, Sound: "break.stone"}
	DoubleStoneSlab     = Block{ID: 43, LightFiltered: 15, Sound: "break.stone"}
	StoneSlab           = Block{ID: 44, LightFiltered: 1, Sound: "break.stone"}
	Bricks              = Block{ID: 45, LightFiltered: 15, Sound: "break.stone"}
	TNT                 = Block{ID: 46, LightFiltered: 15, Sound: "break.grass"}
	Bookshelf           = Block{ID: 47, LightFiltered: 15, Sound: "break.wood"}
	MossStone           = Block{ID: 48, LightFiltered: 15, Sound: "break.stone"}
	Obsidian            = Block{ID: 49, LightFiltered: 15, Sound: "break.stone"}
	Torch               = Block{ID: 50, NonSolid: true, LightEmitted: 14, Sound: "break.wood"}
	Fire                = Block{ID: 51, LightEmitted: 15}
	MonsterSpawner      = Block{ID: 52}
	OakStairs           = Block{ID: 53, LightFiltered: 15, Sound: "break.wood"}
	Chest               = Block{ID: 54, Sound: "break.wood"}
	RedstoneWire        = Block{ID: 55, NonSolid: true}
	DiamondOre          = Block{ID: 56, LightFiltered: 15, Sound: "break.stone"}
	DiamondBlock        = Block{ID: 57, LightFiltered: 15, Sound: "break.stone"}
	CraftingTable       = Block{ID: 58, LightFiltered: 15, Sound: "break.wood"}
	Wheat               = Block{ID: 59, NonSolid: true, Sound: "break.grass"}
	Farmland            = Block{ID: 60, Sound: "break.gravel"}
	Furnace             = Block{ID: 61, LightFiltered: 15, Sound: "break.stone"}
	BurningFurnace      = Block{ID: 62, LightEmitted: 13, Sound: "break.stone"}
	SignGround          = Block{ID: 63, NonSolid: true, Sound: "break.wood"}
	WoodenDoor          = Block{ID: 64, Sound: "break.wood"}
	Ladders             = Block{ID: 65, Sound: "break.wood"}
	Rails               = Block{ID: 66, Sound: "step.stone"}
	CobblestoneStairs   = Block{ID: 67, LightFiltered: 15, Sound: "break.stone"}
	SignWall            = Block{ID: 68, NonSolid: true, Sound: "break.wood"}
	Lever               = Block{ID: 69, NonSolid: true, Sound: "break.wood"}
	StonePressurePlate  = Block{ID: 70, Sound: "break.stone"}
	IronDoor            = Block{ID: 71, Sound: "break.stone"}
	WoodenPressurePlate = Block{ID: 72, Sound: "break.wood"}
	RedstoneOre         = Block{ID: 73, LightFiltered: 15, Sound: "break.stone"}
	RedstoneOreGlowing  = Block{ID: 74, LightEmitted: 9, Sound: "break.stone"}
	RedstoneTorch       = Block{ID: 75, NonSolid: true, Sound: "break.wood"}
	RedstoneTorchActive = Block{ID: 76, NonSolid: true, LightEmitted: 7, Sound: "break.wood"}
	StoneButton         = Block{ID: 77, NonSolid: true, Sound: "break.stone"}
	SnowLayer           = Block{ID: 78, NonSolid: true, Sound: "break.snow"}
	Ice                 = Block{ID: 79, LightFiltered: 2, Sound: "step.stone"}
	SnowBlock           = Block{ID: 80, LightFiltered: 15, Sound: "break.snow"}
	Cactus              = Block{ID: 81, Sound: "break.cloth"}
	ClayBlock           = Block{ID: 82, LightFiltered: 15, Sound: "break.gravel"}
	SugarCane           = Block{ID: 83, Sound: "break.grass"}
	Jukebox             = Block{ID: 84, LightFiltered: 15, Sound: "break.stone"}
	//Fence
	Fence       = Block{ID: 85, Sound: "break.wood", DataValue: 0, DataValueMap: map[byte]string{0: "Oak Fence", 1: "Spruce Fence", 2: "Birch Fence", 3: "Jungle Fence", 4: "Acaica Fence", 5: "Dark Oak Fence"}}
	SpruceFence = Block{ID: 85, Sound: "break.wood", DataValue: 1, DataValueMap: Fence.DataValueMap}
	BirchFence  = Block{ID: 85, Sound: "break.wood", DataValue: 2, DataValueMap: Fence.DataValueMap}
	JungleFence = Block{ID: 85, Sound: "break.wood", DataValue: 3, DataValueMap: Fence.DataValueMap}
	AcaciaFence = Block{ID: 85, Sound: "break.wood", DataValue: 4, DataValueMap: Fence.DataValueMap}
	DarkFence   = Block{ID: 85, Sound: "break.wood", DataValue: 5, DataValueMap: Fence.DataValueMap}
	//
	Pumpkin                = Block{ID: 86, LightFiltered: 15, Sound: "break.wood"}
	Netherrack             = Block{ID: 87, LightFiltered: 15, Sound: "break.stone"}
	Soulsand               = Block{ID: 88, LightFiltered: 15, Sound: "break.sand"}
	Glowstone              = Block{ID: 89, LightEmitted: 15, Sound: "step.stone"}
	NetherPortal           = Block{ID: 90, NonSolid: true, LightEmitted: 11}
	JackOLantern           = Block{ID: 91, LightEmitted: 15, Sound: "break.wood"}
	CakeBlock              = Block{ID: 92, Sound: "break.cloth"}
	RedstoneRepeater       = Block{ID: 93, Sound: "break.wood"}
	RedstoneRepeaterActive = Block{ID: 94, Sound: "break.wood"}
	StainedGlass           = Block{ID: 95, Sound: "step.stone"}
	Trapdoor               = Block{ID: 96, Sound: "break.wood"}
	MonsterEgg             = Block{ID: 97, LightFiltered: 15, Sound: "break.stone"}
	StoneBricks            = Block{ID: 98, LightFiltered: 15, Sound: "break.stone"}
	BrownMushroomBlock     = Block{ID: 99, LightFiltered: 15, Sound: "break.wood"}
	RedMushroomBlock       = Block{ID: 100, LightFiltered: 15, Sound: "break.wood"}
	IronBars               = Block{ID: 101, Sound: "step.stone"}
	GlassPane              = Block{ID: 102, Sound: "step.stone"}
	Melon                  = Block{ID: 103, LightFiltered: 15, Sound: "break.wood"}
	PumpkinStem            = Block{ID: 104, Sound: "break.grass"}
	MelonStem              = Block{ID: 105, Sound: "break.grass"}
	Vines                  = Block{ID: 106, Sound: "break.grass"}
	//FenceGate                = Block{ID: 107, Sound: "break.wood"}
	BrickStairs              = Block{ID: 108, LightFiltered: 15, Sound: "break.stone"}
	StoneBrickStairs         = Block{ID: 109, LightFiltered: 15, Sound: "break.stone"}
	Mycelium                 = Block{ID: 110, LightFiltered: 15, Sound: "break.grass"}
	Lilypad                  = Block{ID: 111, Sound: "break.grass"}
	NetherBrick              = Block{ID: 112, LightFiltered: 15, Sound: "break.stone"}
	NetherBrickFence         = Block{ID: 113, Sound: "break.stone"}
	NetherBrickStairs        = Block{ID: 114, LightFiltered: 15, Sound: "break.stone"}
	NetherWart               = Block{ID: 115, Sound: "break.grass"}
	EnchantmentTable         = Block{ID: 116, Sound: "break.stone"}
	BrewingStand             = Block{ID: 117, Sound: "break.stone"}
	Cauldron                 = Block{ID: 118, Sound: "break.stone"}
	EndPortal                = Block{ID: 119, LightEmitted: 15}
	EndPortalFrame           = Block{ID: 120, Sound: "break.stone"}
	EndStone                 = Block{ID: 121, LightFiltered: 15, Sound: "break.stone"}
	DragonEgg                = Block{ID: 122}
	RedstoneLamp             = Block{ID: 123, LightEmitted: 15, Sound: "step.stone"}
	RedstoneLampActive       = Block{ID: 124, LightFiltered: 15, Sound: "step.stone"}
	WoodenDoubleSlab         = Block{ID: 125, LightFiltered: 15, Sound: "break.wood"}
	WoodenSlab               = Block{ID: 126, LightFiltered: 1, Sound: "break.wood"}
	CocoaPod                 = Block{ID: 127}
	SandstoneStairs          = Block{ID: 128, LightFiltered: 15, Sound: "break.stone"}
	EmeraldOre               = Block{ID: 129, LightFiltered: 15, Sound: "break.stone"}
	EnderChest               = Block{ID: 130, LightEmitted: 7, Sound: "break.stone"}
	TripwireHook             = Block{ID: 131, Sound: "break.stone"}
	Tripwire                 = Block{ID: 132, Sound: "break.stone"}
	EmeraldBlock             = Block{ID: 133, LightFiltered: 15, Sound: "break.stone"}
	SpruceStairs             = Block{ID: 134, LightFiltered: 15, Sound: "break.wood"}
	BirchStairs              = Block{ID: 135, LightFiltered: 15, Sound: "break.wood"}
	JungleStairs             = Block{ID: 136, LightFiltered: 15, Sound: "break.wood"}
	CommandBlock             = Block{ID: 137, LightFiltered: 15, Sound: "break.stone"}
	Beacon                   = Block{ID: 138, LightEmitted: 15, Sound: "break.stone"}
	CobblestoneWall          = Block{ID: 139, Sound: "break.stone"}
	FlowerPot                = Block{ID: 140, Sound: "break.stone"}
	Carrots                  = Block{ID: 141, Sound: "break.grass"}
	Potatoes                 = Block{ID: 142, Sound: "break.grass"}
	WoodenButton             = Block{ID: 143, Sound: "break.wood"}
	Head                     = Block{ID: 144, Sound: "break.stone"}
	Anvil                    = Block{ID: 145, Sound: "random.anvil.land"}
	TrappedChest             = Block{ID: 146, LightFiltered: 15, Sound: "break.wood"}
	GoldPressurePlate        = Block{ID: 147, Sound: "break.stone"}
	IronPressurePlate        = Block{ID: 148, Sound: "break.stone"}
	RedstoneComparator       = Block{ID: 149, Sound: "break.wood"}
	RedstoneComparatorActive = Block{ID: 150, LightEmitted: 9, Sound: "break.wood"}
	DaylightSensor           = Block{ID: 151, Sound: "break.wood"}
	RedstoneBlock            = Block{ID: 152, LightFiltered: 15, Sound: "break.stone"}
	NetherQuartzOre          = Block{ID: 153, LightFiltered: 15, Sound: "break.stone"}
	Hopper                   = Block{ID: 154, Sound: "break.stone"}
	QuartzBlock              = Block{ID: 155, LightFiltered: 15, Sound: "break.stone"}
	QuartzStairs             = Block{ID: 156, LightFiltered: 15, Sound: "break.stone"}
	ActivatorRail            = Block{ID: 157, Sound: "step.stone"}
	Dropper                  = Block{ID: 158, LightFiltered: 15, Sound: "break.stone"}
	StainedClay              = Block{ID: 159, LightFiltered: 15, Sound: "break.stone"}
	StainedGlassPane         = Block{ID: 160, Sound: "step.stone"}
	Leaves2                  = Block{ID: 161, LightFiltered: 1, Sound: "break.grass"}
	Log2                     = Block{ID: 162, LightFiltered: 15, Sound: "break.wood"}
	AcaciaStairs             = Block{ID: 163, LightFiltered: 15, Sound: "break.wood"}
	DarkStairs               = Block{ID: 164, LightFiltered: 15, Sound: "break.wood"}
	SlimeBlock               = Block{ID: 165, Sound: "break.slime"}
	Barrier                  = Block{ID: 166, Sound: "break.stone"}
	IronTrapDoor             = Block{ID: 167, Sound: "break.stone"}
	Prismarine               = Block{ID: 168, LightFiltered: 15, Sound: "break.stone"}
	SeaLantern               = Block{ID: 169, LightEmitted: 15, Sound: "break.glass"}
	HayBlock                 = Block{ID: 170, LightFiltered: 15, Sound: "break.grass"}
	Carpet                   = Block{ID: 171, Sound: "break.cloth"}
	HardenedClay             = Block{ID: 172, LightFiltered: 15, Sound: "break.stone"}
	CoalBlock                = Block{ID: 173, LightFiltered: 15, Sound: "break.stone"}
	PackedIce                = Block{ID: 174, LightFiltered: 15, Sound: "step.stone"}
	DoublePlant              = Block{ID: 175, NonSolid: true, Sound: "break.grass", DataValueMap: map[byte]string{0: "Sunflower", 1: "Tall Grass"}}
	BannerGround             = Block{ID: 176, Sound: "break.wood"}
	BannerWall               = Block{ID: 177, Sound: "break.wood"}
	InvertedDaylightSensor   = Block{ID: 178, Sound: "break.wood"}
	RedSandStone             = Block{ID: 179, Sound: "break.wood"}
	RedSandStoneStairs       = Block{ID: 180, Sound: "break.wood"}
	DoubleStoneSlab2         = Block{ID: 181, Sound: "break.wood"}
	StoneSlab2               = Block{ID: 182, Sound: "break.wood"}
	//FenceGate
	SpruceFenceGate  = Block{ID: 183, Sound: "break.wood"}
	BirchFenceGate   = Block{ID: 184, Sound: "break.wood"}
	JungleFenceGate  = Block{ID: 185, Sound: "break.wood"}
	DarkFenceGate    = Block{ID: 186, Sound: "break.wood"}
	AcaciaFenceGate  = Block{ID: 187, Sound: "break.wood"}
	FenceGate        = Block{ID: 188, Sound: "break.wood"}
	CrimsonFenceGate = Block{ID: 513, Sound: "break.wood"}
	WarpedFenceGate  = Block{ID: 514, Sound: "break.wood"}
	//Door
	SpruceDoor                = Block{ID: 193, Sound: "break.wood"}
	BirchDoor                 = Block{ID: 194, Sound: "break.wood"}
	JungleDoor                = Block{ID: 195, Sound: "break.wood"}
	AcaciaDoor                = Block{ID: 196, Sound: "break.wood"}
	DarkDoor                  = Block{ID: 197, Sound: "break.wood"}
	EndRod                    = Block{ID: 198, Sound: "break.wood"}
	ChorusPlant               = Block{ID: 199, Sound: "break.stone"}
	ChorusFlower              = Block{ID: 200, Sound: "break.stone"}
	PurPurBlock               = Block{ID: 201, Sound: "break.stone"}
	PurPurPillar              = Block{ID: 202, Sound: "break.stone"}
	PurPurStairs              = Block{ID: 203, Sound: "break.stone"}
	PurPurDoubleSlab          = Block{ID: 204, Sound: "break.stone"}
	PurPurSlab                = Block{ID: 205, Sound: "break.stone"}
	EndBricks                 = Block{ID: 206, Sound: "break.stone"}
	BeetRoot                  = Block{ID: 207, NonSolid: true, Sound: "break.stone"}
	GrassPath                 = Block{ID: 208, Sound: "break.stone"}
	EndGateWay                = Block{ID: 209, NonSolid: true, Sound: "break.stone"}
	RepeatingCommandBlock     = Block{ID: 210, Sound: "break.stone"}
	ChainCommandBlock         = Block{ID: 211, Sound: "break.stone"}
	FrostedIce                = Block{ID: 212, Sound: "break.stone"}
	MagmaBlock                = Block{ID: 213, Sound: "break.stone"}
	NetherWartBlock           = Block{ID: 214, Sound: "break.stone"}
	RedNetherBrick            = Block{ID: 215, Sound: "break.stone"}
	BoneBlock                 = Block{ID: 216, Sound: "break.stone"}
	StructureVoid             = Block{ID: 217, Sound: "break.stone"}
	Observer                  = Block{ID: 218, Sound: "break.stone"}
	WhiteShulkerBox           = Block{ID: 219, Sound: "break.stone"}
	OrangeShulkerBox          = Block{ID: 220, Sound: "break.stone"}
	MagentaShulkerBox         = Block{ID: 221, Sound: "break.stone"}
	LightBlueShulkerBox       = Block{ID: 222, Sound: "break.stone"}
	YellowShulkerBox          = Block{ID: 223, Sound: "break.stone"}
	LimeShulkerBox            = Block{ID: 224, Sound: "break.stone"}
	PinkShulkerBox            = Block{ID: 225, Sound: "break.stone"}
	GrayShulkerBox            = Block{ID: 226, Sound: "break.stone"}
	LightGrayShulkerBox       = Block{ID: 227, Sound: "break.stone"}
	CyanShulkerBox            = Block{ID: 228, Sound: "break.stone"}
	PurpleShulkerBox          = Block{ID: 229, Sound: "break.stone"}
	BlueShulkerBox            = Block{ID: 230, Sound: "break.stone"}
	BrownShulkerBox           = Block{ID: 231, Sound: "break.stone"}
	GreenShulkerBox           = Block{ID: 232, Sound: "break.stone"}
	RedShulkerBox             = Block{ID: 233, Sound: "break.stone"}
	BlackShulkerBox           = Block{ID: 234, Sound: "break.stone"}
	WhiteGlazedTerracotta     = Block{ID: 235, Sound: "break.stone"}
	OrangeGlazedTerracotta    = Block{ID: 236, Sound: "break.stone"}
	MagentaGlazedTerracotta   = Block{ID: 237, Sound: "break.stone"}
	LightBlueGlazedTerracotta = Block{ID: 238, Sound: "break.stone"}
	YellowGlazedTerracotta    = Block{ID: 239, Sound: "break.stone"}
	LimeGlazedTerracotta      = Block{ID: 240, Sound: "break.stone"}
	PinkGlazedTerracotta      = Block{ID: 241, Sound: "break.stone"}
	GrayGlazedTerracotta      = Block{ID: 242, Sound: "break.stone"}
	LightGrayGlazedTerracotta = Block{ID: 243, Sound: "break.stone"}
	CyanGlazedTerracotta      = Block{ID: 244, Sound: "break.stone"}
	PurpleGlazedTerracotta    = Block{ID: 245, Sound: "break.stone"}
	BlueGlazedTerracotta      = Block{ID: 246, Sound: "break.stone"}
	BrownGlazedTerracotta     = Block{ID: 247, Sound: "break.stone"}
	GreenGlazedTerracotta     = Block{ID: 248, Sound: "break.stone"}
	RedGlazedTerracotta       = Block{ID: 249, Sound: "break.stone"}
	BlackGlazedTerracotta     = Block{ID: 250, Sound: "break.stone"}
	Concrete                  = Block{ID: 251, Sound: "break.stone"}
	ConcretePowder            = Block{ID: 252, Sound: "break.sand"}
	StructureBlock            = Block{ID: 255, Sound: "break.stone"}
	Undefined                 = Block{ID: 253, Sound: "break.stone"}
)

//Block stores infomation on a block type
type Block struct {
	//Name          string
	ID            uint16
	LightEmitted  byte
	LightFiltered byte
	Sound         string
	NonSolid      bool
	DataValue     byte
	DataValueMap  map[byte]string //This is only here for conveniance later, maybe removed
	//BlockStates are stored in bytes, the "Attatched" block state is boolean and so the byte is either 0 or 1
	//this is done so every single block doesn't have uneccesary data
	BlockStates map[string]byte
	//Blockstates: - note: the internal numbers maybe different to what the client interprets them as I'm just guessing but will be fixed when testing
	// Age           byte - for plant growth and respawn anchor
	// Attatched     bool - used for Tripwire and TripwireHook
	// Attachment    byte - 0: ceiling, 1: double_wall, 2: floor, 3: single_wall - used for the bell
	// Axis          byte - 0: x-facing, 1: y-facing, 2: z-facing - used for Log, stem, basalt, bone, chain, hay, purpur, quartz and nportal
	// bites         byte - 0~6 - used for cake
	// bottom        bool - used for scaffolding
	// conditional   bool - used for cmd BlockStates
	// delay         byte - used for redstone repeater
	// disarmed      bool - used for Tripwire
	// distance      byte - Scaffolding and Leaves
	// down          bool - used for mushroom blocks and ChorusPlant
	// drag          bool - used for bubble columns
	// east          byte - 0 and 1 for fence and others, 2,3,4 for Redstone: none, side and up
	// eggs          byte - 0-4 how many eggs in block
	// enabled       bool - whether hopper can collect and transfer items
	// extended      bool - pistons
	// eye           bool - whether an eye is in the portal EndPortalFrame
	// face 				 byte - what side of a block the attatched block is on - 0: ceiling, 1: floor, 2: wall
	// facing				 byte - 0: north, 1:east, 2:south, 3:west, 4:up, 5:down
	// half          byte - 0: lower, 1: upper, 2: bottom, 3: top
	// hanging			 bool - used for lantern +  Varients
	// has_book      bool - used for lectern
	////
	/// this should use 3 seperate booleans according to gamepedia but for internal use I'm simplifying it to a byte as 3 bools use 3 bytes
	/// and I find it uneccesary to have them when we can just check how many bottles should show, this may change depending on how test shows
	////
	// has_bottle    byte - used for brewing stand, 0: no bottles, bit 1 (0x01): bottle in slot 1, bit 2 (0x02): bottle in slot 2, bit 3 (0x04): bottle in slot 3
	// has_record		 bool - used for jukebox
	// hatch         byte - used for turtle eggs
	// hinge				 byte - used for door - 0: left, 1: InvertedDaylightSensor
	// in_wall			 bool - used for fence gate
	// instrument    byte - used for noteblock
	// inverted      bool - used for daylight sensor
	// layers        byte - used for SnowLayer
	// leaves        byte - used for bamboo, 0: none, 1: small, 2: large
	// level         byte - used for cauldron-0~3, composter-0~8, water/lava-0~15
	// lit           bool - used for Blast furnace, campfire, furnace, redstone_ore, RedstoneTorch, RedstoneWallTorch, RedstoneLamp, Smoker, SoulCampfire
	// locked        bool - used for redstone repeaters
	// mode          byte - used for redstone componenets and StructureBlock
	// moisture      byte - used for farmland
	// north         bool - used for things
	// note          byte - used for noteblock - 0~24
	// occupied      bool - used for beds
	// open          bool - used for door, FenceGate, Trapdoor and barrel
	// part          byte - used for beds, 0: foot, 1: Head
	// persistent    bool - used for leaves
	// pickles			 byte - used for pickles - 1 to 4
	// power	       byte/nibble - used for redstone - 1 to 15
	// powered       bool - used for redstone
	// rotation      byte/nibble - used for mob heads, banners and signs - 0 to 15
	// shape         byte - too long of a list
	// short         bool - piston head's arm is 4/16th of a block shorter
	// signal_fire   bool - campfire + Varients
	// snowy         bool - grass/podzol/mycelium
	// south         byte - look at east
	// stage         byte - used for saplings + bamboo - 0 to 1
	// triggered     bool - dispenser/Dropper
	//

	// WaterLogged   bool -
}

func GetBlockID(ID int) Block {
	if BlockID == nil {
		InitGlobalID()
	}
	I := Blocks[ID] //BlockID[ID]
	return I
}
