package server

import (
	"HoneyBEE/nbt"
	"HoneyBEE/packet"

	"github.com/panjf2000/gnet"
)

//
// A temporary mock up of chunk loading until I figure everything out properly
//

func ChunkLoad(c gnet.Conn) error {
	// C := CreateChunk()
	//
	// tmpCL := packet.CreateWriterWithCapacity(16) //temp chunk location writer
	// tmpCL.WriteInt(0)
	// tmpCL.WriteInt(0)
	//
	// tmpUL := packet.CreateWriterWithCapacity(16) //temp chunk location writer
	// tmpUL.WriteVarInt(0)
	// tmpUL.WriteVarInt(0)
	//
	//UL := CreateLightData()
	//Send chunk
	if err := c.AsyncWrite(CreateLightData(0, 0)); err != nil {
		c.Close() //send
		return err
	}
	if err := c.AsyncWrite(CreateChunk(0, 0)); err != nil { //CreateChunk(0, 0)
		c.Close() //send
		return err
	}
	// Log.Debug("Data of chunk", CreateChunk(0, 0))
	// Log.Debug("Packet Size: ", len(PW.GetData()))
	var X int32 //j
	var Z int32 //i
	for Z = 0; Z < 12; Z++ {
		for X = 0; X < 12; X++ {
			if X != 0 || Z != 0 {
				if err := c.AsyncWrite(CreateLightData(X, -Z)); err != nil {
					c.Close()
					return err
				}
				if err := c.AsyncWrite(CreateChunk(X, Z)); err != nil { // Send chunk
					c.Close()
					return err
				}
				if Z == 0 {
					if err := c.AsyncWrite(CreateLightData(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateChunk(-X, Z)); err != nil {
						c.Close()
						return err
					}
				}
				if X == 0 {
					if err := c.AsyncWrite(CreateLightData(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateChunk(X, -Z)); err != nil {
						c.Close()
						return err
					}
				}
				if X != 0 && Z != 0 {
					if err := c.AsyncWrite(CreateLightData(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateChunk(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateLightData(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateChunk(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateLightData(-X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := c.AsyncWrite(CreateChunk(-X, -Z)); err != nil {
						c.Close()
						return err
					}
				}
			}
		}
	}
	Log.Debug("Sent chunks!")
	PW := packet.CreatePacketWriterWithCapacity(0x20, 128) //Initialise World Border
	PW.WriteDouble(0.0)
	PW.WriteDouble(0.0)
	PW.WriteDouble(348.0)
	PW.WriteDouble(348.0)
	PW.WriteVarLong(0)
	PW.WriteVarInt(29999984)
	PW.WriteVarInt(0)
	PW.WriteVarInt(0)
	Log.Debug("Sent Init World Border")
	if err := c.AsyncWrite(PW.GetPacket()); err != nil {
		c.Close()
		return err
	}
	return nil
}

// func TRYChunkLoad() {
// 	Chunk := new(packet.ChunkData_CB)
// 	Chunk.ChunkX = 0
// 	Chunk.ChunkZ = 0
// 	Chunk.BitMaskLength = 1
// 	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 0b00000001)
// 	NBTE := nbt.CreateNBTEncoder()
// 	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", make([]int64, 37)))
// 	NBTE.Encode()
// 	Log.Debug("X: ", 0, "Z: ", 0)
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			if i != 0 || j != 0 {
// 				Log.Debug("X: ", j, "Z: ", i)
// 				if i == 0 {
// 					Log.Debug("X: ", -j, "Z: ", i)
// 				}
// 				if j == 0 {
// 					Log.Debug("X: ", j, "Z: ", -i)
// 				}
// 				if j != 0 && i != 0 {
// 					Log.Debug("X: ", j, "Z: ", -i)
// 					Log.Debug("X: ", -j, "Z: ", i)
// 					Log.Debug("X: ", -j, "Z: ", -i)
// 				}
// 			}
// 		}
// 	}
// }

func WriteBlocks(BitsPerBlock int, Blocks []int16) { // To-do
	var EncodedBlock []int64
	for i := 0; i < len(Blocks); i++ {
		EncodedBlock[i] = 0
	}
}

func WriteBiomeID(ID int32) []int32 {
	B := make([]int32, 1024)
	return B
}

func CreateStaticHeightMap() nbt.NBTEncoder {
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
	// Log.Debug("HeightMapL:", len(BA))
	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", BA))
	NBTE.EndCompoundTag()
	return NBTE
}

func CreateLightData(X, Z int32) []byte {
	UL := packet.CreatePacketWriterWithCapacity(0x25, 5000)
	UL.WriteVarInt(X)
	UL.WriteVarInt(Z)
	UL.WriteBoolean(true)
	// Sky light mask
	UL.WriteVarInt(1)
	UL.WriteULong(1)
	// Block Light mask
	UL.WriteVarInt(1)
	UL.WriteULong(1)
	// Empty Sky light mask
	UL.WriteVarInt(1)
	UL.WriteULong(0b1111111111111111111111111111111111111111111111111111111111111110)
	// Empty Block light mask
	UL.WriteVarInt(1)
	UL.WriteULong(0b1111111111111111111111111111111111111111111111111111111111111110) //111111111111111111111111111111111111111111111111111111111111110
	//SkyLight array
	UL.WriteVarInt(1)
	UL.WriteVarInt(2048)
	for i := 0; i < 2048; i++ {
		UL.WriteUByte(0xFF)
	}
	UL.WriteVarInt(1)
	UL.WriteVarInt(2048)
	for i := 0; i < 2048; i++ {
		UL.WriteUByte(0xFF)
	}
	return UL.GetPacket()
}

func CreateChunkSection() []byte {
	CS := packet.CreateWriterWithCapacity(5000) // Chunk Section
	CS.WriteUShort(4096)                        // Number of non-air blocks
	CS.WriteUByte(8)                            // bits per block
	CS.WriteVarInt(6)                           // pallete length
	CS.WriteVarInt(0)                           //0
	CS.WriteVarInt(33)                          //1st
	CS.WriteVarInt(73)                          //2nd
	CS.WriteVarInt(33)                          //3rd
	CS.WriteVarInt(72)                          //4th
	CS.WriteVarInt(3953)                        //5th /3953 = redstone ore /3954 = deep slate redstone ore
	CS.WriteVarInt(512)                         //Number of longs
	AL := make([]byte, 4096)                    // Write 512 "longs", using bytes because bitsperblock is 8
	for i := 0; i < len(AL)-256; i++ {
		AL[i] = 4
	}
	for i := 3840; i < 4096; i++ {
		AL[i] = 5
	}
	CS.WriteArray(AL)
	return CS.GetData()
}

func CreateChunk(X, Z int32) []byte {
	Chunk := *new(packet.ChunkData_CB)
	Chunk.ChunkX = X
	Chunk.ChunkZ = Z
	Chunk.BitMaskLength = 1
	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 0b0000000000000000000000000000000000000000000000000000000000001111) // Only 4 sections for now
	HM := CreateStaticHeightMap()
	Chunk.HeightMaps = HM.Encode()
	Chunk.BiomeLength = 1024
	ChunkSection := CreateChunkSection()
	Chunk.Size = len(ChunkSection) + len(ChunkSection) + len(ChunkSection) + len(ChunkSection)
	C := packet.CreatePacketWriterWithCapacity(0x22, 34768) //construct chunk
	C.WriteInt(Chunk.ChunkX)
	C.WriteInt(Chunk.ChunkZ)
	C.WriteVarInt(Chunk.BitMaskLength)
	C.WriteLong(0b0000000000000000000000000000000000000000000000000000000000001111) //Array(Chunk.PrimaryBitMask)
	C.WriteArray(Chunk.HeightMaps)
	C.WriteVarInt(Chunk.BiomeLength)
	C.WriteArray(make([]byte, 1024))
	C.WriteVarInt(int32(Chunk.Size)) // Write chunk size
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteVarInt(0)                 // Write block entity length
	return C.GetPacket()
}

/* preperation for 1.18, more data needed
func CreateChunkSection() []byte {
	CS := packet.CreateWriterWithCapacity(5000) // Chunk Section
	CS.WriteUShort(4096)                        // Number of non-air blocks
	CS.WriteUByte(8)                            // bits per block
	CS.WriteVarInt(6)                           // pallete length
	CS.WriteVarInt(0)                           //0
	CS.WriteVarInt(33)                          //1st
	CS.WriteVarInt(73)                          //2nd
	CS.WriteVarInt(33)                          //3rd
	CS.WriteVarInt(72)                          //4th
	CS.WriteVarInt(3955)                        //5th /3953 = redstone ore /3954 = deep slate redstone ore
	CS.WriteVarInt(512)                         //Number of longs
	AL := make([]byte, 4096)                    // Write 512 "longs", using bytes because bitsperblock is 8
	for i := 0; i < len(AL)-256; i++ {
		AL[i] = 4
	}
	for i := 3840; i < 4096; i++ {
		AL[i] = 5
	}
	CS.WriteArray(AL)
	//Biome pallete container
	CS.WriteUByte(0)
	CS.WriteVarInt(0)
	// CS.WriteVarInt(256)
	CS.WriteLongArray(make([]int64, 256))
	return CS.GetData()
}

func CreateChunk(X, Z int32) []byte {
	Chunk := *new(packet.ChunkData_CB)
	Chunk.ChunkX = X
	Chunk.ChunkZ = Z
	Chunk.BitMaskLength = 1
	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 0b0000000000000000000000000000000000000000000000000000000000001111) // Only 4 sections for now
	HM := CreateStaticHeightMap()
	Chunk.HeightMaps = HM.Encode()
	Chunk.BiomeLength = 1024
	ChunkSection := CreateChunkSection()
	Chunk.Size = len(ChunkSection) + len(ChunkSection) + len(ChunkSection) + len(ChunkSection)
	C := packet.CreatePacketWriterWithCapacity(0x22, 34768) //construct chunk
	C.WriteInt(Chunk.ChunkX)
	C.WriteInt(Chunk.ChunkZ)
	// C.WriteVarInt(Chunk.BitMaskLength)
	// C.WriteLong(0b0000000000000000000000000000000000000000000000000000000000001111) //Array(Chunk.PrimaryBitMask)
	C.WriteArray(Chunk.HeightMaps)
	C.WriteVarInt(int32(Chunk.Size)) // Write chunk size
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteArray(ChunkSection)       // Write the chunk section
	C.WriteVarInt(0)                 // Write block entity length
	C.WriteBoolean(true)
	// Sky light mask
	C.WriteVarInt(1)
	C.WriteULong(1)
	// Block Light mask
	C.WriteVarInt(1)
	C.WriteULong(1)
	// Empty Sky light mask
	C.WriteVarInt(1)
	C.WriteULong(0b1111111111111111111111111111111111111111111111111111111111111110)
	// Empty Block light mask
	C.WriteVarInt(1)
	C.WriteULong(0b1111111111111111111111111111111111111111111111111111111111111110) //111111111111111111111111111111111111111111111111111111111111110
	//SkyLight array
	C.WriteVarInt(1)
	C.WriteVarInt(2048)
	for i := 0; i < 2048; i++ {
		C.WriteUByte(0xFF)
	}
	C.WriteVarInt(1)
	C.WriteVarInt(2048)
	for i := 0; i < 2048; i++ {
		C.WriteUByte(0xFF)
	}
	return C.GetPacket()
}
*/
