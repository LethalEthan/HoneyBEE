package packet

import (
	"HoneyBEE/world"
)

const GlobalPaletteBPB = 15

//Play 0x22
type ChunkData_CB struct {
	ChunkX              int32
	ChunkZ              int32
	BitMaskLength       int32
	PrimaryBitMask      []int64
	HeightMaps          []byte
	BiomeLength         int32
	Biomes              []int32
	Size                int
	Data                []byte
	NumberBlockEntities int
	BlockEntities       NBT
}

//Play 0x22
type ChunkAndLightData_CB struct {
	// Location
	ChunkX int32
	ChunkZ int32

	HeightMaps          []byte
	Size                int
	Data                []byte
	ChunkSections       []world.ChunkSection
	NumberBlockEntities int
	BlockEntities       []byte
	// Lighting
	TrustEdges          bool
	SkyLightMask        uint64
	BlockLightMask      uint64
	EmptySkyLightMask   uint64
	EmptyBlockLightMask uint64
	SkyLightArrayCount  int32
	SkyLightArray       struct {
		SkyLightArray [2048]byte // always 2048
	}
	BlockLightArrayCount int32
	BlockLightArray      struct {
		BlockLightArray [2048]byte // always 2048
	}
}

func GenerateChunk(X int32, Z int32) ChunkAndLightData_CB {
	C := *new(ChunkAndLightData_CB)
	C.ChunkX = X
	C.ChunkZ = Z
	// C.HeightMaps = GenerateHeightMap(C.Data)
	return C
}

func (Chunk *ChunkData_CB) Encode(X int, Z int) {

}
