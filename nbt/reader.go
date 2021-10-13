package nbt

import (
	"HoneyBEE/utils"
	"fmt"
	"math"
)

type NBTReader struct {
	data       []byte
	seeker     int
	end        int
	tagsAmount int
	result     []interface{}
}

func CreateNBTReader(Data []byte) (*NBTReader, error) {
	if len(Data) >= 3 {
		NBTR := new(NBTReader)
		NBTR.data = Data
		NBTR.seeker = 0
		NBTR.end = len(Data)
		NBTR.result = make([]interface{}, 0, 32)
		return NBTR, nil
	}
	return nil, NBTEmpty
}

func (NBTR *NBTReader) seek(offset int) {
	NBTR.seeker += offset
}

// func (NBTR *NBTReader) SeekWithEOF(offset int) error {
// 	if offset+NBTR.seeker > NBTR.end {
// 		return fmt.Errorf("Seek reached end: %d offset: %d", NBTR.seeker, offset) //errors.New("Seek reached End")
// 	} else {
// 		NBTR.seeker += offset
// 		return nil
// 	}
// }

func (NBTR *NBTReader) AutoRead() ([]interface{}, error) {
	for {
		var Type byte
		var Name string
		var err error
		Type, Name, err = NBTR.readTagType()
		if err != nil {
			return nil, err
		}
		switch Type {
		case TagEnd:
			NBTR.tagsAmount++
			return NBTR.result, nil
		case TagByte:
			Byte, err := NBTR.readByte()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TByte{Name, Byte})
		case TagShort:
			Short, err := NBTR.readShort()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TShort{Name, int16(Short)})
		case TagInt:
			Int, err := NBTR.readInt()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TInt{Name, int32(Int)})
		case TagLong:
			Long, err := NBTR.readLong()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TLong{Name, int64(Long)})
		case TagFloat:
			Float, err := NBTR.readFloat()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TFloat{Name, float32(Float)})
		case TagDouble:
			Double, err := NBTR.readDouble()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TDouble{Name, float64(Double)})
		case TagByteArray:
			ByteArray, err := NBTR.readByteArray()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, TByteArray{Name, ByteArray, len(ByteArray), 0})
		case TagString:
			String, err := NBTR.readString()
			if err != nil {
				return nil, err
			}
			NBTR.result = append(NBTR.result, String)
		case TagList:

		case TagCompound:
			NBTR.result = append(NBTR.result, Type)
		case TagIntArray:

		case TagLongArray:
		}
	}
}

func (NBTR *NBTReader) readTagType() (byte, string, error) {
	Type, err := NBTR.readByte()
	if err != nil {
		return 255, "", err
	}
	if Type == TagEnd {
		return 0, "", nil
	}
	TagName, err := NBTR.readString()
	if err != nil {
		return 255, "", err
	}
	return Type, TagName, nil
}

func (NBTR *NBTReader) readByte() (byte, error) {
	Byte, err := NBTR.readWithEOFSeek(1) //NBTR.data[NBTR.seeker]
	//err := NBTR.SeekWithEOF(1)
	if err != nil {
		return 0, err
	}
	return Byte[0], nil
}

func (NBTR *NBTReader) readShort() (int16, error) {
	ShortBytes, err := NBTR.readWithEOFSeek(2)
	//NBTR.SeekWithEOF(2)
	Short, err := utils.ByteArrayToInt16(ShortBytes)
	if err != nil {
		return 0, err
	}
	return Short, nil
}

func (NBTR *NBTReader) readInt() (int32, error) {
	IntBytes, err := NBTR.readWithEOFSeek(4)
	Int, err := utils.ByteArrayToInt32(IntBytes)
	if err != nil {
		return 0, err
	}
	return Int, nil
}

func (NBTR *NBTReader) readLong() (int64, error) {
	LongBytes, err := NBTR.readWithEOFSeek(8)
	Long, err := utils.ByteArrayToInt64(LongBytes)
	if err != nil {
		return 0, err
	}
	return Long, nil
}

func (NBTR *NBTReader) readFloat() (float32, error) {
	FloatBytes, err := NBTR.readInt()
	if err != nil {
		return 0, err
	}
	Float := math.Float32frombits(uint32(FloatBytes))
	return Float, nil
}

func (NBTR *NBTReader) readDouble() (float64, error) {
	DoubleBytes, err := NBTR.readLong()
	if err != nil {
		return 0, err
	}
	Double := math.Float64frombits(uint64(DoubleBytes))
	return Double, nil
}

func (NBTR *NBTReader) readByteArray() ([]byte, error) {
	Length, err := NBTR.readInt()
	if err != nil {
		panic(err)
	}
	BA, err := NBTR.readWithEOFSeek(int(Length))
	if err != nil {
		panic(err)
	}
	return BA, nil
}

func (NBTR *NBTReader) readString() (string, error) {
	StringLength, err := NBTR.readShort()
	if err != nil {
		return "", err
	}
	String, err := NBTR.readWithEOFSeek(int(StringLength)) /*NBTR.data[NBTR.seeker : NBTR.seeker+int(StringLength)]*/
	if err != nil {
		return "", err
	}
	return string(String), nil
}

func (NBTR *NBTReader) readWithEOFSeek(offset int) ([]byte, error) {
	//Check if offset overflows
	if NBTR.seeker+offset > NBTR.end {
		utils.PrintHexFromBytes("overflow", NBTR.data) //fmt.Print(NBTR.data)
		return []byte{0}, fmt.Errorf("offset: %d seeker: %d end: %d", offset, NBTR.seeker, NBTR.end)
	}
	//Return data
	Data := NBTR.data[NBTR.seeker : NBTR.seeker+offset]
	NBTR.seek(offset)
	return Data, nil
}

// func (NBTR *NBTReader) AppendByteSlice(Data []byte) {
// 	NBTR.data = append(NBTR.data, Data...)
// }
