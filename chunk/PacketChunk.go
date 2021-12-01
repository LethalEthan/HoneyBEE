package chunk

type ChunkPacket struct {
	ChunkX                int64
	ChunkZ                int64
	FullChunk             bool
	PBitMask              int32
	HeightMaps            []byte
	Biomes                []byte //Optional array of integers
	Size                  int32
	Data                  []byte //Number of elements is equal to the number of bits set in PBitMask. Sections Sent from top to bottom, i.e first section is Y=0 to Y=15
	NumberOfBlockEntities int32
	BlockEntities         []byte
}
