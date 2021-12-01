package nbt

import (
	"encoding/binary"
	"fmt"
)

type NBTDecoder struct {
	data            []byte
	index           int
	Result          []interface{}
	currentCompound *Compound
}

func CreateNBTDecoder(Data []byte) NBTDecoder {
	return NBTDecoder{Data, 0, make([]interface{}, 0, 16), nil}
}

func (NBTD *NBTDecoder) seek(offset int) error {
	if NBTD.index+offset > len(NBTD.data) {
		return fmt.Errorf("Seek will over run buffer, seek: %d end: %d offset: %d", NBTD.index, len(NBTD.data), offset)
	}
	NBTD.index += offset
	return nil
}

func (NBTD *NBTDecoder) SetData(Data []byte) {
	NBTD.data = Data
	NBTD.index = 0
	NBTD.Result = NBTD.Result[:0]
	NBTD.currentCompound = nil
}

func (NBTD *NBTDecoder) Decode() {
	Log.Debug("Decode function")
	NBTD.Result = NBTD.Result[:0] //Reset length to 0 but keep capacity
	if err := NBTD.decode(); err != nil {
		panic(err)
	}
	Log.Debug(NBTD.Result...)
}

// decodes all tags into result
func (NBTD *NBTDecoder) decode() (err error) {
	var Type byte
	var Name string
	for {
		if Type, Name, err = NBTD.DecodeTag(); err != nil {
			return
		}
		switch Type {
		case TagEnd:
			NBTD.currentCompound.value = append(NBTD.currentCompound.value, End{})
			if NBTD.currentCompound == nil {
				return
			} else {
				NBTD.currentCompound = NBTD.currentCompound.previousTag
			}
		case TagByte:
			NBTD.currentCompound.value = append(NBTD.currentCompound.value, Byte{Name, NBTD.data[NBTD.index]})
			if err = NBTD.seek(1); err != nil {
				return
			}
		case TagShort:
			short := binary.BigEndian.Uint16(NBTD.data[NBTD.index : NBTD.index+2]) // read short
			NBTD.Result = append(NBTD.Result, Short{Name, int16(short)})           // create short tag
			if err = NBTD.seek(2); err != nil {
				return
			}
		case TagCompound:
			C := &Compound{Name, make([]interface{}, 0, 16), nil}
			if NBTD.currentCompound != nil {
				C.previousTag = NBTD.currentCompound
			}
			NBTD.currentCompound = C

			//	caseTagEnd:
			// 		NBTD.Result = append(NBTD.Result, End{})
			// 		return
			// 	case TagByte:
			// 		NBTD.Result = append(NBTD.Result, Byte{Name, NBTD.data[NBTD.index]}) // read byte
			// 		NBTD.seek(1)
			// 	case TagShort:
			// 		short := binary.BigEndian.Uint16(NBTD.data[NBTD.index : NBTD.index+2]) // read short
			// 		NBTD.Result = append(NBTD.Result, Short{Name, int16(short)})           // create short tag
			// 		if err = NBTD.seek(2); err != nil {
			// 			return
			// 		}
			// 	case TagInt:
			// 		integer := binary.BigEndian.Uint32(NBTD.data[NBTD.index : NBTD.index+4]) // read int
			// 		NBTD.Result = append(NBTD.Result, Int{Name, int32(integer)})             // create int tag
			// 		if err = NBTD.seek(4); err != nil {
			// 			return
			// 		}
			// 	case TagLong:
			// 		long := binary.BigEndian.Uint64(NBTD.data[NBTD.index : NBTD.index+8]) // read long
			// 		NBTD.Result = append(NBTD.Result, Long{Name, int64(long)})            // create int tag
			// 		if err = NBTD.seek(8); err != nil {
			// 			return
			// 		}
			// 	case TagFloat:
			// 		floatbits := binary.BigEndian.Uint32(NBTD.data[NBTD.index : NBTD.index+4]) // read int
			// 		float := math.Float32frombits(floatbits)                                   //utils.ByteArrayToInt64(NBTD.data[NBTD.index : NBTD.index+8]) // read int
			// 		NBTD.Result = append(NBTD.Result, Float{Name, float})                      // create int tag
			// 		if err = NBTD.seek(4); err != nil {
			// 			return
			// 		}
			// 	case TagDouble:
			// 		doublebits := binary.BigEndian.Uint64(NBTD.data[NBTD.index : NBTD.index+8]) // read long
			// 		double := math.Float64frombits(doublebits)                                  //utils.ByteArrayToInt64(NBTD.data[NBTD.index : NBTD.index+8]) // read int
			// 		NBTD.Result = append(NBTD.Result, Double{Name, double})                     // create int tag
			// 		if err = NBTD.seek(4); err != nil {
			// 			return
			// 		}
			// 	case TagByteArray:
			// 	case TagString:
			// 	case TagList:
			// 	case TagCompound:
			// 		NBTD.Result = append(NBTD.Result, Compound{Name, make([]interface{}, 0, 16, NBTD.currentCompound)})
			// 	case TagIntArray:
			// 	case TagLongArray:
			// 	}
			// }
		}
	}
}
