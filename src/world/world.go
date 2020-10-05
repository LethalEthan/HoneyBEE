package world

import (
	"chunk"
	"config"
	"errors"
	"fmt"
	"sync"
	//"sync"
	"time"

	nibble "github.com/LethalEthan/Go-Nibble"
	logging "github.com/op/go-logging"
)

var (
	RegionMap           = make(map[RegionID]region)
	mutex               = &sync.RWMutex{} //This is mandatory because of the concurrent and parrallel nature of how HoneyGO works
	Log                 = logging.MustGetLogger("HoneyGO")
	DEBUG               bool
	DefaultFlatArray    []byte
	UninitialisedRegion = new(region)
	UninitialisedChunk  = new(chunk.Chunk)
	RegionNotFound      = errors.New("404: Region not found :(")
	ChunkOOB            = errors.New("Chunk Out Of Bounds")
)

type world struct {
	name       string
	size       int
	RegionData map[string][]region
}

type region struct {
	ID   RegionID      //string
	Data []chunk.Chunk //map[string]chunk.Chunk
	//Lock &sync.Mutex{}
}

type RegionID struct {
	X int64
	Z int64
}

func Init() {
	Config := config.GetConfig()
	if Config.Server.DEBUG {
		DEBUG = true
	}
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
	Chunk := new(chunk.Chunk)
	//
	var TNow time.Time
	if DEBUG {
		TNow = time.Now()
	}
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
			cx++
		}
		cx = X * 256 //reset down to the original number
		cz++
	}
	//RegionMap[ID] = *Region //Don't directly insert into map due to race concerns when the world generation pool is active
	PutRegionInMap(*Region)
	///
	///Note: Find a way to store regions without strings as they cosume shit tonnes of ram in a map -\_0_0_/-
	///
	if DEBUG {
		Elapse := time.Since(TNow)
		fmt.Print("\nFinished creating: ", RID, " Time Taken: ", Elapse)
	}
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

//var RegionChunkCache = make(map[string]region.Data) -- WIP

//GetChunkFromRegion - Gets the chunk from the region
func GetChunkFromRegion(Region region, CX int, CZ int) (chunk.Chunk, error) {
	//Each region contains 0~65536 chunks and each location is dependant on the region location
	//i.e |Region 0,1 the X chunks are 0~255 and the Z chunks are 255~510
	//So if we know the region location we can know all the possible chunk locations without ineffeciently trying to scan through them all
	// X := Region.ID.X
	// Z := Region.ID.Z
	ChunkLocationsX := Region.ID.X * 256 //Min XChunk Co-ord
	ChunkLocationsZ := Region.ID.Z * 256 //Min ZChunk Co-ord
	CLZDelta := ChunkLocationsZ + 255    //Max ZChunk Co-ord
	CLXDelta := ChunkLocationsX + 255    //Max XChunk Co-ord
	if CX > int(CLXDelta) || CZ > int(CLZDelta) {
		fmt.Print("Warning: Chunk OOB")
		return *UninitialisedChunk, ChunkOOB
	}
	//Math
	T := int(CLZDelta) - CZ //50
	T = T * 256
	T = T + (int(CLXDelta) - CX)
	T = 65535 - T
	return Region.Data[T], nil
}
