package nbt

const (
	TagEnd       = iota //0  -00 Tag_End Only used in Tag_Compound
	TagByte             //1  -01 Tag_Byte byte
	TagShort            //2  -02 Tag_Short int16
	TagInt              //3  -03 Tag_Int int32
	TagLong             //4  -04 Tag_Long int64
	TagFloat            //5  -05 Tag_Float float32
	TagDouble           //6  -06 Tag_Double float64
	TagByteArray        //7  -07 Tag_ByteArray []byte
	TagString           //8  -08 Tag_String string
	TagList             //9  -09 Tag_List any tag with no names
	TagCompound         //10 -0a Tag_Compound string
	TagIntArray         //11 -0b Tag_Int_Array []int32
	TagLongArray        //12 -0c Tag_Long_Array []int64
)
