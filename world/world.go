package world

import (
	"errors"
	"fmt"
	"sync"

	nibble "github.com/LethalEthan/Go-Nibble"
	logging "github.com/op/go-logging"
)

const NetherDIMID = -1
const OverworldDIMID = 0
const TheEndDIMID = 1

var WorldRegistry = make(map[int]world)
var WorldIDCounter int = 10
var worldcountermutex sync.Mutex

var (
	RegionMap           = make(map[RegionID]region)
	mutex               = &sync.RWMutex{} //Allow regions to be sharded?
	Log                 = logging.MustGetLogger("HoneyBEE")
	DefaultFlatArray    []byte
	UninitialisedRegion = new(region)
	UninitialisedChunk  = new(ChunkColumn)
)

const (
	RegionNotFound             = "404: Region not found :("
	ChunkOOB                   = "Chunk Out Of Bounds"
	ErrWorldEntryAlreadyExists = "World entry already exists: overworld, nether and the_end are added to the registry upon start up. You can not delete these from the registry only deny access or change world generation variables."
	ErrWorldRegistryFull       = "World entry is full, no more worlds can be added!"
)

type world struct {
	name      string
	dim_id    int
	worldsize int
	seed      int64
	typeof    int8 //overworld, nether, end
	// RegionData map[string][]region
}

type world_options struct {
	Seed                 int64 // Set Make_Seed to true for new pseudorandom seed
	FixedTime            int64 // optional
	LogicalHeight        int32
	Coordinate_Scale     float32
	Ambient_Light        float32
	Max_Height           int16
	Min_Height           int16
	IsHardcore           bool
	IsFlat               bool
	Piglin_Safe          bool
	Natural              bool
	Has_FixedTime        bool
	Has_Raids            bool
	Has_Skylight         bool
	Has_Ceiling          bool
	Ultra_Warm           bool
	Bed_Works            bool
	Respawn_Anchor_Works bool
	Make_Seed            bool // Set true for seed to be generared
}

func AddWorldtoRegistry(name string, size int, typeof int8, seed int64) error {
	world := new(world)
	world.name = name
	world.seed = seed
	world.typeof = typeof
	worldcountermutex.Lock()
	switch typeof {
	case -1:
		if _, e := WorldRegistry[-1]; !e && name == "nether" {
			world.dim_id = NetherDIMID
			WorldRegistry[-1] = *world
		}
	case 0:
		if _, e := WorldRegistry[0]; !e && name == "overworld" {
			world.dim_id = OverworldDIMID
			WorldRegistry[0] = *world
		}
	case 1:
		if _, e := WorldRegistry[1]; !e && name == "the_end" {
			world.dim_id = TheEndDIMID
			WorldRegistry[1] = *world
		}
	default:
		if WorldIDCounter >= 10 && WorldIDCounter <= -10 {
			WorldRegistry[WorldIDCounter] = *world
			WorldIDCounter++
		} else {
			return errors.New(ErrWorldRegistryFull)
		}
	}
	worldcountermutex.Unlock()
	return nil
}

func Init() {
	DefaultFlatArray = make([]byte, 16384)
	for i := 0; i < 16384; i++ {
		DefaultFlatArray[i] = nibble.CreateNibbleMerged(1, 1)
	}
}

func CreateRegion(X int, Z int) {
	//Region is 256*256 Chunks
	Region := new(region)
	// RID := new(RegionID)
	// RID.X = X
	// RID.Z = Z
	Region.ID.X = X
	Region.ID.Z = Z
	Region.Data = make([]ChunkColumn, 65536)
	// Region.ChunkModify = make(chan chunk.Chunk, 10)
	Chunk := new(ChunkColumn)
	//
	// Chunk.Blocks = make([]byte, 16384)
	// Chunk.BitsPerBlock = 4
	//for i := 0; i < 16384; i++ {
	//Chunk.Blocks[i] = nibble.CreateNibbleMerged(1, 1)
	//}
	// Chunk.Blocks = DefaultFlatArray //Only for testing purposes until proper world gen and custom gen is complete which will be far off
	// Chunk.NumBlocks = uint16(len(Chunk.Blocks) / int(Chunk.BitsPerBlock))
	cx := X * 256
	cz := Z * 256
	Index := 0
	//Create 256*256 Chunks -- X increments first in the DataArray
	for i := 0; i < 256; i++ {
		for i := 0; i < 256; i++ {
			Chunk.ChunkX = cx
			Chunk.ChunkZ = cz
			Region.Data[Index] = *Chunk
			Index++
			if cx < 0 {
				cx--
			} else {
				cx++
			}
		}
		cx = X * 256 //reset down to the original number
		if cz < 0 {
			cz--
		} else {
			cz++
		}
	}
}

//GetRegionByID - Retrieves Region by the ID safely
func GetRegionByID(ID RegionID) (region, bool, error) {
	mutex.RLock()
	R, bool := RegionMap[ID]
	mutex.RUnlock()
	if bool {
		return R, bool, nil
	}
	return *UninitialisedRegion, bool, errors.New(RegionNotFound)
}

//GetRegionByID - Retrieves Region by the X/Z ints safely
func GetRegionByInt(X int, Z int) (region, bool, error) {
	ID := new(RegionID)
	ID.X = X
	ID.Z = Z
	mutex.RLock()
	R, bool := RegionMap[*ID]
	mutex.RUnlock()
	//Simple check to see if it's initialised
	if bool {
		return R, bool, nil
	}
	return *UninitialisedRegion, bool, errors.New(RegionNotFound)

}

//GetChunkFromRegion - Gets the chunk from the region
func GetChunkFromRegion(Region region, ChunkX int, ChunkZ int) (ChunkColumn, error) {
	//Each region contains 0~65536 chunks and each location is dependant on the region location
	//i.e |Region 0,1 the X chunks are 0~255 and the Z chunks are 255~510
	//So if we know the region location we can know all the possible chunk locations without ineffeciently trying to scan through them all
	var ChunkLocationsX int
	var ChunkLocationsZ int
	//Make negative numbers positive so the logic still works
	if Region.ID.X < 0 { //If X is negative
		ChunkLocationsX = -Region.ID.X * 256 //Min XChunk Co-ord
		ChunkX = -ChunkX
	} else {
		ChunkLocationsX = Region.ID.X * 256 //Min XChunk Co-ord
	}
	if Region.ID.Z < 0 {
		ChunkLocationsZ = -Region.ID.Z * 256 //Min ZChunk Co-ord
		ChunkZ = -ChunkZ
	} else {
		ChunkLocationsZ = Region.ID.Z * 256 //Min ZChunk Co-ord
	}
	CLZDelta := ChunkLocationsZ + 255 //Max ZChunk Co-ord
	CLXDelta := ChunkLocationsX + 255 //Max XChunk Co-ord
	if ChunkX > int(CLXDelta) || ChunkZ > int(CLZDelta) || ChunkX < int(ChunkLocationsX) || ChunkZ < int(ChunkLocationsX) {
		fmt.Print("Warning: Chunk OOB")
		return *UninitialisedChunk, errors.New(ChunkOOB)
	}
	//Math
	T := int(CLZDelta) - ChunkZ      //Minus the Z chunk we try find outta the max possible Zchunk
	T = T * 256                      //Z increments every 256 chunks since each region goes by 256*256 and starts incrementing with X then increments Z for every 256 chunks
	T = T + (int(CLXDelta) - ChunkX) //Minus the X chunk we try find outta the max possible Xchunk then add to var T (position in array)
	T = 65535 - T                    //Take away the position of the chunk from the max length of the array
	return Region.Data[T], nil
}

func GetChunkRangeFromRegion(Region region) (int, int, int, int) {
	MinCX := Region.ID.X * 256
	MinCZ := Region.ID.Z * 256
	MaxCX := MinCX + 255
	MaxCZ := MinCZ + 255
	return MinCX, MinCZ, MaxCX, MaxCZ
}
