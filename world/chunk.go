package world

import (
	"HoneyBEE/nbt"

	"github.com/google/uuid"
)

var Regions = make(map[RegionID]region)

const GlobalPaletteBPB = 13

type PalletedDataContainer struct {
	BitsPerEntry    byte
	Palette         Pallete
	DataArrayLength int32
	DataArray       []uint64 // Compacted data array
}

type Pallete struct {
	PalleteLength int32 // 0 will mean a direct pallete
	Pallete       []int32
}

type ChunkSection struct {
	BlockCount  int16
	BlockStates PalletedDataContainer
	Biomes      PalletedDataContainer
}

type ChunkColumn struct {
	ChunkX        int
	ChunkZ        int
	Data          []ChunkSection
	BlockEntities []BlockEnity
}

type BlockEnity struct {
	ID       uuid.UUID
	PackedXZ byte
	Y        int16
	Type     int32
	Data     nbt.NBTEncoder
}

// func CreateChunkSection() (CS ChunkSection) {
// 	CS.BlockCount = 4096
// 	_, CS.BlockStates.BitsPerEntry = FindBitsPerBlock(CS.BlockStates.Palette.Pallete)
// 	return CS
// }

func GenerateHeightMap(chunk []byte) []byte { // Currently static
	NBTE := nbt.CreateNBTEncoder()
	BA := make([]int64, 37)
	for i := 0; i < len(BA); i++ { // Setup static heightmap
		if i < 36 { // last 3 bytes are unused
			BA[i] |= 0x40 << 55
			BA[i] |= 0x40 << 46
			BA[i] |= 0x40 << 37
			BA[i] |= 0x40 << 28
			BA[i] |= 0x40 << 19
			BA[i] |= 0x40 << 10
			BA[i] |= 0x40 << 1
		} else { // Whether it should be LSB or MSB is something I do not know
			BA[i] |= 0x40 << 28
			BA[i] |= 0x40 << 19
			BA[i] |= 0x40 << 10
			BA[i] |= 0x40 << 1
		}
	}
	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", BA))
	NBTE.EndCompoundTag()
	return NBTE.Encode()
}

// FindBitsPerBlock - finds the different blockstates and returns them along with the bits per block
func FindBitsPerBlock(data []int32) (map[int]int32, int) {
	var blockstates = make(map[int]int32)
	for i, v := range data { // search through chunk
		if _, ok := blockstates[int(v)]; !ok { // if data value does not exists in map
			blockstates[i] = v // add value to map
		}
	}
	switch k := len(blockstates); {
	case k <= 1:
		return blockstates, 0
	case k <= 16:
		return blockstates, 4
	case k <= 32:
		return blockstates, 5
	case k <= 64:
		return blockstates, 6
	case k <= 128:
		return blockstates, 7
	case k <= 256:
		return blockstates, 8
	case k <= 512:
		return blockstates, 9
	case k <= 1024:
		return blockstates, 10
	case k <= 2048:
		return blockstates, 11
	case k <= 4096:
		return blockstates, 12
	case k <= 8192:
		return blockstates, 13
	case k <= 16384:
		return blockstates, 14
	default:
		return blockstates, GlobalPaletteBPB
	}
}

func CreatePalette(BitsPerBlock int, BlockStates map[int]uint64) {
	PDC := new(PalletedDataContainer)
	PDC.BitsPerEntry = byte(BitsPerBlock)
	PDC.Palette.PalleteLength = int32(len(BlockStates))

}

func GenerateRegion(x, z int) {
	Region := new(region)
	Region.ID.X = x
	Region.ID.Z = z
	Region.Data = make([]ChunkColumn, 0, 65535)
}
