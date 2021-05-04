package npacket

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
)

type PacketReader struct {
	Data   []byte
	Seeker int64
	End    int64
	Status int
}

var EOF = errors.New("Packet EOF")

//CreatePacketReader - Creates Packet Reader
func CreatePacketReader(Data []byte) *PacketReader {
	pr := new(PacketReader)
	pr.Data = Data
	pr.Seeker = 0
	pr.End = int64(len(Data))
	pr.Status = Begin
	return pr
}

//Seek - Seek through the Data array
func (pr *PacketReader) Seek(offset int64, stat int) (int64, error) {
	switch stat {
	case Begin:
		{
			if offset < 0 {
				return pr.Seeker, fmt.Errorf("Seek of %d is below zero", offset)
			}
			pr.Status = Current
			if offset >= pr.End {
				pr.Seeker = pr.End
				pr.Status = End
			} else {
				pr.Seeker = offset
				pr.Status = Current
			}
			return pr.Seeker, nil
		}
	case Current:
		{
			if pr.Seeker+offset < 0 {
				return pr.Seeker, fmt.Errorf("Seek adjustment of %d from beginning seeks below zero", offset)
			}
			if pr.Seeker+offset >= pr.End {
				pr.Seeker = pr.End
				pr.Status = End
			} else {
				pr.Seeker += offset
			}
			return pr.Seeker, nil
		}
	case End:
		{
			if pr.End+offset > 0 {
				return pr.Seeker, fmt.Errorf("Seek adjustment of %d from End seeks above zero", offset)
			}
			//Hardly used but is meant to seek backwards from the end
		}
	}
	return 0, fmt.Errorf("An invalid whence value was submitted")
}

func (pr *PacketReader) CheckForEOF() bool {
	return pr.Seeker >= pr.End
}

func (pr *PacketReader) CheckForEOFWithSeek(SeekTo int64) bool {
	if pr.Seeker > pr.End {
		return true
	}
	if pr.Seeker+SeekTo > pr.End {
		return true
	}
	return false
}

///
/// Reminder: stat argument isn't needed it's stored in pt.State! remove it
///

//whence is where the current Seek is and offset is how far the Seek should offset to
func (pr *PacketReader) SeekWithEOF(offset int64, stat int) (int64, error) {
	if offset+pr.Seeker > pr.End {
		return offset, errors.New("Seek reached End")
	}
	//Seek after EOF check
	offset, err := pr.Seek(offset, stat)
	if err != nil {
		return offset, err
	}
	return offset, nil
}

//func (pr *PacketReader) Read(p []byte) (int, error) {
// if pr.CheckForEOF() {
// 	return 0, io.EOF
// }
//
// num := copy(p, pr.Data[pr.Seeker:])
//
// _, err := pr.SeekWithEOF(int64(num), io.SeekCurrent)
//
// if err != nil {
// 	return num, err
// }
//
// return num, nil
//}

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
		return false, errors.New("Invalid value found in boolean, likely incorrect seek")
	}
}

//ReadByte - reads a single byte from the packet and returns it, it returns a zero and an io.EOF if the packet has been already read to the End.
func (pr *PacketReader) ReadByte() (byte, error) {
	Byte, err := pr.ReadUnsignedByte()
	return Byte, err
}

func (pr *PacketReader) ReadUnsignedByte() (byte, error) {
	if pr.CheckForEOFWithSeek(1) {
		return 0, errors.New("EOF: UnsignedByte")
	}
	//Get byte from slice
	Byte := pr.Data[pr.Seeker]
	//Move the Seek
	_, err := pr.SeekWithEOF(1, Current)
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
	if pr.CheckForEOFWithSeek(2) {
		return 0, io.EOF
	}
	//Get the 2 bytes that make up the short
	_, err := pr.SeekWithEOF(2, Current)
	if err != nil {
		return 0, err
	}
	short := binary.BigEndian.Uint16(pr.Data[pr.Seeker : pr.Seeker+2])
	//Move the Seek

	return short, nil
}

func (pr *PacketReader) ReadInt() (int32, error) {
	if pr.CheckForEOFWithSeek(4) {
		return 0, EOF
	}
	//Get the 4 bytes that make up the int
	Integer := int32(binary.BigEndian.Uint32(pr.Data[pr.Seeker : pr.Seeker+4]))
	//Move the Seek
	_, err := pr.SeekWithEOF(4, io.SeekCurrent)
	if err != nil {
		return Integer, err
	}
	return Integer, nil
}

func (pr *PacketReader) ReadLong() (int64, error) {
	if pr.CheckForEOFWithSeek(8) {
		return 0, io.EOF
	}
	//Get the 8 bytes that make up the long
	long := int64(binary.BigEndian.Uint64(pr.Data[pr.Seeker : pr.Seeker+8]))
	//Move the Seek
	_, err := pr.SeekWithEOF(8, io.SeekCurrent)
	if err != nil {
		return long, err
	}
	return long, nil
}

func (pr *PacketReader) ReadFloat() (float32, error) {
	if pr.CheckForEOFWithSeek(4) {
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
	if pr.CheckForEOFWithSeek(8) {
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
		return "", errors.New("Error on begin start string")
	}
	//Read string size
	StringSize, _, err := pr.ReadVarInt()
	if err != nil {
		return "", err
	}
	if pr.CheckForEOF() {
		return "", EOF
	}
	//StringSize check
	if StringSize < 0 {
		return "", errors.New("string size of %d invalid" + strconv.Itoa(int(StringSize)))
	}
	if int64(StringSize) > pr.End {
		return "", errors.New("StringSize exceeds EOF")
	}
	if int64(StringSize)+pr.Seeker > pr.End {
		return "", errors.New("String size + seeker = EOF")
	}
	//Read the string
	StringVal := string(pr.Data[pr.Seeker : pr.Seeker+int64(StringSize)])
	//move the Seek
	_, err = pr.SeekWithEOF(int64(StringSize), Current)
	if err != nil {
		return StringVal, err
	}
	return StringVal, nil
}

func (pr *PacketReader) ReadVarInt() (int32, uint32, error) {
	if pr.CheckForEOF() {
		return 0, 0, errors.New("EOF: ReadVarInt")
	}
	var Result int32
	var NumRead uint32
	for {
		Byte, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, 0, err
		}
		val := int32((Byte & 0x7F))
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
	}
	return Result, NumRead, nil
}

func (pr *PacketReader) ReadVarLong() (int64, error) {
	if pr.CheckForEOF() {
		return 0, io.EOF
	}
	var Result int64
	var NumRead uint64
	for {
		Byte, err := pr.ReadUnsignedByte()
		if err != nil {
			return 0, err
		}
		val := int64((Byte & 0x7F))
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
	}
	_, err := pr.SeekWithEOF(int64(NumRead), io.SeekCurrent)
	if err != nil {
		return Result, err
	}
	return Result, nil
}

//ReadArray - Returns the array (slice) of the packet Data
func (pr *PacketReader) ReadArray() ([]byte, error) {
	return pr.Data, nil
}

const (
	Begin   = 0
	Current = 1
	End     = 2
)
