package packet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
)

var errEmpty = errors.New("array is empty, cannot decode")

//DecodeBoolean - reads a single byte from the packet, and interprets it as a boolean.
//It throws an error and returns false if it has a problem either reading from the packet or encounters a value outside of the boolean range.
func DecodeBoolean(D byte) (bool, error) {
	switch D {
	case 0x00:
		return false, nil
	case 0x01:
		return true, nil
	default:
		return false, errors.New("invalid value found in boolean!, likely incorrect byte")
	}
}

//DecodeByte - reads a single byte from the packet and returns it, it returns a zero and an io.EOF if the packet has been already read to the End.
func DecodeByte(D byte) byte {
	Byte := DecodeUnsignedByte(D)
	return Byte
}

func DecodeUnsignedByte(D byte) byte {
	return D
}

func DecodeShort(D []byte) (int16, error) {
	short, err := DecodeUShort(D)
	return int16(short), err
}

func DecodeUShort(D []byte) (uint16, error) {
	if len(D) > 2 {
		return 0, errors.New("data greater than 2 for short")
	}
	//Get the 2 bytes that make up the short
	Short := binary.BigEndian.Uint16(D[0:2])
	return Short, nil
}

func DecodeInteger(D []byte) (int32, error) {
	if len(D) > 4 {
		return 0, errors.New("data greater than 4 for VarInt")
	}
	//Get the 4 bytes that make up the int
	Integer := int32(binary.BigEndian.Uint32(D[0:4]))
	return Integer, nil
}

func DecodeLong(D []byte) (int64, error) {
	//Get the 8 bytes that make up the long
	if len(D) > 8 {
		return 0, errors.New("data greater than 8 for Long")
	}
	Long := int64(binary.BigEndian.Uint64(D[0:8]))
	return Long, nil
}

func DecodeFloat(D []byte) (float32, error) {
	//Decode the Int
	FloatBytes, err := DecodeInteger(D)
	if err != nil {
		return 0, err
	}
	//Turn the int into float32
	return math.Float32frombits(uint32(FloatBytes)), nil
}

func DecodeDouble(D []byte) (float64, error) {
	//Decode the long
	if len(D) == 0 {
		return 0, errEmpty
	}
	DoubleBytes, err := DecodeLong(D)
	if err != nil {
		return 0, err
	}
	//Turn the long into float64
	return math.Float64frombits(uint64(DoubleBytes)), nil
}

func DecodeString(D []byte, StringSize int) (string, error) {
	if StringSize <= 0 {
		return "", fmt.Errorf("string size of %d invalid" + strconv.Itoa(int(StringSize)))
	}
	stringVal := string(D[0:StringSize])
	return stringVal, nil
}

func DecodeVarInt(D []byte) (int32, uint32, error) {
	var Result int32
	var NumRead uint32
	if len(D) == 0 {
		return 0, 0, errEmpty
	}
	for {
		Byte := DecodeUnsignedByte(D[NumRead])
		//
		val := int32((Byte & 0x7F))
		Result |= (val << (7 * NumRead))
		//Increment
		NumRead++
		//Size check
		if NumRead > 5 || int(NumRead) > len(D) {
			return 0, 0, errors.New("varint was over five bytes without termination")
		}
		//Terminate
		if Byte&0x80 == 0 {
			break
		}
	}
	return Result, NumRead, nil
}

func DecodeVarLong(D []byte) (int64, error) {
	var Result int64
	var NumRead uint64
	if len(D) == 0 {
		return 0, errEmpty
	}
	for {
		Byte := DecodeUnsignedByte(D[NumRead])
		//
		val := int64((Byte & 0x7F))
		Result |= (val << (7 * NumRead))
		//Increment
		NumRead++
		//Size check
		if NumRead > 10 {
			return 0, fmt.Errorf("varlong was over 10 bytes without termination")
		}
		//Terminate
		if Byte&0x80 == 0 {
			break
		}
	}
	return Result, nil
}
