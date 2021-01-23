package Packet

import (
	"encoding/binary"
	"math"
)

type PacketWriter struct {
	Data       []byte
	packetID   int32
	packetSize int32
}

func CreatePacketWriter(packetID int32) *PacketWriter {
	pw := new(PacketWriter)   //new packet with data struct Above
	pw.packetID = packetID    //packetID passed via function arguments
	pw.Data = make([]byte, 0) //Data is created with a byte array
	pw.WriteVarInt(packetID)  //write PacketID to packet
	return pw
}

func (pw *PacketWriter) GetPacket() []byte {
	return append(pw.CreateVarLong(int64(pw.packetSize)), pw.Data...)
}

func (pw *PacketWriter) appendByteSlice(Data []byte) {
	pw.Data = append(pw.Data, Data...)

	pw.packetSize += int32(len(Data))
}

//WriteBoolean - Write Boolean to packet
func (pw *PacketWriter) WriteBoolean(val bool) {
	if val {
		pw.WriteUnsignedByte(0x01) //true
	} else {
		pw.WriteUnsignedByte(0x00) //false
	}
}

//WriteByte - Write Byte to packet (int8)
func (pw *PacketWriter) WriteByte(val int8) {
	pw.WriteUnsignedByte(byte(val))
}

//WriteUnsignedByte - Write Unsigned Byte to packet (uint8)
func (pw *PacketWriter) WriteUnsignedByte(val byte) {
	pw.appendByteSlice([]byte{val})
}

//WriteShort - Write Short to packet (int16)
func (pw *PacketWriter) WriteShort(val int16) {
	pw.WriteUnsignedShort(uint16(val))
}

//WriteUnsignedShort- Write Unsigned Short to packet (uint16)
func (pw *PacketWriter) WriteUnsignedShort(val uint16) {
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, val)

	pw.appendByteSlice(buff)
}

//WriteInt - Write Integer to packet (int32)
func (pw *PacketWriter) WriteInt(val int32) {
	pw.writeUnsignedInt(uint32(val))
}

//writeUnsignedInt - Write Unsigned Integer to packet (uint32)
func (pw *PacketWriter) writeUnsignedInt(val uint32) {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, val)

	pw.appendByteSlice(buff)
}

//WriteLong - Write Long to packet (int64)
func (pw *PacketWriter) WriteLong(val int64) {
	pw.writeUnsignedLong(uint64(val))
}

//writeUnsignedLong - Write Unsigned Long (unint64)
func (pw *PacketWriter) writeUnsignedLong(val uint64) {
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, val)

	pw.appendByteSlice(buff)
}

//WriteFloat - Write Float to packet (float32)
func (pw *PacketWriter) WriteFloat(val float32) {
	pw.writeUnsignedInt(math.Float32bits(val))
}

//WriteDouble - Write Double to packet (float64)
func (pw *PacketWriter) WriteDouble(val float64) {
	pw.writeUnsignedLong(math.Float64bits(val))
}

//WriteArray - Write an array of bytes ([]byte)
func (pw *PacketWriter) WriteArray(val []byte) {
	pw.appendByteSlice(val)
}

//WriteString - Write String to packet (string)
func (pw *PacketWriter) WriteString(val string) {
	pw.WriteVarInt(int32(len(val)))
	pw.appendByteSlice([]byte(val))
}

//WriteVarInt - Write VarInt to packet (int32)
func (pw *PacketWriter) WriteVarInt(val int32) {
	pw.WriteVarLong(int64(val))
}

//WriteVarLong - Write VarLong (int64)
func (pw *PacketWriter) WriteVarLong(val int64) {
	//tt := pw.CreateVarLong(val)
	//Log.Debug("!!!!: ", tt)
	pw.appendByteSlice(pw.CreateVarLong(val))
}

//CreateVarLong - Creates a VarLong
func (pw *PacketWriter) CreateVarLong(val int64) []byte {
	var buff []byte
	for {
		temp := byte(val & 0x7F)
		val = int64(val >> 7)
		if val != 0 {
			temp |= 0x80
		}
		buff = append(buff, temp)
		if val == 0 {
			break
		}
	}
	return buff
}

//
// func (pw *PacketWriter) UseInterface(i interface{}) {
// 	switch i.(type) {
// 	case:
// 	}
// }
