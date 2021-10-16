package items

import (
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("HoneyBEE")

var ItemID map[float32]string

//Initialise ItemID map
func InitItemID() map[float32]string {
	//Due to items having a sub ID such as charcoal: 263:1, we'll have to utilise a float32/string map
	ItemID = map[float32]string{
		//Iron Tools
		256: "Iron_Shovel",
		257: "Iron_Picaxe",
		258: "Iron_Axe",
		//Idk Category
		259: "Flint_And_Steel",
		260: "Apple",
		//Bow/Arrow
		261: "Bow",
		262: "Arrow",
		//Valuables
		263:    "Coal",
		263.01: "Charcoal",
		264:    "Diamond",
		265:    "Iron_Ingot",
		266:    "Gold_Ingot",
		267:    "Iron_Sword",
		//Wooden Tools
		268: "Wooden_Sword",
		269: "Wooden_Shovel",
		270: "Wooden_Pickaxe",
		271: "Wooden_Axe",
		//Stone Tools
		272: "Stone_Sword",
		273: "Stone_Shovel",
		274: "Stone_Pickaxe",
		275: "Stone_Axe",
		//Diamond Tools
		276: "Diamond_Sword",
		277: "Diamond_Shovel",
		278: "Diamond_Pickaxe",
		279: "Diamond_Axe",
		280: "Stick",
		//Stewie
		281: "Bowl",
		282: "Mushroom_Stew",
		//Golden Tools
		283: "Golden_Sword",
		284: "Golden_Shovel",
		285: "Golden_Pickaxe",
		286: "Golden_Axe",
		287: "String",
		288: "Feather",
		289: "GunPowder",
		//Hoes
		290: "Wooden_Hoe",
		291: "Stone_Hoe",
		292: "Iron_Hoe",
		293: "Diamond_Hoe",
		294: "Golden_Hoe",
		//Bread Stuff
		295: "Wheat_Seeds",
		296: "Wheat",
		297: "Bread",
		//Leather Armour
		298: "Leather_Helmet",
		299: "Leather_Tunic",
		300: "Leather_Pants",
		301: "Leather_Boots",
		//Chainmail Armour
		302: "Chainmail_Helmet",
		303: "Chainmail_Chestplate",
		304: "Chainmail_Leggings",
		305: "Chainmail_Boots",
		//Iron Armour
		306: "Iron_Helmet",
		307: "Iron_Chestplate",
		308: "Iron_Leggings",
		309: "Iron_Boots",
		//Diamond Armour
		310: "Diamond_Helmet",
		311: "Diamond_Chestplate",
		312: "Diamond_Leggings",
		313: "Diamond_Boots",
		//Gold Armour
		314: "Golden_Helemet",
		315: "Golden_Chestplate",
		316: "Golden_Leggings",
		317: "Golden_Boots",
		//Random/Food/Items
		318:    "Flint",
		319:    "Raw_Porkchop",
		320:    "Cooked_Porkchop",
		321:    "Painting",
		322:    "Golden_Apple",
		322.01: "Enchanted_Golden_Apple",
		323:    "Sign",
		324:    "Oak_Door",
		//Bukkits
		325: "Bucket",
		326: "Water_Bucket",
		327: "Lava_Bucket",
		//Misc
		328: "Minecart",
		329: "Saddle",
		330: "Iron_Door",
		331: "Redstone",
		332: "Snowball",
		333: "Boat",
		334: "Leather",
		//Thanks for being the odd 1 out
		335: "Milk_Bucket",
		336: "Brick",
		337: "Clay_Ball",
		338: "Sugar_Cane",
		339: "Paper",
		340: "Book",
		341: "Slime_Ball",
		//Other Minecart
		342: "Chest_Minecart",
		343: "Furnace_Minecart",
		//Misc
		344: "Egg",
		345: "Compass",
		346: "Fishing_Rod",
		347: "Clock",
		348: "Glowstone_Dust",
		//Fish
		349:    "Raw_Fish",
		349.01: "Raw_Salmon",
		349.02: "Clownfish",
		349.03: "Pufferfish",
		350:    "Cooked_Fish",
		350.01: "Cooked_Salmon",
		//Dye
		351:    "Ink_Sack",
		351.01: "Rose_Red",
		351.02: "Cactus_Green",
		351.03: "Coco_Beans",
		351.04: "Lapis_Lazuli",
		351.05: "Purple_Dye",
		351.06: "Cyan_Dye",
		351.07: "Light Gray Dye",
		351.08: "Gray_Dye",
		351.09: "Pink_Dye",
		351.10: "Lime_Dye",
		351.11: "Dandelion_Yellow",
		351.12: "Light_Blue_Dye",
		351.13: "Magenta_Dye",
		351.14: "Orange_Dye",
		351.15: "Bone_Meal", //Doot
		//
		352: "Bone",
		353: "Sugar",
		354: "Cake",
		355: "Bed",
		356: "Redstone_Repeater",
		357: "Cookie",
		358: "Map", //Filled_Map?
		359: "Shears",
		360: "Melon",
		//Seeds
		361: "Pumpkin_Seeds",
		362: "Melon_Seeds",
		//Meat Food
		363: "Raw_Beef",
		364: "Cooked_Beef",
		365: "Raw_Chicken",
		366: "Cooked_Chicken",
		367: "Rotten_Flesh",
		368: "Ender_Pearl",
		369: "Blaze_Rod",
		370: "Ghast_Tear",
		371: "Gold_Nugget",
		372: "Nether_Wart",
		373: "Potion",
		374: "Glass_Bottle",
		375: "Spider_Eye",
		376: "Fermented_Spider_Eye",
		377: "Blaze_Powder",
		378: "Magma_Cream",
		379: "Brewing_Stand",
		380: "Cauldron",
		381: "Ender_Eye",
		382: "Speckled_Melon",
		//SpawnEggs
		383.04: "Spawn_Elder_Guardian",
		383.05: "Spawn_Wither_Skeleton",
		383.06: "Spawn_Stray",
		383.23: "Spawn_Husk",
		383.27: "Spawn_Zombie_Villager",
		383.28: "Spawn_Skeleton_Horse",
		383.29: "Spawn_Zombie_Horse",
		383.31: "Spawn_Donkey",
		383.32: "Spawn_Mule",
		383.34: "Spawn_Evoker",
		383.35: "Spawn_Vex",
		383.36: "Spawn_Vindicator",
		383.50: "Spawn_Creeper",
		383.51: "Spawn_Skeleton",
		383.52: "Spawn_Spider",
	}
	return ItemID
}

// func GetItemID(ID int) {
// 	if ItemID == nil {
// 		InitItemID()
// 	}
// 	I := ItemID[ID]
// 	log.Info("Block: ", I)
// }

type Item struct {
	Name  string
	ID    uint16
	Count byte
	//tag
}

type ItemEntity struct {
	Age         int16
	Health      byte
	PickupDelay int16    //this probably will be removed as it depends on ticks
	Owner       string   //used by the give command, will be available to use with honeycomb so plugins can tailor items to a specific players
	Owners      []string //Declare multiple players that can pick up this item such as a team of players
	Thrower     string
	ItemS       Item
}
