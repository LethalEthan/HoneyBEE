package chunk

import (
	"blocks"
	"fmt"
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

var (
	//TBD once everything is in place
	//CLCX   = make(map[int]*Chunk)
	//CLCY   = make(map[int]*Chunk)
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
	SectionVolume      int16 = 4096
	SectionsNum              = 16
)

//Note: Currently has a lot of debug and testing stuff, is not finalised
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
	//VarTool.ParseVarIntFromArray(I)
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
func GenChunk(CX int64, CY int64) Chunk {
	//GenSectionPalette(1, 1) //NumOfBlocks, BlockID
	Chunk := new(Chunk)
	Chunk.ChunkPosX = CX
	Chunk.ChunkPosZ = CY
	Chunk.NumSecs = 2
	//t := SectionVolume * int16(Chunk.NumSecs)
	Chunk.Blocks = make([]byte, 512)
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
	CS.BitsPerBlock = 8
}
