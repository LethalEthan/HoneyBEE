package chunk

import (
	logging "github.com/op/go-logging"
)

type Location int64 //TBD

//Around 16kb, 64kb with 1 section filled with stone nibbles
type Chunk struct {
	ChunkPosX    int64
	ChunkPosZ    int64
	Biomes       []byte
	Blocks       []byte // Replace with sections later
	NumBlocks    uint16
	NumSecs      uint8
	Modfied      bool
	BitsPerBlock byte
	IsFlat       bool //This flag states if the chunk follows a repeatable pattern such as a flat world to reduce the amount of bytes being used to signify the same amount of blocks -- To be implemented
}

var (
	Log = logging.MustGetLogger("HoneyBEE")
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
