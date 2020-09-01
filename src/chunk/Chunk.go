package chunk

import (
	"blocks"
	"fmt"

	nibble "github.com/LethalEthan/Go-Nibble"
)

type Location int64

type Chunk struct {
	ChunkPosX int64
	ChunkPosZ int64
	Biomes    []byte
	Blocks    []byte
	NumBlocks uint16
	NumSecs   uint8
	Modfied   bool
}

type CMap struct {
	PX int
	PZ int
}

var (
	//TBD once everything is in place
	//CLCX   = make(map[int]*Chunk)
	//CLCY   = make(map[int]*Chunk)
	Biomes = make([]byte, (SectionWidth * SectionLength))
	//ChunkData - 2D slice [X][Z]
	ChunkData = make([][][]byte, 512)
)

const (
	GlobalBitsPerBlock uint8 = 14
	MinBitsPerBlock    uint8 = 4
	MaxBitsPerBlock    uint8 = 8
	ChunkWidth         int   = 16
	ChunkLength        int   = 16
	ChunkHeight        int   = 256
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

func BuildChunk(X int, Z int) {
	//test := [10][4096]byte{}
	//fmt.Print(test)
	ChunkData[X] = make([][]byte, 100)
	ChunkData[X][Z] = make([]byte, 32768)
	layer := []byte{}
	for i := 0; i < 65536; i++ {
		layer = append(layer, 1)
	}
	//print(layer)
	fmt.Print(len(layer))
	//layer := PackByteToNibble(ChunkData[X][Z])
	layer = PackByteToNibble(layer)
	ChunkData[X][Z] = layer
	//ChunkData[X][Y] = []byte{0, 0, 135, 90, 7, 4, 7, 44, 73, 10}
	fmt.Print( /*ChunkData[X][Z],*/ "Length: ", len(ChunkData[X][Z]))
	UnPackByteToNibble(ChunkData[X][Z])
}

//PackByteToNibble - Packs bytes (8bits) in a chunk to nibbles (4bits)
func PackByteToNibble(B []byte) []byte {
	//Make Blank slice
	B2 := make([]byte, 0)
	for i := 0; i < len(B)-1; i += 2 {
		M := nibble.CreateNibbleMerged(B[i], B[i+1]) //Create Nibbles
		B2 = append(B2, M)                           //Append nibbles to byte array
	}
	B = nil //Set previoud array to nil
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
