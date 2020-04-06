package chunk

import (
	"VarTool"
	"blocks"
	"fmt"
)

type Location int64

type Chunk struct {
	ChunkPos  Location
	Biomes    []byte
	Blocks    []byte
	NumBlocks uint16
	Modfied   bool
}

var (
	CLC    = make(map[int]*Chunk)
	Biomes = make([]byte, (SectionWidth * SectionLength))
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
	SectionVolume      int16 = (SectionWidth * SectionHeight * SectionLength)
	SectionsNum              = 16
)

func CreateNewChunkSection() { //*Chunk {
	chunk := new(ChunkSection)
	chunk.BlockCount = SectionVolume
	chunk.BitsPerBlock = 8
	chunk.Palette.PaletteLength = 16
	I := make([]int32, 64)
	//I[0:64] = 1
	for i := 0; i < len(I); i++ {
		I[i] = 1
	}
	VarTool.ParseVarIntFromArray(I)
	chunk.Palette.PalleteData = []VarInt{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	chunk.DataArrayLength = 3
	chunk.DataArray = make([]int64, 4096) //chunk.DataArrayLength)
	//return chunk
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
func GenChunk() {
	GenSectionPalette(1, 1) //NumOfBlocks, BlockID
}
