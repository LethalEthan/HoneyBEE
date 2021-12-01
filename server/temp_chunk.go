package server

import (
	"HoneyBEE/nbt"
	"HoneyBEE/packet"
	"time"

	"github.com/panjf2000/gnet"
)

//
// A temporary mock up of chunk loading until I figure everything out properly
//

func ChunkLoad(c gnet.Conn) error {
	Chunk := *new(packet.ChunkData_CB)
	Chunk.ChunkX = 0
	Chunk.ChunkZ = 0
	Chunk.BitMaskLength = 1
	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 1) // Only 1 section
	NBTE := nbt.CreateNBTEncoder()
	BA := make([]int64, 37)
	for i := 0; i < len(BA); i++ { // Setup static heightmap
		if i < 36 { // last 3 bytes are unused
			BA[i] |= 0x0F << 55
			BA[i] |= 0x0F << 46
			BA[i] |= 0x0F << 37
			BA[i] |= 0x0F << 28
			BA[i] |= 0x0F << 19
			BA[i] |= 0x0F << 10
			BA[i] |= 0x0F << 1
		} else { // Whether it should be LSB or MSB is something I do not know
			BA[i] |= 0x0F << 28
			BA[i] |= 0x0F << 19
			BA[i] |= 0x0F << 10
			BA[i] |= 0x0F << 1
		}
	}
	Log.Debug("HeightMapL:", len(BA))
	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", BA))
	NBTE.EndCompoundTag()
	Chunk.HeightMaps = NBTE.Encode()
	Chunk.BiomeLength = 0
	CS := packet.CreateWriterWithCapacity(32800) // Chunk Section
	CS.WriteShort(4096)
	CS.WriteUByte(8)  //bits per block
	CS.WriteVarInt(1) //Pallette length
	CS.WriteVarInt(5) //BlockID for index 0
	// CS.WriteVarInt(0)
	// CS.WriteVarInt(5)
	CS.WriteVarInt(512)
	AL := make([]byte, 32768) //4
	for i := 0; i < len(AL); i++ {
		AL[i] = 0
	}
	CS.WriteArray(AL)
	CS.WriteVarInt(0) //End Of chunk section
	ChunkSection := CS.GetData()
	Chunk.Size = len(ChunkSection)
	C := packet.CreateWriterWithCapacity(34768) //construct chunk
	C.WriteVarInt(Chunk.BitMaskLength)
	C.WriteLongArray(Chunk.PrimaryBitMask)
	C.WriteArray(Chunk.HeightMaps)
	C.WriteVarInt(1024)
	C.WriteArray(make([]byte, 1024)) //(WriteBiomeID(0))
	C.WriteVarInt(int32(Chunk.Size))
	C.WriteArray(ChunkSection)
	C.WriteVarInt(0)
	//
	tmpCL := packet.CreateWriterWithCapacity(16) //temp chunk location writer
	tmpCL.WriteInt(0)
	tmpCL.WriteInt(0)
	//
	tmpUL := packet.CreateWriterWithCapacity(16) //temp chunk location writer
	tmpUL.WriteVarInt(0)
	tmpUL.WriteVarInt(0)
	//
	UL := packet.CreateWriterWithCapacity(3072)
	UL.WriteBoolean(true)
	// Sky light mask
	UL.WriteVarInt(1)
	UL.WriteLong(1)
	// Block Light mask
	UL.WriteVarInt(1)
	UL.WriteLong(1)
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
	//Send chunk
	PW := packet.CreatePacketWriterWithCapacity(0x25, 8192)
	PW.WriteArray(tmpUL.GetData())
	PW.WriteArray(UL.GetData())
	if err := c.AsyncWrite(PW.GetPacket()); err != nil {
		c.Close()
		return err
	}
	PW.ResetData(0x22)
	PW.WriteArray(tmpCL.GetData()) //Write X and Z
	PW.WriteArray(C.GetData())     // Add other chunk data
	if err := c.AsyncWrite(PW.GetPacket()); err != nil {
		c.Close() //send
		return err
	}
	Log.Debug("Packet Size: ", len(PW.GetData()))
	var j int32
	var i int32
	time.Sleep(5 * time.Second)
	for i = 0; i < 12; i++ {
		for j = 0; j < 12; j++ {
			if i != 0 || j != 0 {
				// Log.Debug("X: ", j, "Z: ", i)
				Chunk.ChunkX = j
				Chunk.ChunkZ = i
				tmpCL.ResetDataSlice()
				tmpCL.WriteInt(Chunk.ChunkX)
				tmpCL.WriteInt(Chunk.ChunkZ)
				tmpUL.ResetDataSlice()
				tmpUL.WriteVarInt(Chunk.ChunkX)
				tmpUL.WriteVarInt(Chunk.ChunkZ)
				PW.ResetData(0x25)
				PW.WriteArray(tmpUL.GetData())
				PW.WriteArray(UL.GetData())
				if err := c.AsyncWrite(PW.GetPacket()); err != nil {
					c.Close()
					return err
				}
				PW.ResetData(0x22)
				PW.WriteArray(tmpCL.GetData())
				PW.WriteArray(C.GetData())
				if err := c.AsyncWrite(PW.GetPacket()); err != nil {
					c.Close()
					return err
				}
				if i == 0 {
					//Log.Debug("X: ", -j, "Z: ", i)
					Chunk.ChunkX = -j
					Chunk.ChunkZ = i
					tmpCL.ResetDataSlice()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					tmpUL.ResetDataSlice()
					tmpUL.WriteVarInt(Chunk.ChunkX)
					tmpUL.WriteVarInt(Chunk.ChunkZ)
					PW.ResetData(0x25)
					PW.WriteArray(tmpUL.GetData())
					PW.WriteArray(UL.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
				}
				if j == 0 {
					// Log.Debug("X: ", j, "Z: ", -i)
					Chunk.ChunkX = j
					Chunk.ChunkZ = -i
					tmpCL.ResetDataSlice()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					tmpUL.ResetDataSlice()
					tmpUL.WriteVarInt(Chunk.ChunkX)
					tmpUL.WriteVarInt(Chunk.ChunkZ)
					PW.ResetData(0x25)
					PW.WriteArray(tmpUL.GetData())
					PW.WriteArray(UL.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
				}
				if j != 0 && i != 0 {
					//Log.Debug("X: ", j, "Z: ", -i)
					// Log.Debug("X: ", -j, "Z: ", i)
					// Log.Debug("X: ", -j, "Z: ", -i)
					Chunk.ChunkX = j
					Chunk.ChunkZ = -i
					tmpCL.ResetDataSlice()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					tmpUL.ResetDataSlice()
					tmpUL.WriteVarInt(Chunk.ChunkX)
					tmpUL.WriteVarInt(Chunk.ChunkZ)
					PW.ResetData(0x25)
					PW.WriteArray(tmpUL.GetData())
					PW.WriteArray(UL.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					//
					Chunk.ChunkX = -j
					Chunk.ChunkZ = i
					tmpCL.ResetDataSlice()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					tmpUL.ResetDataSlice()
					tmpUL.WriteVarInt(Chunk.ChunkX)
					tmpUL.WriteVarInt(Chunk.ChunkZ)
					PW.ResetData(0x25)
					PW.WriteArray(tmpUL.GetData())
					PW.WriteArray(UL.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					//
					Chunk.ChunkX = -j
					Chunk.ChunkZ = -i
					tmpCL.ResetDataSlice()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					tmpUL.ResetDataSlice()
					tmpUL.WriteVarInt(Chunk.ChunkX)
					tmpUL.WriteVarInt(Chunk.ChunkZ)
					PW.ResetData(0x25)
					PW.WriteArray(tmpUL.GetData())
					PW.WriteArray(UL.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := c.AsyncWrite(PW.GetPacket()); err != nil {
						c.Close()
						return err
					}
				}
			}
		}
	}
	Log.Debug("Sent chunks!")
	PW.ResetData(0x20) //Initialise World Border
	PW.WriteDouble(2000.0)
	PW.WriteDouble(2000.0)
	PW.WriteDouble(2001.0)
	PW.WriteDouble(2001.0)
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
