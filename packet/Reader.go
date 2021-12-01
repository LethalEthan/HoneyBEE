package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/google/uuid"
)

type PacketReader struct {
	data  []byte
	index int
}

//CreatePacketReader - Creates Packet Reader
func CreatePacketReader(data []byte) PacketReader {
	pr := *new(PacketReader)
	pr.data = data
	pr.index = 0
	return pr
}

//Seek - Seek through the data array
func (pr *PacketReader) seek(offset int) int {
	pr.index += offset
	return pr.index
}

func (pr *PacketReader) SetData(data []byte) {
	pr.index = 0
	pr.data = data
}

func (pr *PacketReader) GetSeeker() int {
	return pr.index
}

func (pr *PacketReader) GetEnd() int {
	return len(pr.data)
}

func (pr *PacketReader) SeekTo(pos int) bool {
	if pos > 0 {
		if pos > len(pr.data) {
			return false
		}
		pr.index = pos
		return true
	}
	return false
}

//CheckForEOF
func (pr *PacketReader) CheckForEOF() bool {
	return pr.index > len(pr.data)
}

func (pr *PacketReader) CheckEOFOffset(offset int) bool {
	if offset > 0 {
		if pr.index > len(pr.data) {
			return true
		}
		if pr.index+offset > len(pr.data) {
			return true
		}
		return false
	}
	return true
}

//whence is where the current Seek is and offset is how far the Seek should offset to
func (pr *PacketReader) Seek(offset int) (int, error) {
	if offset > 0 {
		if pr.index+offset > len(pr.data) {
			return 0, errors.New("seek reached end")
		}
		//Seek after EOF check
		pr.seek(offset)
		return offset, nil
	}
	return 0, errors.New("offset is negative")
}

//ReadBoolean - reads a single byte from the packet, and interprets it as a boolean.
//It throws an error and returns false if it has a problem either reading from the packet or encounters a value outside of the boolean range.
func (pr *PacketReader) ReadBoolean() (bool, error) {
	bool, err := pr.ReadUByte()
	return bool > 0, err
}

//ReadByte - reads a single byte from the packet and returns it, it returns a zero and an io.EOF if the packet has been already read to the end.
func (pr *PacketReader) ReadByte() (int8, error) {
	Byte, err := pr.ReadUByte()
	return int8(Byte), err
}

func (pr *PacketReader) ReadUByte() (byte, error) {
	if pr.CheckEOFOffset(1) {
		return 0, errors.New("EOF: UnsignedByte")
	}
	//Get byte from slice
	Byte := pr.data[pr.index]
	//Move the Seek
	_, err := pr.Seek(1)
	if err != nil {
		return Byte, err
	}
	return Byte, nil
}

func (pr *PacketReader) ReadShort() (int16, error) {
	short, err := pr.ReadUShort()
	return int16(short), err
}

func (pr *PacketReader) ReadUShort() (uint16, error) {
	if pr.CheckEOFOffset(2) {
		return 0, io.EOF
	}
	//Get the 2 bytes that make up the short
	short := binary.BigEndian.Uint16(pr.data[pr.index : pr.index+2])
	_, err := pr.Seek(2)
	if err != nil {
		return 0, err
	}
	return short, nil
}

func (pr *PacketReader) ReadInt() (int32, error) {
	I, err := pr.ReadUInt()
	return int32(I), err
}

func (pr *PacketReader) ReadUInt() (uint32, error) {
	if pr.CheckEOFOffset(4) {
		return 0, io.EOF
	}
	//Get the 4 bytes that make up the int
	UInteger := binary.BigEndian.Uint32(pr.data[pr.index : pr.index+4])
	//Move the Seek
	_, err := pr.Seek(4)
	if err != nil {
		return UInteger, err
	}
	return UInteger, nil
}

func (pr *PacketReader) ReadLong() (int64, error) {
	UL, err := pr.ReadULong()
	return int64(UL), err
}

func (pr *PacketReader) ReadULong() (uint64, error) {
	if pr.CheckEOFOffset(8) {
		return 0, io.EOF
	}
	//Get the 8 bytes that make up the long
	ULong := binary.BigEndian.Uint64(pr.data[pr.index : pr.index+8])
	//Move the Seek
	_, err := pr.Seek(8)
	if err != nil {
		return ULong, err
	}
	return ULong, nil
}

func (pr *PacketReader) ReadFloat() (float32, error) {
	if pr.CheckEOFOffset(4) {
		return 0, io.EOF
	}
	//Read the Int
	floatBits, err := pr.ReadInt()
	if err != nil {
		return 0, err
	}
	//Turn the int into float32
	return math.Float32frombits(uint32(floatBits)), nil
}

func (pr *PacketReader) ReadDouble() (float64, error) {
	if pr.CheckEOFOffset(8) {
		return 0, io.EOF
	}
	//Read the long
	doubleBits, err := pr.ReadLong()
	if err != nil {
		return 0, err
	}
	//Turn the long into float64
	return math.Float64frombits(uint64(doubleBits)), nil
}

func (pr *PacketReader) ReadString() (string, error) {
	if pr.CheckForEOF() {
		return "", errors.New("error on begin start string")
	}
	//Read string size
	StringSize, _, err := pr.ReadVarInt()
	if err != nil {
		return "", err
	}
	if pr.CheckForEOF() {
		return "", errors.New("error on second EOF check")
	}
	//StringSize check
	if StringSize < 0 {
		return "", errors.New("string size of %d invalid" + strconv.Itoa(int(StringSize)))
	}
	if pr.CheckEOFOffset(int(StringSize)) {
		return "", errors.New("StringSize exceeds EOF")
	}
	//Read the string
	StringVal := string(pr.data[pr.index : pr.index+int(StringSize)])
	//move the Seek
	_, err = pr.Seek(int(StringSize))
	if err != nil {
		return StringVal, err
	}
	return StringVal, nil
}

func (pr *PacketReader) ReadVarInt() (int32, byte, error) {
	if pr.CheckForEOF() {
		return 0, 0, errors.New("EOF: ReadVarInt")
	}
	var Result uint32
	var NumRead byte
	var Byte byte
	var val uint32
	var err error
	Byte, err = pr.ReadUByte()
	if err != nil {
		return 0, 0, err
	}
	for {
		val = uint32(Byte & 0x7F)
		Result |= (val << (7 * NumRead))
		//Increment
		NumRead++
		//Size check
		if NumRead > 5 {
			return 0, 0, fmt.Errorf("varint was over five bytes without termination")
		}
		//Termination
		if Byte&0x80 == 0 {
			break
		}
		if Byte, err = pr.ReadUByte(); err != nil {
			return 0, NumRead, err
		}
	}
	return int32(Result), NumRead, nil
}

func (pr *PacketReader) ReadVarLong() (int64, error) {
	if pr.CheckForEOF() {
		return 0, errors.New("EOF: ReadVarLong")
	}
	var Result uint64
	var NumRead byte
	var Byte byte
	var val uint64
	var err error
	Byte, err = pr.ReadUByte()
	if err != nil {
		return 0, err
	}
	for {
		val = uint64(Byte & 0x7F)
		Result |= (val << (7 * NumRead))
		//Increment
		NumRead++
		//Size check
		if NumRead > 10 {
			return 0, fmt.Errorf("varlong was over 10 bytes without termination")
		}
		//Termination
		if Byte&0x80 == 0 {
			break
		}
		if Byte, err = pr.ReadUByte(); err != nil {
			return 0, err
		}
	}
	return int64(Result), nil
}

func (pr *PacketReader) ReadUUID() (uuid.UUID, error) {
	if pr.CheckEOFOffset(16) {
		return uuid.Nil, io.EOF
	}
	UUIDBytes := pr.data[pr.index : pr.index+16]
	_, err := pr.Seek(16)
	if err != nil {
		return uuid.Nil, err
	}
	UUID, err := uuid.FromBytes(UUIDBytes)
	if err != nil {
		return uuid.Nil, err
	}
	return UUID, err
}

//ReadArray - Returns the array (slice) of the packet data
func (pr *PacketReader) ReadByteArray(length int) ([]byte, error) {
	//fmt.Print("Current: ", pr.index, "len: ", length)
	if pr.CheckEOFOffset(length) {
		return []byte{0}, io.EOF
	}
	data := pr.data[pr.index : pr.index+length]
	//fmt.Print("datalen: ", len(data))
	if _, err := pr.Seek(length); err != nil {
		return []byte{0}, io.EOF
	}
	//fmt.Println("index: ", pr.index)
	return data, nil
}

func (pr *PacketReader) ReadVarIntArray(length int) ([]int32, error) {
	var data = make([]int32, 0, length*4)
	for i := 0; i < length; i++ {
		l, _, err := pr.ReadVarInt()
		if err != nil {
			return []int32{0}, err
		}
		data = append(data, l)
	}
	// Log.Debug("Seeker: ", pr.index)
	return data, nil
}

func (pr *PacketReader) ReadLongArray(length int) ([]int64, error) {
	if pr.CheckEOFOffset(length * 8) {
		return []int64{0}, io.EOF
	}
	var data = make([]int64, 0, length)
	for i := 0; i < length; i++ {
		l, err := pr.ReadLong()
		if err != nil {
			return []int64{0}, err
		}
		data = append(data, l)
	}
	// Log.Debug("Seeker: ", pr.index)
	return data, nil
}

func (pr *PacketReader) ReadRestOfByteArrayNoSeek() []byte {
	return pr.data[pr.index:]
}

func (pr *PacketReader) ReadPosition() (X, Y, Z int64, err error) {
	POS, err := pr.ReadULong()
	if err != nil {
		return 0, 0, 0, err
	}
	X = int64(POS >> 38)
	Y = int64(POS & 0xFFF)
	Z = int64(POS << 26 >> 38)
	return
}

func (pr *PacketReader) ReadChunkSectionPosition() (X, Y, Z int64, err error) {
	CSPOS, err := pr.ReadULong()
	if err != nil {
		return 0, 0, 0, err
	}
	X = int64(CSPOS >> 42)
	Y = int64(CSPOS << 44 >> 44)
	Z = int64(CSPOS << 22 >> 42)
	return
}

// func (pr *PacketReader) ReadBlock() (BlockStateID int16, X,Y,Z int) {
// 	BPOS := pr.ReadULong()
// 	X = BPOS >> finshme
// 	Y = BPOS >> finishme
// 	Z = BPOS >> finishme
// }

func (pr *PacketReader) ReadIdentifier() (Identifier, error) {
	I, err := pr.ReadString()
	if err != nil {
		return "", err
	}
	return Identifier(I), nil
}

func (pr *PacketReader) ReadIdentifierArray(length int) ([]Identifier, error) {
	var err error
	IA := make([]Identifier, length)
	for i := 0; i < length; i++ {
		IA[i], err = pr.ReadIdentifier()
		if err != nil {
			return []Identifier{""}, err
		}
	}
	return IA, nil
}
