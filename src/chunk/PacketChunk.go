package chunk

type VarInt int32
type NBT byte //Temporary

type ChunkPacket struct {
	ChunkX                int
	ChunkY                int
	FullChunk             bool
	PBitMask              VarInt
	HeightMaps            NBT
	Biomes                []byte //Optional array of integers
	Size                  VarInt
	Data                  []byte //Number of elements is equal to the number of bits set in PBitMask. Sections Sent from top to bottom, i.e first section is Y=0 to Y=15
	NumberOfBlockEntities VarInt
	BlockEntities         []NBT //ArrayOfNBT
}
