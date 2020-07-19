package VarTool //A tool to help decode and encode Variable Integer values used by minecrafts' protocol
//Right now this is a massive mess of spaget code
import (
	"encoding/binary"
	"fmt"
	"net"
)

type VarInt int32
type VarLong int64

//VarInt's are no longer than 5 bytes and are BigEndian
func DecodeVarInt(DM VarInt) (int32, error) {
	var result int32
	var numRead uint32
	for {
		value := int32((DM & 0x7F))
		result |= (value << (7 * numRead))
		numRead++
		if numRead > 5 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}
		if DM&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func EncodeVarInt(EM int32) (int32, error) {
	buf := make([]byte, 6)
	test := binary.PutVarint(buf, int64(EM))
	test2 := int32(test)
	return test2, nil
}

func DecodeVarLong(DM VarLong) (int64, error) {
	var result int64
	var numRead uint64
	for {
		val := int64((DM & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 10 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}

		if DM&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func ParseVarIntFromConnection(conn net.Conn) (int32, error) {
	var result int32
	var numRead uint32
	buff := make([]byte, 1)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			return 0, err
		}
		val := int32((buff[0] & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 5 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}

		if buff[0]&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func ParseVarLongFromConnection(conn net.Conn) (int64, error) {
	var result int64
	var numRead uint64
	buff := make([]byte, 1)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			return 0, err
		}
		val := int64((buff[0] & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 10 {
			return 0, fmt.Errorf("varlong was over ten bytes without termination")
		}

		if buff[0]&0x80 == 0 {
			break
		}
	}
	return result, nil
}

//ParseBoolFromConnection - Parses true/false data from connection
func ParseBoolFromConnection(conn net.Conn) (bool, error) {
	//var result bool
	var numRead bool
	numRead = false
	//error := 0
	return numRead, nil
}

// func EncodeVarIntFromArray(I []int32) (int32, error) {
// 	buff := make([]byte, 65) //len(I))
//
// 	for i := 0; i < 64; i++ {
//
// 		val := int32(I[i] & 0x7F) //int32((buff[0] & 0x7F))
// 		t := VarInt.Encode(VarInt(val))
// 		buff[i] = t[0:]
//
// 	}
// 	fmt.Print("Buffer: ", buff)
// 	buf := make([]byte, 10)
// 	fmt.Print("test: ", binary.PutVarint(buf, int64(1)))
// 	return 0, nil //buff, nil
// }

func (v VarInt) Encode() (vi []byte) {
	num := uint32(v)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		vi = append(vi, byte(b))
		if num == 0 {
			break
		}
	}
	return
}

// func (v *VarInt) Decode(r DecodeReader) error {
// 	var n uint32
// 	for i := 0; ; i++ {
// 		sec, err := r.ReadByte()
// 		if err != nil {
// 			return err
// 		}
//
// 		n |= uint32(sec&0x7F) << uint32(7*i)
//
// 		if sec&0x80 == 0 {
// 			break
// 		} else if i > 5 {
// 			return errors.New("VarInt is too big")
// 		}
// 	}
//
// 	*v = VarInt(n)
// 	return nil
// }
