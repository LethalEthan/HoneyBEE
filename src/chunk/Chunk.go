package chunk

import (
	"blocks"
	"fmt"
	"strconv"
	"strings"

	nibble "github.com/LethalEthan/Go-Nibble"
	logging "github.com/op/go-logging"
)

type Location int64

//Around 16kb
type Chunk struct {
	ChunkPosX    int64
	ChunkPosZ    int64
	Biomes       []byte
	Blocks       []byte // Replace with sections later
	NumBlocks    uint16
	NumSecs      uint8
	Modfied      bool
	BitsPerBlock byte
	Repeat       byte //This flag states if the chunk follows a repeatable pattern such as a flat world to reduce the amount of bytes being used to signify the same amount of blocks -- To be implemented
}

var (
	Biomes = make([]byte, (SectionWidth * SectionLength))
	//ChunkData - 2D slice [X][Z]
	//ChunkData  = make([][][]int64, 512)
	//ChunkData uses a string/Chunk map to store the byte arrays of blocks which
	ChunkData = make(map[string]Chunk)
	Log       = logging.MustGetLogger("HoneyGO")
	COORDS    string
)

const (
	GlobalBitsPerBlock uint8 = 14
	MinBitsPerBlock    uint8 = 4
	MaxBitsPerBlock    uint8 = 8
	ChunkWidth         int   = 16
	ChunkLength        int   = 16
	ChunkHeight        int   = 256
	ChunkVolume        int   = 65536
	SectionHeight      int16 = 16
	SectionWidth       int16 = 16
	SectionLength      int16 = 16
	SectionVolume      int16 = 4096
	SectionsNum              = 16
)

//Note: Currently has a lot of debug and testing stuff, is not finalised
func CreateNewChunkSection() *ChunkSection { //*Chunk {
	chunk := new(ChunkSection)
	chunk.BlockCount = SectionVolume * 2
	chunk.BitsPerBlock = 4
	//chunk.Palette.PaletteLength = 16
	// I := make([]int32, 64)
	// //I[0:64] = 1
	// for i := 0; i < len(I); i++ {
	// 	I[i] = 1
	// }
	//VarTool.ParseVarIntFromArray(I)
	//chunk.Palette.PalleteData = []VarInt{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	chunk.DataArrayLength = 256
	chunk.DataArray = make([]int64, 256) //chunk.DataArrayLength)
	chunk.DataArray = BuildDataArray(0, chunk)
	for i := 0; i < 256; i++ {
		chunk.DataArray[i] = chunk.DataArray[0]
	}
	fmt.Print(chunk.DataArray)
	return chunk
}

//Very WIP
func BuildDataArray(Index int, chunk *ChunkSection) []int64 {
	var t int
	t = 60
	for i := 0; i < 16; i++ {
		chunk.DataArray[Index] |= 1 << t
		//chunk.DataArray[Index] = 1 << t
		t = t - 4
		fmt.Print("\n")
		fmt.Printf("%064b", chunk.DataArray[Index])
	}
	return chunk.DataArray
}

func Getblocks() uint16 {
	// C.NumBlocks = 16 * 16 * 16
	return 4096
}

func GenSectionPalette(NB int, BID int) {
	for i := 0; i > len(blocks.Blocks); i++ {
		//blocks.Blocks.ID
		//B2 := blocks.BlockID
		// if B2 == BID {
		// 	print("Block Found: ", blocks.BlockID[i])
		// }
	}
}

func GlobalPalette() map[int]string {
	BlockID := blocks.InitGlobalID()
	fmt.Print("BlockID 0: ", BlockID[0])
	// fmt.Print("Air", blocks.Air)
	return BlockID
}

//Testing using only stone blocks
func GenChunk(CX int64, CY int64) Chunk {
	//GenSectionPalette(1, 1) //NumOfBlocks, BlockID
	Chunk := new(Chunk)
	Chunk.ChunkPosX = CX
	Chunk.ChunkPosZ = CY
	Chunk.NumSecs = 2
	//t := SectionVolume * int16(Chunk.NumSecs)
	Chunk.Blocks = make([]byte, 256)
	for i := 0; i <= 7; i++ {
		Chunk.Blocks[i] = 1
	}
	return *Chunk
}

func SendChunkPacket(Chunk *Chunk, BitMask VarInt, DAL int) {
	CP := new(ChunkPacket)
	CP.ChunkX = Chunk.ChunkPosX
	CP.ChunkZ = Chunk.ChunkPosZ
	CP.FullChunk = true
	CP.PBitMask = BitMask
	//CP.HeightMaps = []int64{}
	CP.Size = VarInt(DAL)
	CS := new(ChunkSection)
	CS.BitsPerBlock = 4
}

func COORDSToInts(COORDS string) (int64, int64) {
	i := strings.Index(COORDS, ",")
	//X = COORDS[:i]
	X, _ := strconv.ParseInt(COORDS[:i], 10, 64)
	Z, _ := strconv.ParseInt(COORDS[i+1:len(COORDS)], 10, 64)
	return X, Z
}

func IntsToCOORDS(X int, Z int) string {
	COORDS := strconv.Itoa(X) + "," + strconv.Itoa(Z)
	return COORDS
}

//BuildChunk - WIP
func BuildChunk(X int, Z int, BPB byte) Chunk {
	COORDS = strconv.Itoa(X) + "," + strconv.Itoa(Z)
	fmt.Print(COORDS)
	switch BPB {
	case 4:
		Log.Debug("BPB: 4")
		C := new(Chunk)
		C.NumBlocks = 4096
		C.Blocks = make([]byte, 16384)
		for i := 0; i < len(C.Blocks); i++ {
			C.Blocks[i] = nibble.CreateNibbleMerged(1, 1)
		}
		return *C
	case 8:
		C := new(Chunk)
		return *C
	default:
		C := new(Chunk)
		return *C
	}
}

///
///   TO BE REMOVED, REPLACED BY NIBBLE PACKAGE
///
//PackByteToNibble - Packs bytes (8bits) in a chunk to nibbles (4bits)
func PackByteToNibble(B []byte) []byte {
	//Make Blank slice
	B2 := make([]byte, 0)
	for i := 0; i < len(B)-1; i += 2 {
		M := nibble.CreateNibbleMerged(B[i], B[i+1]) //Create Nibbles
		B2 = append(B2, M)                           //Append nibbles to byte array
	}
	B = nil //Set previous array to nil
	return B2
}

func UnPackByteToNibble(B []byte) []byte {
	//Create Blank slice
	B2 := make([]byte, 0)
	for i := 0; i < 32768; i++ {
		//Read nibbles
		M := nibble.ReadNibble1(B[i])
		M2 := nibble.ReadNibble2(B[i])
		//Append em
		B2 = append(B2, M)
		B2 = append(B2, M2)
	}
	B = nil //Set previous array to nil
	return B2
}

///
///
///

func CompactByteArrayToint64(BA []byte) []int64 {
	//var ChunkS int64
	var CA []int64
	var Index int
	//ChunkS = int64(TStone)
	for j := 0; j < len(BA); j += 64 {
		for i := 0; i <= 64; i += 8 {
			CA[Index] = int64(BA[i+j] + BA[i+j+1] + BA[i+j+2] + BA[i+j+3] + BA[i+j+4] + BA[i+j+5] + BA[i+j+6] + BA[i+j+7])
			//fmt.Printf("%064b", ChunkS)
		}
		Index++
	}
	return CA
}
