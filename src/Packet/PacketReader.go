package Packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
)

type PacketReader struct {
	data []byte
	seek int64
	end  int64
}

//CreatePacketReader - Creates Packet Reader
func CreatePacketReader(data []byte) *PacketReader {
	pr := new(PacketReader)
	pr.data = data
	pr.seek = 0
	pr.end = int64(len(data))
	return pr
}

func (pr *PacketReader) Seek(offset int64, whence int) (int64, error) {

	switch whence {
	case io.SeekStart:
		{
			if offset < 0 {
				return pr.seek, fmt.Errorf("seek of %d is below zero", offset)
			}
			if offset > pr.end {
				pr.seek = pr.end
			} else {
				pr.seek = offset
			}
			return pr.seek, nil
		}
	case io.SeekCurrent:
		{
			if pr.seek+offset < 0 {
				return pr.seek, fmt.Errorf("seek adjustment of %d from beginning seeks below zero", offset)
			}
			if pr.seek+offset > pr.end {
				pr.seek = pr.end
			} else {
				pr.seek += offset
			}
			return pr.seek, nil
		}
	case io.SeekEnd:
		{
			if pr.end+offset < 0 {
				return pr.seek, fmt.Errorf("seek adjustment of %d from end seeks below zero", offset)
			}
			if pr.end+offset > pr.end {
				pr.seek = pr.end
			} else {
				pr.seek = pr.end + offset
			}
			return pr.seek, nil
		}
	}
	return 0, fmt.Errorf("An invalid whence value was submitted")
}

func (pr *PacketReader) checkForEOF() bool {
	return pr.seek >= pr.end
}

//whence is where the current seek is and offset is how far the seek should offset to
func (pr *PacketReader) seekWithEOF(offset int64, whence int) (int64, error) {
	offset, err := pr.Seek(offset, whence)
	if err != nil {
		return offset, err
	}
	if offset > pr.end {
		return offset, io.EOF
	}
	return offset, nil
}

func (pr *PacketReader) Read(p []byte) (int, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	num := copy(p, pr.data[pr.seek:])

	_, err := pr.seekWithEOF(int64(num), io.SeekCurrent)

	if err != nil {
		return num, err
	}

	return num, nil
}

//ReadBoolean - reads a single byte from the packet, and interprets it as a boolean.
//It throws an error and returns false if it has a problem either reading from the packet or encounters a value outside of the boolean range.
func (pr *PacketReader) ReadBoolean() (bool, error) {
	bool, err := pr.ReadByte()
	if err != nil {
		return false, err
	}
	switch bool {
	case 0x00:
		return false, nil
	case 0x01:
		return true, nil
	default:
		return false, errors.New("Invalid value found in boolean")
	}
	//
	// if bool != 0x00 && bool != 0x01 {
	// 	return false, fmt.Errorf("value %X not a boolean value", bool)
	// }
	//
	// return bool != 0x00, nil
}

//ReadByte - reads a single byte from the packet and returns it, it returns a zero and an io.EOF if the packet has been already read to the end.
func (pr *PacketReader) ReadByte() (byte, error) {
	Byte, err := pr.ReadUnsignedByte()

	return Byte, err
}

func (pr *PacketReader) ReadUnsignedByte() (byte, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}
	//Get byte from slice
	Byte := pr.data[pr.seek]
	//Move the seek
	_, err := pr.seekWithEOF(1, io.SeekCurrent)
	if err != nil {
		return Byte, err
	}
	return Byte, nil
}

func (pr *PacketReader) ReadShort() (int16, error) {
	short, err := pr.ReadUnsignedShort()
	return int16(short), err
}

func (pr *PacketReader) ReadUnsignedShort() (uint16, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}
	//Get the 2 bytes that make up the short
	short := binary.BigEndian.Uint16(pr.data[pr.seek : pr.seek+2])
	//Move the seek
	_, err := pr.seekWithEOF(2, io.SeekCurrent)
	if err != nil {
		return short, err
	}
	return short, nil
}

func (pr *PacketReader) ReadInt() (int32, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}
	//Get the 4 bytes that make up the int
	Integer := int32(binary.BigEndian.Uint32(pr.data[pr.seek : pr.seek+4]))
	//Move the seek
	_, err := pr.seekWithEOF(4, io.SeekCurrent)
	if err != nil {
		return Integer, err
	}
	return Integer, nil
}

func (pr *PacketReader) ReadLong() (int64, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}
	//Get the 8 bytes that make up the long
	long := int64(binary.BigEndian.Uint64(pr.data[pr.seek : pr.seek+8]))
	//Move the seek
	_, err := pr.seekWithEOF(8, io.SeekCurrent)
	if err != nil {
		return long, err
	}
	return long, nil
}

func (pr *PacketReader) ReadFloat() (float32, error) {
	if pr.checkForEOF() {
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
	if pr.checkForEOF() {
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
	if pr.checkForEOF() {
		return "", io.EOF
	}
	//Read string size
	stringSize, err := pr.ReadVarInt()
	if err != nil {
		return "", err
	}
	//StringSize check
	if stringSize < 0 {
		return "", errors.New("string size of %d invalid" + strconv.Itoa(int(stringSize)))
	}
	if int64(stringSize) >= pr.end || int64(stringSize)+pr.seek > pr.end {
		return "", io.EOF
	}
	//Read the string
	stringVal := string(pr.data[pr.seek : pr.seek+int64(stringSize)])
	//move the seek
	_, err = pr.seekWithEOF(int64(stringSize), io.SeekCurrent)
	if err != nil {
		return stringVal, err
	}
	return stringVal, nil
}

func (pr *PacketReader) ReadVarInt() (int32, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	var result int32
	var numRead uint32
	for {
		Byte, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, err
		}
		val := int32((Byte & 0x7F))
		result |= (val << (7 * numRead))
		//Increment
		numRead++
		//Size check
		if numRead > 5 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}
		//if Byte and 128 == 0
		if Byte&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func (pr *PacketReader) ReadVarInt2() (int32, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}
	return 0, io.EOF
}

func (pr *PacketReader) ReadVarLong() (int64, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	var result int64
	var numRead uint64
	for {
		Byte, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, err
		}
		val := int64((Byte & 0x7F))
		result |= (val << (7 * numRead))
		//Increment
		numRead++
		//Size check
		if numRead > 10 {
			return 0, fmt.Errorf("varlong was over 10 bytes without termination")
		}
		//
		if Byte&0x80 == 0 {
			break
		}
	}
	return result, nil
}

//ReadArray - Returns the array (slice) of the packet data
func (pr *PacketReader) ReadArray() ([]byte, error) {
	return pr.data, nil
}
