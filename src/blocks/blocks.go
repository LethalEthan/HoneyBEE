package blocks

import logging "github.com/op/go-logging"

var log = logging.MustGetLogger("HoneyGO")

//BlockData
var (
	Air                       = Block{ID: 0}
	Stone                     = Block{ID: 1, LightFiltered: 15, Sound: "break.stone"}
	Grass                     = Block{ID: 2, LightFiltered: 15, Sound: "break.grass"}
	Dirt                      = Block{ID: 3, LightFiltered: 15, Sound: "break.gravel"}
	Cobblestone               = Block{ID: 4, LightFiltered: 15, Sound: "break.stone"}
	WoodenPlanks              = Block{ID: 5, LightFiltered: 15, Sound: "break.wood"}
	Sapling                   = Block{ID: 6, Sound: "break.grass"}
	Bedrock                   = Block{ID: 7, LightFiltered: 15, Sound: "break.stone"}
	Water                     = Block{ID: 8, LightFiltered: 2}
	WaterStill                = Block{ID: 9, LightFiltered: 2}
	Lava                      = Block{ID: 10, LightEmitted: 15}
	LavaStill                 = Block{ID: 11, LightEmitted: 15}
	Sand                      = Block{ID: 12, LightFiltered: 15, Sound: "break.sand"}
	Gravel                    = Block{ID: 13, LightFiltered: 15, Sound: "break.gravel"}
	GoldOre                   = Block{ID: 14, LightFiltered: 15, Sound: "break.stone"}
	IronOre                   = Block{ID: 15, LightFiltered: 15, Sound: "break.stone"}
	CoalOre                   = Block{ID: 16, LightFiltered: 15, Sound: "break.stone"}
	Log                       = Block{ID: 17, LightFiltered: 15, Sound: "break.wood"}
	Leaves                    = Block{ID: 18, LightFiltered: 1, Sound: "break.grass"}
	Sponge                    = Block{ID: 19, LightFiltered: 15, Sound: "break.grass"}
	Glass                     = Block{ID: 20, Sound: "step.stone"}
	LapisLazuliOre            = Block{ID: 21, LightFiltered: 15, Sound: "break.stone"}
	LapisLazuliBlock          = Block{ID: 22, LightFiltered: 15, Sound: "break.stone"}
	Dispenser                 = Block{ID: 23, LightFiltered: 15, Sound: "break.stone"}
	Sandstone                 = Block{ID: 24, LightFiltered: 15, Sound: "break.stone"}
	NoteBlock                 = Block{ID: 25, LightFiltered: 15, Sound: "break.stone"}
	Bed                       = Block{ID: 26}
	PoweredRail               = Block{ID: 27, Sound: "step.stone"}
	DetectorRail              = Block{ID: 28, Sound: "step.stone"}
	StickyPiston              = Block{ID: 29, Sound: "break.stone"}
	Cobweb                    = Block{ID: 30, Sound: "break.stone"}
	TallGrass                 = Block{ID: 31, Sound: "break.grass"}
	DeadBush                  = Block{ID: 32, Sound: "break.grass"}
	Piston                    = Block{ID: 33, Sound: "break.stone"}
	PistonHead                = Block{ID: 34}
	Wool                      = Block{ID: 35, LightFiltered: 15, Sound: "break.cloth"}
	MovedByPiston             = Block{ID: 36}
	Dandelion                 = Block{ID: 37, Sound: "break.grass"}
	Flower                    = Block{ID: 38, Sound: "break.grass"}
	BrownMushroom             = Block{ID: 39, Sound: "break.grass"}
	RedMushroom               = Block{ID: 40, Sound: "break.grass"}
	GoldBlock                 = Block{ID: 41, LightFiltered: 15, Sound: "break.stone"}
	IronBlock                 = Block{ID: 42, LightFiltered: 15, Sound: "break.stone"}
	DoubleStoneSlab           = Block{ID: 43, LightFiltered: 15, Sound: "break.stone"}
	StoneSlab                 = Block{ID: 44, LightFiltered: 1, Sound: "break.stone"}
	Bricks                    = Block{ID: 45, LightFiltered: 15, Sound: "break.stone"}
	TNT                       = Block{ID: 46, LightFiltered: 15, Sound: "break.grass"}
	Bookshelf                 = Block{ID: 47, LightFiltered: 15, Sound: "break.wood"}
	MossStone                 = Block{ID: 48, LightFiltered: 15, Sound: "break.stone"}
	Obsidian                  = Block{ID: 49, LightFiltered: 15, Sound: "break.stone"}
	Torch                     = Block{ID: 50, LightEmitted: 14, Sound: "break.wood"}
	Fire                      = Block{ID: 51, LightEmitted: 15}
	MonsterSpawner            = Block{ID: 52}
	OakStairs                 = Block{ID: 53, LightFiltered: 15, Sound: "break.wood"}
	Chest                     = Block{ID: 54, Sound: "break.wood"}
	RedstoneWire              = Block{ID: 55}
	DiamondOre                = Block{ID: 56, LightFiltered: 15, Sound: "break.stone"}
	DiamondBlock              = Block{ID: 57, LightFiltered: 15, Sound: "break.stone"}
	CraftingTable             = Block{ID: 58, LightFiltered: 15, Sound: "break.wood"}
	Wheat                     = Block{ID: 59, Sound: "break.grass"}
	Farmland                  = Block{ID: 60, Sound: "break.gravel"}
	Furnace                   = Block{ID: 61, LightFiltered: 15, Sound: "break.stone"}
	BurningFurnace            = Block{ID: 62, LightEmitted: 13, Sound: "break.stone"}
	SignGround                = Block{ID: 63, Sound: "break.wood"}
	WoodenDoor                = Block{ID: 64, Sound: "break.wood"}
	Ladders                   = Block{ID: 65, Sound: "break.wood"}
	Rails                     = Block{ID: 66, Sound: "step.stone"}
	CobblestoneStairs         = Block{ID: 67, LightFiltered: 15, Sound: "break.stone"}
	SignWall                  = Block{ID: 68, Sound: "break.wood"}
	Lever                     = Block{ID: 69, Sound: "break.wood"}
	StonePressurePlate        = Block{ID: 70, Sound: "break.stone"}
	IronDoor                  = Block{ID: 71, Sound: "break.stone"}
	WoodenPressurePlate       = Block{ID: 72, Sound: "break.wood"}
	RedstoneOre               = Block{ID: 73, LightFiltered: 15, Sound: "break.stone"}
	RedstoneOreGlowing        = Block{ID: 74, LightEmitted: 9, Sound: "break.stone"}
	RedstoneTorch             = Block{ID: 75, Sound: "break.wood"}
	RedstoneTorchActive       = Block{ID: 76, LightEmitted: 7, Sound: "break.wood"}
	StoneButton               = Block{ID: 77, Sound: "break.stone"}
	SnowLayer                 = Block{ID: 78, Sound: "break.snow"}
	Ice                       = Block{ID: 79, LightFiltered: 2, Sound: "step.stone"}
	SnowBlock                 = Block{ID: 80, LightFiltered: 15, Sound: "break.snow"}
	Cactus                    = Block{ID: 81, Sound: "break.cloth"}
	ClayBlock                 = Block{ID: 82, LightFiltered: 15, Sound: "break.gravel"}
	SugarCane                 = Block{ID: 83, Sound: "break.grass"}
	Jukebox                   = Block{ID: 84, LightFiltered: 15, Sound: "break.stone"}
	Fence                     = Block{ID: 85, Sound: "break.wood"}
	Pumpkin                   = Block{ID: 86, LightFiltered: 15, Sound: "break.wood"}
	Netherrack                = Block{ID: 87, LightFiltered: 15, Sound: "break.stone"}
	Soulsand                  = Block{ID: 88, LightFiltered: 15, Sound: "break.sand"}
	Glowstone                 = Block{ID: 89, LightEmitted: 15, Sound: "step.stone"}
	NetherPortal              = Block{ID: 90, LightEmitted: 11}
	JackOLantern              = Block{ID: 91, LightEmitted: 15, Sound: "break.wood"}
	CakeBlock                 = Block{ID: 92, Sound: "break.cloth"}
	RedstoneRepeater          = Block{ID: 93, Sound: "break.wood"}
	RedstoneRepeaterActive    = Block{ID: 94, Sound: "break.wood"}
	StainedGlass              = Block{ID: 95, Sound: "step.stone"}
	Trapdoor                  = Block{ID: 96, Sound: "break.wood"}
	MonsterEgg                = Block{ID: 97, LightFiltered: 15, Sound: "break.stone"}
	StoneBricks               = Block{ID: 98, LightFiltered: 15, Sound: "break.stone"}
	BrownMushroomBlock        = Block{ID: 99, LightFiltered: 15, Sound: "break.wood"}
	RedMushroomBlock          = Block{ID: 100, LightFiltered: 15, Sound: "break.wood"}
	IronBars                  = Block{ID: 101, Sound: "step.stone"}
	GlassPane                 = Block{ID: 102, Sound: "step.stone"}
	Melon                     = Block{ID: 103, LightFiltered: 15, Sound: "break.wood"}
	PumpkinStem               = Block{ID: 104, Sound: "break.grass"}
	MelonStem                 = Block{ID: 105, Sound: "break.grass"}
	Vines                     = Block{ID: 106, Sound: "break.grass"}
	FenceGate                 = Block{ID: 107, Sound: "break.wood"}
	BrickStairs               = Block{ID: 108, LightFiltered: 15, Sound: "break.stone"}
	StoneBrickStairs          = Block{ID: 109, LightFiltered: 15, Sound: "break.stone"}
	Mycelium                  = Block{ID: 110, LightFiltered: 15, Sound: "break.grass"}
	Lilypad                   = Block{ID: 111, Sound: "break.grass"}
	NetherBrick               = Block{ID: 112, LightFiltered: 15, Sound: "break.stone"}
	NetherBrickFence          = Block{ID: 113, Sound: "break.stone"}
	NetherBrickStairs         = Block{ID: 114, LightFiltered: 15, Sound: "break.stone"}
	NetherWart                = Block{ID: 115, Sound: "break.grass"}
	EnchantmentTable          = Block{ID: 116, Sound: "break.stone"}
	BrewingStand              = Block{ID: 117, Sound: "break.stone"}
	Cauldron                  = Block{ID: 118, Sound: "break.stone"}
	EndPortal                 = Block{ID: 119, LightEmitted: 15}
	EndPortalFrame            = Block{ID: 120, Sound: "break.stone"}
	EndStone                  = Block{ID: 121, LightFiltered: 15, Sound: "break.stone"}
	DragonEgg                 = Block{ID: 122}
	RedstoneLamp              = Block{ID: 123, LightEmitted: 15, Sound: "step.stone"}
	RedstoneLampActive        = Block{ID: 124, LightFiltered: 15, Sound: "step.stone"}
	WoodenDoubleSlab          = Block{ID: 125, LightFiltered: 15, Sound: "break.wood"}
	WoodenSlab                = Block{ID: 126, LightFiltered: 1, Sound: "break.wood"}
	CocoaPod                  = Block{ID: 127}
	SandstoneStairs           = Block{ID: 128, LightFiltered: 15, Sound: "break.stone"}
	EmeraldOre                = Block{ID: 129, LightFiltered: 15, Sound: "break.stone"}
	EnderChest                = Block{ID: 130, LightEmitted: 7, Sound: "break.stone"}
	TripwireHook              = Block{ID: 131, Sound: "break.stone"}
	Tripwire                  = Block{ID: 132, Sound: "break.stone"}
	EmeraldBlock              = Block{ID: 133, LightFiltered: 15, Sound: "break.stone"}
	SpruceStairs              = Block{ID: 134, LightFiltered: 15, Sound: "break.wood"}
	BirchStairs               = Block{ID: 135, LightFiltered: 15, Sound: "break.wood"}
	JungleStairs              = Block{ID: 136, LightFiltered: 15, Sound: "break.wood"}
	CommandBlock              = Block{ID: 137, LightFiltered: 15, Sound: "break.stone"}
	Beacon                    = Block{ID: 138, LightEmitted: 15, Sound: "break.stone"}
	CobblestoneWall           = Block{ID: 139, Sound: "break.stone"}
	FlowerPot                 = Block{ID: 140, Sound: "break.stone"}
	Carrots                   = Block{ID: 141, Sound: "break.grass"}
	Potatoes                  = Block{ID: 142, Sound: "break.grass"}
	WoodenButton              = Block{ID: 143, Sound: "break.wood"}
	Head                      = Block{ID: 144, Sound: "break.stone"}
	Anvil                     = Block{ID: 145, Sound: "random.anvil.land"}
	TrappedChest              = Block{ID: 146, LightFiltered: 15, Sound: "break.wood"}
	GoldPressurePlate         = Block{ID: 147, Sound: "break.stone"}
	IronPressurePlate         = Block{ID: 148, Sound: "break.stone"}
	RedstoneComparator        = Block{ID: 149, Sound: "break.wood"}
	RedstoneComparatorActive  = Block{ID: 150, LightEmitted: 9, Sound: "break.wood"}
	DaylightSensor            = Block{ID: 151, Sound: "break.wood"}
	RedstoneBlock             = Block{ID: 152, LightFiltered: 15, Sound: "break.stone"}
	NetherQuartzOre           = Block{ID: 153, LightFiltered: 15, Sound: "break.stone"}
	Hopper                    = Block{ID: 154, Sound: "break.stone"}
	QuartzBlock               = Block{ID: 155, LightFiltered: 15, Sound: "break.stone"}
	QuartzStairs              = Block{ID: 156, LightFiltered: 15, Sound: "break.stone"}
	ActivatorRail             = Block{ID: 157, Sound: "step.stone"}
	Dropper                   = Block{ID: 158, LightFiltered: 15, Sound: "break.stone"}
	StainedClay               = Block{ID: 159, LightFiltered: 15, Sound: "break.stone"}
	StainedGlassPane          = Block{ID: 160, Sound: "step.stone"}
	Leaves2                   = Block{ID: 161, LightFiltered: 1, Sound: "break.grass"}
	Log2                      = Block{ID: 162, LightFiltered: 15, Sound: "break.wood"}
	AcaciaStairs              = Block{ID: 163, LightFiltered: 15, Sound: "break.wood"}
	DarkStairs                = Block{ID: 164, LightFiltered: 15, Sound: "break.wood"}
	SlimeBlock                = Block{ID: 165, Sound: "break.slime"}
	Barrier                   = Block{ID: 166, Sound: "break.stone"}
	IronTrapDoor              = Block{ID: 167, Sound: "break.stone"}
	Prismarine                = Block{ID: 168, LightFiltered: 15, Sound: "break.stone"}
	SeaLantern                = Block{ID: 169, LightEmitted: 15, Sound: "break.glass"}
	HayBlock                  = Block{ID: 170, LightFiltered: 15, Sound: "break.grass"}
	Carpet                    = Block{ID: 171, Sound: "break.cloth"}
	HardenedClay              = Block{ID: 172, LightFiltered: 15, Sound: "break.stone"}
	CoalBlock                 = Block{ID: 173, LightFiltered: 15, Sound: "break.stone"}
	PackedIce                 = Block{ID: 174, LightFiltered: 15, Sound: "step.stone"}
	DoublePlant               = Block{ID: 175, Sound: "break.grass"}
	BannerGround              = Block{ID: 176, Sound: "break.wood"}
	BannerWall                = Block{ID: 177, Sound: "break.wood"}
	InvertedDaylightSensor    = Block{ID: 178, Sound: "break.wood"}
	RedSandStone              = Block{ID: 179, Sound: "break.wood"}
	RedSandStoneStairs        = Block{ID: 180, Sound: "break.wood"}
	DoubleStoneSlab2          = Block{ID: 181, Sound: "break.wood"}
	StoneSlab2                = Block{ID: 182, Sound: "break.wood"}
	SpruceFenceGate           = Block{ID: 183, Sound: "break.wood"}
	BirchFenceGate            = Block{ID: 184, Sound: "break.wood"}
	JungleFenceGate           = Block{ID: 185, Sound: "break.wood"}
	DarkFenceGate             = Block{ID: 186, Sound: "break.wood"}
	AcaciaFenceGate           = Block{ID: 187, Sound: "break.wood"}
	SpruceFence               = Block{ID: 188, Sound: "break.wood"}
	BirchFence                = Block{ID: 189, Sound: "break.wood"}
	JungleFence               = Block{ID: 190, Sound: "break.wood"}
	DarkFence                 = Block{ID: 191, Sound: "break.wood"}
	AcaciaFence               = Block{ID: 192, Sound: "break.wood"}
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
	BeetRoot                  = Block{ID: 207, Sound: "break.stone"}
	GrassPath                 = Block{ID: 208, Sound: "break.stone"}
	EndGateWay                = Block{ID: 209, Sound: "break.stone"}
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

var BlockID map[int]string

//Initialise BlockID map
func InitGlobalID() map[int]string {
	BlockID = map[int]string{
		0:   "Air",
		1:   "Stone",
		2:   "Grass",
		3:   "Dirt",
		4:   "Cobblestone",
		5:   "WoodenPlanks",
		6:   "Sapling",
		7:   "Bedrock",
		8:   "Water",
		9:   "WaterStill",
		10:  "Lava",
		11:  "LavaStill",
		12:  "Sand",
		13:  "Gravel",
		14:  "GoldOre",
		15:  "IronOre",
		16:  "CoalOre",
		17:  "Log",
		18:  "Leaves",
		19:  "Sponge",
		20:  "Glass",
		21:  "LapisLazuliOre",
		22:  "LapisLazuliBlock",
		23:  "Dispenser",
		24:  "Sandstone",
		25:  "NoteBlock",
		26:  "Bed",
		27:  "PoweredRail",
		28:  "DetectorRail",
		29:  "StickyPiston",
		30:  "Cobweb",
		31:  "TallGrass", //Includes Fern
		32:  "DeadBush",
		33:  "Piston",
		34:  "PistonHead",
		35:  "Wool", //Includes Coloured Varients
		36:  "MovedByPiston",
		37:  "Dandelion",
		38:  "Flower", //Includes: Oxeye Daisy; Pink, White, Orange, red tulip; Azure; Allium; Blue Orchid and poppy
		39:  "BrownMushroom",
		40:  "RedMushroom",
		41:  "GoldBlock",
		42:  "IronBlock",
		43:  "DoubleStoneSlab", //This includes petrified oakslab
		44:  "StoneSlab",       //This includes petrified oakslab
		45:  "Bricks",
		46:  "TNT",
		47:  "Bookshelf",
		48:  "MossStone",
		49:  "Obsidian",
		50:  "Torch",
		51:  "Fire",
		52:  "MonsterSpawner",
		53:  "OakStairs",
		54:  "Chest",
		55:  "RedstoneWire",
		56:  "DiamondOre",
		57:  "DiamondBlock",
		58:  "CraftingTable",
		59:  "Wheat",
		60:  "Farmland",
		61:  "Furnace",
		62:  "BurningFurnace",
		63:  "SignGround",
		64:  "WoodenDoor",
		65:  "Ladders",
		66:  "Rails",
		67:  "CobblestoneStairs",
		68:  "SignWall",
		69:  "Lever",
		70:  "StonePressurePlate",
		71:  "IronDoor",
		72:  "WoodenPressurePlate",
		73:  "RedstoneOre",
		74:  "RedstoneOreGlowing",
		75:  "RedstoneTorch",
		76:  "RedstoneTorchActive",
		77:  "StoneButton",
		78:  "SnowLayer",
		79:  "Ice",
		80:  "SnowBlock",
		81:  "Cactus",
		82:  "ClayBlock",
		83:  "SugarCane",
		84:  "Jukebox",
		85:  "Fence",
		86:  "Pumpkin",
		87:  "Netherrack",
		88:  "Soulsand",
		89:  "Glowstone",
		90:  "NetherPortal",
		91:  "JackOLantern",
		92:  "CakeBlock",
		93:  "RedstoneRepeater",
		94:  "RedstoneRepeaterActive",
		95:  "StainedGlass",
		96:  "Trapdoor",
		97:  "MonsterEgg",
		98:  "StoneBricks",
		99:  "BrownMushroomBlock",
		100: "RedMushroomBlock",
		101: "IronBars",
		102: "GlassPane",
		103: "Melon",
		104: "PumpkinStem",
		105: "MelonStem",
		106: "Vines",
		107: "FenceGate",
		108: "BrickStairs",
		109: "StoneBrickStairs",
		110: "Mycelium",
		111: "Lilypad",
		112: "NetherBrick",
		113: "NetherBrickFence",
		114: "NetherBrickStairs",
		115: "NetherWart",
		116: "EnchantmentTable",
		117: "BrewingStand",
		118: "Cauldron",
		119: "EndPortal",
		120: "EndPortalFrame",
		121: "EndStone",
		122: "DragonEgg",
		123: "RedstoneLamp",
		124: "RedstoneLampActive",
		125: "WoodenDoubleSlab",
		126: "WoodenSlab",
		127: "CocoaPod",
		128: "SandstoneStairs",
		129: "EmeraldOre",
		130: "EnderChest",
		131: "TripwireHook",
		132: "Tripwire",
		133: "EmeraldBlock",
		134: "SpruceStairs",
		135: "BirchStairs",
		136: "JungleStairs",
		137: "CommandBlock",
		138: "Beacon",
		139: "CobblestoneWall",
		140: "FlowerPot",
		141: "Carrots",
		142: "Potatoes",
		143: "WoodenButton",
		144: "Head",
		145: "Anvil",
		146: "TrappedChest",
		147: "GoldPressurePlate",
		148: "IronPressurePlate",
		149: "RedstoneComparator",
		150: "RedstoneComparatorActive",
		151: "DaylightSensor",
		152: "RedstoneBlock",
		153: "NetherQuartzOre",
		154: "Hopper",
		155: "QuartzBlock",
		156: "QuartzStairs",
		157: "ActivatorRail",
		158: "Dropper",
		159: "StainedClay",
		160: "StainedGlassPane",
		161: "Leaves2",
		162: "Log2",
		163: "AcaciaStairs",
		164: "DarkStairs",
		165: "SlimeBlock",
		166: "Barrier",
		167: "IronTrapDoor",
		168: "Prismarine",
		169: "SeaLantern",
		170: "HayBlock",
		171: "Carpet",
		172: "HardenedClay",
		173: "CoalBlock",
		174: "PackedIce",
		175: "DoublePlant",
		176: "BannerGround",
		177: "BannerWall",
		178: "InvertedDaylightSensor",
		179: "RedSandStone",
		180: "RedSandStoneStairs",
		181: "DoubleStoneSlab2",
		182: "StoneSlab2",
		183: "SpruceFenceGate",
		184: "BirchFenceGate",
		185: "JungleFenceGate",
		186: "DarkFenceGate",
		187: "AcaciaFenceGate",
		188: "SpruceFence",
		189: "BirchFence",
		190: "JungleFence",
		191: "DarkFence",
		192: "AcaciaFence",
		193: "SpruceDoor",
		194: "BirchDoor",
		195: "JungleDoor",
		196: "AcaciaDoor",
		197: "DarkDoor",
		198: "EndRod",
		199: "ChorusPlant",
		200: "ChorusFlower",
		201: "PurPurBlock",
		202: "PurPurPillar",
		203: "PurPurStairs",
		204: "PurPurDoubleSlab",
		205: "PurPurSlab",
		206: "EndBricks",
		207: "BeetRoot",
		208: "GrassPath",
		209: "EndGateWay",
		210: "RepeatingCommandBlock",
		211: "ChainCommandBlock",
		212: "FrostedIce",
		213: "MagmaBlock",
		214: "NetherWartBlock",
		215: "RedNetherBrick",
		216: "BoneBlock",
		217: "StructureVoid",
		218: "Observer",
		219: "WhiteShulkerBox",
		220: "OrangeShulkerBox",
		221: "MagentaShulkerBox",
		222: "LightBlueShulkerBox",
		223: "YellowShulkerBox",
		224: "LimeShulkerBox",
		225: "PinkShulkerBox",
		226: "GrayShulkerBox",
		227: "LightGrayShulkerBox",
		228: "CyanShulkerBox",
		229: "PurpleShulkerBox",
		230: "BlueShulkerBox",
		231: "BrownShulkerBox",
		232: "GreenShulkerBox",
		233: "RedShulkerBox",
		234: "BlackShulkerBox",
		235: "WhiteGlazedTerracotta", //Lol does anyone actually use these blocks?
		236: "OrangeGlazedTerracotta",
		237: "MagentaGlazedTerracotta",
		238: "LightBlueGlazedTerracotta",
		239: "YellowGlazedTerracotta",
		240: "LimeGlazedTerracotta",
		241: "PinkGlazedTerracotta",
		242: "GrayGlazedTerracotta",
		243: "LightGrayGlazedTerracotta",
		244: "CyanGlazedTerracotta",
		245: "PurpleGlazedTerracotta",
		246: "BlueGlazedTerracotta",
		247: "BrownGlazedTerracotta",
		248: "GreenGlazedTerracotta",
		249: "RedGlazedTerracotta",
		250: "BlackGlazedTerracotta",
		251: "Concrete",
		252: "ConcretePowder",
		253: "Undefined",
		254: "Undefined",
		255: "StructureBlock",
		256: "Undefined",
	}
	return BlockID
}

//Block stores infomation on a block type
type Block struct {
	ID            uint16
	LightEmitted  byte
	LightFiltered byte
	Sound         string
	Solid         bool
}

func GetBlockID(ID int) {
	if BlockID == nil {
		InitGlobalID()
	}
	I := BlockID[ID]
	log.Info("Block: ", I)
}
