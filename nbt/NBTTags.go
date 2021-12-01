package nbt

import "HoneyBEE/utils"

const (
	TagEnd       byte = iota //0  -00 Tag_End Only used in Tag_Compound
	TagByte                  //1  -01 Tag_Byte byte
	TagShort                 //2  -02 Tag_Short int16
	TagInt                   //3  -03 Tag_Int int32
	TagLong                  //4  -04 Tag_Long int64
	TagFloat                 //5  -05 Tag_Float float32
	TagDouble                //6  -06 Tag_Double float64
	TagByteArray             //7  -07 Tag_ByteArray []byte
	TagString                //8  -08 Tag_String string
	TagList                  //9  -09 Tag_List any tag with no names
	TagCompound              //10 -0a Tag_Compound string
	TagIntArray              //11 -0b Tag_Int_Array []int32
	TagLongArray             //12 -0c Tag_Long_Array []int64
)

//EncodeTag - Encodes the name and tag into the encoder buf only if there is a name, a lack of names means it's in a list
func (NBTE *NBTEncoder) EncodeTag(Type byte, Name string) {
	if Name != "" {
		NBTE.data = append(NBTE.data, Type)
		NBTE.data = append(NBTE.data, utils.Int16ToByteArray(int16(len(Name)))...)
		NBTE.data = append(NBTE.data, []byte(Name)...)
	} else {
		NBTE.data = append(NBTE.data, Type)
		NBTE.data = append(NBTE.data, []byte{0, 0}...)
	}
}

//EncodeTag - Encodes the name and tag into the encoder buf only if there is a name, a lack of names means it's in a list
func (NBTD *NBTDecoder) DecodeTag() (Type byte, Name string, err error) {
	Type = NBTD.data[NBTD.index] // Read tag type
	NBTD.seek(1)
	if Type == TagEnd {
		return
	}
	NL, err := utils.ByteArrayToInt16(NBTD.data[NBTD.index : NBTD.index+2]) // Read name length
	if err != nil {
		panic(err)
	}
	if err = NBTD.seek(2); err != nil {
		return
	}
	if NL > 0 {
		Name = string(NBTD.data[NBTD.index:NL])   //Read name
		if err = NBTD.seek(int(NL)); err != nil { // Seek name length
			return
		}
	}
	return
}
