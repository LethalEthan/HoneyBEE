package Packet

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
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
	res, err := pr.ReadByte()

	if err != nil {
		return false, err
	}

	if res != 0x00 && res != 0x01 {
		return false, fmt.Errorf("value %X not a boolean value", res)
	}

	return res != 0x00, nil
}

//ReadByte - reads a single byte from the packet and returns it, it returns a zero and an io.EOF if the packet has been already read to the end.
func (pr *PacketReader) ReadByte() (byte, error) {
	bytee, err := pr.ReadUnsignedByte()

	return byte(bytee), err
}

func (pr *PacketReader) ReadUnsignedByte() (byte, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	bytee := pr.data[pr.seek]

	_, err := pr.seekWithEOF(1, io.SeekCurrent)

	if err != nil {
		return bytee, err
	}

	return bytee, nil
}

func (pr *PacketReader) ReadShort() (int16, error) {
	short, err := pr.ReadUnsignedShort()
	return int16(short), err
}

func (pr *PacketReader) ReadUnsignedShort() (uint16, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	short := binary.BigEndian.Uint16(pr.data[pr.seek : pr.seek+2])

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

	longShort := int32(binary.BigEndian.Uint32(pr.data[pr.seek : pr.seek+4]))

	_, err := pr.seekWithEOF(4, io.SeekCurrent)

	if err != nil {
		return longShort, err
	}

	return longShort, nil
}

func (pr *PacketReader) ReadLong() (int64, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	long := int64(binary.BigEndian.Uint64(pr.data[pr.seek : pr.seek+8]))

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

	floatBits, err := pr.ReadInt()

	if err != nil {
		return 0, err
	}

	return math.Float32frombits(uint32(floatBits)), nil
}

func (pr *PacketReader) ReadDouble() (float64, error) {
	if pr.checkForEOF() {
		return 0, io.EOF
	}

	doubleBits, err := pr.ReadLong()

	if err != nil {
		return 0, err
	}

	return math.Float64frombits(uint64(doubleBits)), nil
}

func (pr *PacketReader) ReadString() (string, error) {
	if pr.checkForEOF() {
		return "", io.EOF
	}

	stringSize, err := pr.ReadVarInt()

	if err != nil {
		return "", err
	}

	if stringSize < 0 {
		return "", fmt.Errorf("string size of %d invalid", stringSize)
	}

	stringVal := string(pr.data[pr.seek : pr.seek+int64(stringSize)])

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
		bytee, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, err
		}
		val := int32((bytee & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 5 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}

		if bytee&0x80 == 0 {
			break
		}
	}
	/*TESTING:
	var buff = make([]byte, 5)
	for i := 0; i < 5; i++ {
		buff[i], err = pr.ReadUnsignedByte()
	}
	//tt, t := pr.ReadUnsignedByte()
	test, err := binary.Varint(buff)
	fmt.Print("TESTING: ", test, "ERR:", err)*/
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
		bytee, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, err
		}
		val := int64((bytee & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 10 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}

		if bytee&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func (pr *PacketReader) ReadArray() ([]byte, error) {
	return pr.data, nil
}
