package chunk

import (
	"fmt"
	"math"
)

//Chunk Sections are 16*16*16

type ChunkSection struct {
	BlockCount   int16    //Number of "non-air" blocks (anything other than air,cave air and void air)
	BitsPerBlock byte     //How many bits are used to encode a block
	Palette      struct { //Depends on format, Indirect OR Direct
		GlobalPallete []int16
		PaletteLength VarInt
		PalleteData   []VarInt
	}
	DataArrayLength VarInt  //Number of longs in Array
	DataArray       []int64 //Compacted list of 4096 (16*16*16) indicies pointing to state ID's in the palette
}

type PalleteI interface {
	GetBitsPerBlock() uint8
}

func (CS ChunkSection) GetBitsPerBlock() uint8 {
	return 14
}
func (CC ChunkColumn) GetBitsPerBlock() uint8 {
	return 14
}

func DirectPallete(Chunk *ChunkSection) {
	if Chunk.BitsPerBlock >= 9 {
		X := math.Log2(float64(Chunk.BlockCount))
		fmt.Print("Maths: ", X)
	}
}

func GenChunkSection(BC int16) *ChunkSection {
	CS := new(ChunkSection)
	CS.BlockCount = BC
	CS.BitsPerBlock = 8
	//CS.Palette.PaletteLength = VarTool.EncodeVarLong(1)
	//CS.Palette.PalleteData = VarInt.Encode()

	CS.DataArray = make([]int64, 512)
	for i := 0; i < 512; i++ {
		CS.DataArray[i] = 0xFFFFFFFFFFFFFFF
	}
	DL := VarInt(len(CS.DataArray))
	CS.DataArrayLength = DL
	return CS
}

func (v VarInt) Encode( /*v int64*/ ) (vi []byte) {
	num := uint32(v)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		vi = append(vi, byte(b))
		if num == 0 {
			break
		}
	}
	return
}
