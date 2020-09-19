package world

import (
	"chunk"
	"fmt"
	"strconv"
	"sync"
	//"sync"
	"time"

	nibble "github.com/LethalEthan/Go-Nibble"
	logging "github.com/op/go-logging"
)

var RegionMap = make(map[string]region)
var mutex = &sync.RWMutex{} //This is mandatory because of the concurrent and parrallel nature of how HoneyGO works
var Log = logging.MustGetLogger("HoneyGO")

type world struct {
	name       string
	size       int
	RegionData map[string][]region
}

type region struct {
	ID string
	Data []chunk.Chunk //map[string]chunk.Chunk
	//Lock &sync.Mutex{}
}

func Init() {
	//TBD
}

func CreateRegion(X int64, Z int64) {
	//Region is 256*256 Chunks
	ID := strconv.Itoa(int(X)) + "," + strconv.Itoa(int(Z))
	Region := new(region)
	Region.ID = ID
	Region.Data = make([]chunk.Chunk, 65536)
	Chunk := new(chunk.Chunk)
	//
	TNow := time.Now()
	Chunk.Blocks = make([]byte, 16384)
	Chunk.BitsPerBlock = 4
	for i := 0; i < 16384; i++ {
		Chunk.Blocks[i] = nibble.CreateNibbleMerged(1, 1)
	}
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
	Elapse := time.Since(TNow)
	fmt.Print("Finished")
	fmt.Print("Time Taken: ", Elapse)
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
	//Seperate world into 4 sections and have them ran on different goroutines?
	// XZP
	// XZN
	// ZXP
	// ZXN
}

//GetRegionByID - Retrieves Region by the ID safely
func GetRegionByID(ID string) region {
	mutex.RLock()
	R := RegionMap[ID]
	mutex.RUnlock()
	return R
}

//GetRegionByID - Retrieves Region by the X/Z ints safely
func GetRegionByInt(X int, Z int) region {
	ID := strconv.Itoa(int(X)) + "," + strconv.Itoa(int(Z))
	mutex.RLock()
	R := RegionMap[ID]
	mutex.RUnlock()
	return R
}

//var RegionChunkCache = make(map[string]region.Data) -- WIP

//GetChunkFromRegion - Gets the chunk from the region
func GetChunkFromRegion(Region region, CX int, CZ int /*ChunkLoc string*/) chunk.Chunk {
	//Each region contains 0~65536 chunks and each location is dependant on the region location
	//i.e |Region 0,1 the X chunks are 0~255 and the Z chunks are 255~510
	//So if we know the region location we can know all the possible chunk locations without ineffeciently trying to scan through them all
	//fmt.Print(Region.ID)
	X, Z := chunk.COORDSToInts(Region.ID)
	ChunkLocationsX := X * 256
	ChunkLocationsZ := Z * 256
	CLZDelta := ChunkLocationsZ + 255
	CLXDelta := ChunkLocationsX + 255
	var T = 0
	for i := CZ; i < int(CLZDelta); i++ {
		T += 256
	}
	for i := CX; i < /*CX*/ int(CLXDelta); i++ {
		T++
	}
	T = 65535 - T
	return Region.Data[T]
}
