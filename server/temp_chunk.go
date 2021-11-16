package server

import (
	"HoneyBEE/nbt"
	"HoneyBEE/packet"
)

func (ClientConn *Client) ChunkLoad() {
	Chunk := *new(packet.ChunkData_CB)
	Chunk.ChunkX = 0
	Chunk.ChunkZ = 0
	Chunk.BitMaskLength = 1
	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 1) //only 1 section
	NBTE := nbt.CreateNBTEncoder()
	BA := make([]int64, 37)
	for i := 0; i < len(BA); i++ { //Setup static heightmap
		if i < 512 {
			BA[i] |= 0x0F << 55
			BA[i] |= 0x0F << 46
			BA[i] |= 0x0F << 37
			BA[i] |= 0x0F << 28
			BA[i] |= 0x0F << 19
			BA[i] |= 0x0F << 10
			BA[i] |= 0x0F << 1
		} else {
			BA[i] |= 0x0F << 28
			BA[i] |= 0x0F << 19
			BA[i] |= 0x0F << 10
			BA[i] |= 0x0F << 1
		}
	}
	Log.Debug("HeightMapL:", len(BA))
	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", BA)) //BasedOnValue("MOTION_BLOCKING", BA)
	NBTE.EndCompoundTag()
	NBTE.Encode()
	Log.Debug("NBTD: ", NBTE.GetData())
	Chunk.HeightMaps = NBTE.GetData()
	Chunk.BiomeLength = 0
	CS := packet.CreateWriterWithCapacity(32800) //Chunk Section
	CS.WriteShort(4096)
	CS.WriteUnsignedByte(8)
	CS.WriteVarInt(0)
	CS.WriteVarInt(512)
	AL := make([]byte, 32768)
	for i := 0; i < len(AL); i++ {
		AL[i] = 1
	}
	CS.WriteArray(AL)
	CS.WriteVarInt(0) //End Of chunk section
	ChunkSection := CS.GetData()
	Chunk.Size = len(ChunkSection)
	C := packet.CreateWriterWithCapacity(34768) //construct chunk
	C.WriteVarInt(Chunk.BitMaskLength)
	C.WriteLongArray(Chunk.PrimaryBitMask)
	C.WriteArray(Chunk.HeightMaps)
	C.WriteVarInt(int(Chunk.BiomeLength))
	C.WriteVarInt(Chunk.Size)
	C.WriteArray(ChunkSection)
	C.WriteVarInt(0)
	//
	tmpCL := packet.CreateWriterWithCapacity(16) //temp chunk location writer
	tmpCL.WriteInt(0)
	tmpCL.WriteInt(0)
	//
	PW := packet.CreatePacketWriterWithCapacity(0x22, 34768)
	PW.WriteArray(tmpCL.GetData()) //Write X and Z
	PW.WriteArray(C.GetData())     // Add other chunk data
	if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
		ClientConn.Conn.Close() //send
		return
	}
	Log.Debug("Packet Size: ", len(PW.GetData()))
	var j int
	var i int
	for i = 0; i < 12; i++ {
		for j = 0; j < 12; j++ {
			if i != 0 || j != 0 {
				// Log.Debug("X: ", j, "Z: ", i)
				Chunk.ChunkX = j
				Chunk.ChunkZ = i
				tmpCL.ClearData()
				tmpCL.WriteInt(Chunk.ChunkX)
				tmpCL.WriteInt(Chunk.ChunkZ)
				PW.ResetData(0x22)
				PW.WriteArray(tmpCL.GetData())
				PW.WriteArray(C.GetData())
				if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
					ClientConn.Conn.Close()
					return
				}
				if i == 0 {
					//Log.Debug("X: ", -j, "Z: ", i)
					Chunk.ChunkX = -j
					Chunk.ChunkZ = i
					tmpCL.ClearData()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
						ClientConn.Conn.Close()
						return
					}
				}
				if j == 0 {
					// Log.Debug("X: ", j, "Z: ", -i)
					Chunk.ChunkX = j
					Chunk.ChunkZ = -i
					tmpCL.ClearData()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
						ClientConn.Conn.Close()
						return
					}
				}
				if j != 0 && i != 0 {
					//Log.Debug("X: ", j, "Z: ", -i)
					// Log.Debug("X: ", -j, "Z: ", i)
					// Log.Debug("X: ", -j, "Z: ", -i)
					Chunk.ChunkX = j
					Chunk.ChunkZ = -i
					tmpCL.ClearData()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
						ClientConn.Conn.Close()
						return
					}
					//
					Chunk.ChunkX = -j
					Chunk.ChunkZ = i
					tmpCL.ClearData()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
						ClientConn.Conn.Close()
						return
					}
					//
					Chunk.ChunkX = -j
					Chunk.ChunkZ = -i
					tmpCL.ClearData()
					tmpCL.WriteInt(Chunk.ChunkX)
					tmpCL.WriteInt(Chunk.ChunkZ)
					PW.ResetData(0x22)
					PW.WriteArray(tmpCL.GetData())
					PW.WriteArray(C.GetData())
					if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
						ClientConn.Conn.Close()
						return
					}
				}
			}
		}
	}
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
	if err := ClientConn.Conn.AsyncWrite(PW.GetPacket()); err != nil {
		ClientConn.Conn.Close()
		return
	}
}

func TRYChunkLoad() {
	Chunk := new(packet.ChunkData_CB)
	Chunk.ChunkX = 0
	Chunk.ChunkZ = 0
	Chunk.BitMaskLength = 1
	Chunk.PrimaryBitMask = append(Chunk.PrimaryBitMask, 0b00000001)
	NBTE := nbt.CreateNBTEncoder()
	NBTE.AddTag(nbt.CreateLongArrayTag("MOTION_BLOCKING", make([]int64, 37)))
	NBTE.Encode()
	Log.Debug("X: ", 0, "Z: ", 0)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i != 0 || j != 0 {
				Log.Debug("X: ", j, "Z: ", i)
				if i == 0 {
					Log.Debug("X: ", -j, "Z: ", i)
				}
				if j == 0 {
					Log.Debug("X: ", j, "Z: ", -i)
				}
				if j != 0 && i != 0 {
					Log.Debug("X: ", j, "Z: ", -i)
					Log.Debug("X: ", -j, "Z: ", i)
					Log.Debug("X: ", -j, "Z: ", -i)
				}
			}
		}
	}
}

func (ClientConn *Client) GetClosed() bool {
	ClientConn.ClosedMutex.Lock()
	B := !ClientConn.Closed
	ClientConn.ClosedMutex.Unlock()
	return B
}

func (ClientConn *Client) SetClosed(v bool) {
	ClientConn.ClosedMutex.Lock()
	ClientConn.Closed = !v
	ClientConn.ClosedMutex.Unlock()
}
