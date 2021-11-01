package packet

import (
	"encoding/binary"
	"math"

	"github.com/google/uuid"
)

type PacketWriter struct {
	data       []byte
	packetID   int32
	packetSize int
}

func CreatePacketWriter(PacketID int32) *PacketWriter {
	pw := new(PacketWriter)        //new packet with data struct Above
	pw.packetID = PacketID         //PacketID passed via function arguments
	pw.data = make([]byte, 0, 128) //Data is created with a byte array
	pw.WriteVarInt(PacketID)       //write PacketID to packet
	return pw
}

/*CreatePacketWriterWithCapacity - Create a packet writer with capacity on the data slice
max 2097151 if over it will default to a capacity of 128*/
func CreatePacketWriterWithCapacity(PacketID int32, Capacity int) *PacketWriter {
	pw := new(PacketWriter)
	pw.packetID = PacketID
	if Capacity > 0 && Capacity < 2097151 {
		pw.data = make([]byte, 0, Capacity)
		pw.WriteVarInt(PacketID)
		return pw
	}
	pw.data = make([]byte, 0, 128)
	pw.WriteVarInt(PacketID)
	return pw
}

func (pw *PacketWriter) GetData() []byte {
	return pw.data
}

func (pw *PacketWriter) ResetData(packetID int32) {
	pw.data = make([]byte, 0, len(pw.data))
	pw.packetSize = 0
	pw.packetID = packetID
}

func (pw *PacketWriter) GetPacket() []byte {
	pw.packetSize = len(pw.data)
	p := append(pw.CreateVarInt(uint32(pw.packetSize)), pw.data...)
	Log.Debug("PacketSize: ", len(pw.data))
	Log.Debug("Packet Contents: ", pw.data)
	return p
}

func (pw *PacketWriter) GetPacketID() int32 {
	return pw.packetID
}

func (pw *PacketWriter) GetPacketSize() int {
	return pw.packetSize
}

func (pw *PacketWriter) AppendByteSlice(Data []byte) {
	pw.data = append(pw.data, Data...)

	pw.packetSize += len(Data)
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
	pw.AppendByteSlice([]byte{val})
}

//WriteShort - Write Short to packet (int16)
func (pw *PacketWriter) WriteShort(val int16) {
	pw.WriteUnsignedShort(uint16(val))
}

//WriteUnsignedShort- Write Unsigned Short to packet (uint16)
func (pw *PacketWriter) WriteUnsignedShort(val uint16) {
	buff := make([]byte, 2)
	binary.BigEndian.PutUint16(buff, val)

	pw.AppendByteSlice(buff)
}

//WriteInt - Write Integer to packet (int32)
func (pw *PacketWriter) WriteInt(val int32) {
	pw.writeUnsignedInt(uint32(val))
}

//writeUnsignedInt - Write Unsigned Integer to packet (uint32)
func (pw *PacketWriter) writeUnsignedInt(val uint32) {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, val)

	pw.AppendByteSlice(buff)
}

//WriteLong - Write Long to packet (int64)
func (pw *PacketWriter) WriteLong(val int64) {
	pw.writeUnsignedLong(uint64(val))
}

//writeUnsignedLong - Write Unsigned Long (unint64)
func (pw *PacketWriter) writeUnsignedLong(val uint64) {
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, val)

	pw.AppendByteSlice(buff)
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
	pw.AppendByteSlice(val)
}

//WriteString - Write String to packet (string)
func (pw *PacketWriter) WriteString(val string) {
	pw.WriteVarInt(int32(len(val)))
	pw.AppendByteSlice([]byte(val))
}

func (pw *PacketWriter) WriteIdentifier(val Identifier) {
	pw.WriteVarInt(int32(len(val)))
	pw.AppendByteSlice([]byte(val))
}

func (pw *PacketWriter) WriteArrayIdentifier(val []Identifier) {
	for _, v := range val {
		pw.WriteIdentifier(v)
	}
}

//WriteVarInt - Write VarInt to packet (int32)
func (pw *PacketWriter) WriteVarInt(val int32) {
	pw.AppendByteSlice(pw.CreateVarInt(uint32(val)))
}

//WriteVarLong - Write VarLong (int64)
func (pw *PacketWriter) WriteVarLong(val int64) {
	pw.AppendByteSlice(pw.CreateVarLong(uint64(val)))
}

//CreateVarInt - creates VarInt, requires uint to move the sign bit
func (pw *PacketWriter) CreateVarInt(val uint32) []byte {
	var buff = make([]byte, 0, 5)
	var tmp byte
	for {
		tmp = byte(val & 0x7F)
		val = val >> 7
		if val != 0 {
			tmp |= 0x80
		}
		buff = append(buff, tmp)
		if val == 0 {
			break
		}
	}
	return buff
}

//CreateVarLong - Creates a VarLong, requires uint to move the sign bit
func (pw *PacketWriter) CreateVarLong(val uint64) []byte {
	var buff = make([]byte, 0, 10)
	for {
		temp := byte(val & 0x7F)
		val = val >> 7
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

func (pw *PacketWriter) WriteUUID(val uuid.UUID) {
	BU, err := val.MarshalBinary()
	if err != nil {
		Log.Debug("Could not marshal UUID!")
	}
	pw.AppendByteSlice(BU)
}
