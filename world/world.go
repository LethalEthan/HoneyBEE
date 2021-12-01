package world

import (
	"HoneyBEE/chunk"
	"errors"
	"fmt"
	"sync"

	//"sync"

	nibble "github.com/LethalEthan/Go-Nibble"
	logging "github.com/op/go-logging"
)

var (
	RegionMap           = make(map[RegionID]region)
	mutex               = &sync.RWMutex{} //Allow regions to be sharded?
	Log                 = logging.MustGetLogger("HoneyBEE")
	DefaultFlatArray    []byte
	UninitialisedRegion = new(region)
	UninitialisedChunk  = new(chunk.Chunk)
	RegionNotFound      = errors.New("404: Region not found :(")
	ChunkOOB            = errors.New("Chunk Out Of Bounds")
)

type world struct {
	name       string
	DIM_ID     int
	size       int
	RegionData map[string][]region
}

type region struct {
	ID          RegionID      //string
	Data        []chunk.Chunk //map[string]chunk.Chunk
	Loaded      bool
	ChunkModify chan chunk.Chunk
	//Lock &sync.Mutex{}
}

type RegionID struct {
	X int64
	Z int64
}

func Init() {
	DefaultFlatArray = make([]byte, 16384)
	for i := 0; i < 16384; i++ {
		DefaultFlatArray[i] = nibble.CreateNibbleMerged(1, 1)
	}
}

func CreateRegion(X int64, Z int64) {
	//Region is 256*256 Chunks
	//ID := strconv.Itoa(int(X)) + "," + strconv.Itoa(int(Z))
	Region := new(region)
	RID := new(RegionID)
	RID.X = X
	RID.Z = Z
	Region.ID = *RID
	Region.Data = make([]chunk.Chunk, 65536)
	Region.ChunkModify = make(chan chunk.Chunk, 10)
	Chunk := new(chunk.Chunk)
	//
	Chunk.Blocks = make([]byte, 16384)
	Chunk.BitsPerBlock = 4
	//for i := 0; i < 16384; i++ {
	//Chunk.Blocks[i] = nibble.CreateNibbleMerged(1, 1)
	//}
	Chunk.Blocks = DefaultFlatArray //Only for testing purposes until proper world gen and custom gen is complete which will be far off
	Chunk.NumBlocks = uint16(len(Chunk.Blocks) / int(Chunk.BitsPerBlock))
	cx := X * 256
	cz := Z * 256
	Index := 0
	//Create 256*256 Chunks -- X increments first in the DataArray
	for i := 0; i < 256; i++ {
		for i := 0; i < 256; i++ {
			Chunk.ChunkPosX = cx
			Chunk.ChunkPosZ = cz
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
	PutRegionInMap(*Region)
}

//PutRegionInMap - Safely puts regions into a map to be retrieved whenever
func PutRegionInMap(Region region) {
	mutex.Lock()
	RegionMap[Region.ID] = Region
	mutex.Unlock()
}

//TBD
func CreateMultipleRegions() {
	// type E (int64, int64)
	//TestMap := make(map[]chunk.Chunk)
}

//WIP
func WorldManager() {
	World(0, "test")
}

//WIP
func World(ID uint16, Name string) {
	World := new(world)
	World.name = Name
	//Seperate world into 4 sections and have them run on different goroutines?
	// XZP
	// XZN
	// ZXP
	// ZXN
}

//GetRegionByID - Retrieves Region by the ID safely
func GetRegionByID(ID RegionID) (region, bool, error) {
	mutex.RLock()
	R, bool := RegionMap[ID]
	mutex.RUnlock()
	if bool {
		return R, bool, nil
	}
	return *UninitialisedRegion, bool, RegionNotFound
}

//GetRegionByID - Retrieves Region by the X/Z ints safely
func GetRegionByInt(X int64, Z int64) (region, bool, error) {
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
	return *UninitialisedRegion, bool, RegionNotFound

}

//GetChunkFromRegion - Gets the chunk from the region
func GetChunkFromRegion(Region region, ChunkX int, ChunkZ int) (chunk.Chunk, error) {
	//Each region contains 0~65536 chunks and each location is dependant on the region location
	//i.e |Region 0,1 the X chunks are 0~255 and the Z chunks are 255~510
	//So if we know the region location we can know all the possible chunk locations without ineffeciently trying to scan through them all
	// X := Region.ID.X
	// Z := Region.ID.Z
	var ChunkLocationsX int64
	var ChunkLocationsZ int64
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
		return *UninitialisedChunk, ChunkOOB
	}
	//Math
	T := int(CLZDelta) - ChunkZ      //Minus the Z chunk we try find outta the max possible Zchunk
	T = T * 256                      //Z increments every 256 chunks since each region goes by 256*256 and starts incrementing with X then increments Z for every 256 chunks
	T = T + (int(CLXDelta) - ChunkX) //Minus the X chunk we try find outta the max possible Xchunk then add to var T (position in array)
	T = 65535 - T                    //Take away the position of the chunk from the max length of the array
	return Region.Data[T], nil
}

func GetChunkRangeFromRegion(Region region) (int64, int64, int64, int64) {
	MinCX := Region.ID.X * 256
	MinCZ := Region.ID.Z * 256
	MaxCX := MinCX + 255
	MaxCZ := MinCZ + 255
	return MinCX, MinCZ, MaxCX, MaxCZ
}

//RunRegion - Does the region logic and changes the chunks on change, function name will probably change.
func (R region) RunRegion() {
	for {
		select {
		case CM := <-R.ChunkModify:
			RegionChunk, err := GetChunkFromRegion(R, int(CM.ChunkPosX), int(CM.ChunkPosZ))
			if err != nil {
				panic(err)
			}
			_ = RegionChunk
		}
	}
}
