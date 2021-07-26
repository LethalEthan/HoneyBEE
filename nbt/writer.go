package nbt

import (
	"fmt"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyGO")
var TestData = []byte{10, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 8, 0, 4, 110, 97, 109, 101, 0, 9, 66, 97, 110, 97, 110, 114, 97, 109, 97, 0}

type NBTWriter struct {
	Name            string
	Data            []byte
	Result          []*TCompound //NBT tags usually onlt use root compound tag but this allows for it to have multiple
	CurrentTag      *TCompound
	size            int
	currentTagIndex int
	numElements     int
	numCompounds    int
}

func CreateNBTWriter(Name string) *NBTWriter {
	NBTW := new(NBTWriter)
	NBTW.Data = make([]byte, 0, 512)
	NBTW.Result = make([]*TCompound, 0, 2)
	Root := new(TCompound)
	Root.Name = Name
	Root.Value = make([]interface{}, 0, 16)
	NBTW.Result = append(NBTW.Result, Root)
	NBTW.CurrentTag = Root
	NBTW.numCompounds++
	return NBTW
}

//AddTag - Adds a tag to the current compound tag, requires a T{TagType} object e.g. TShort
//This can be dangerous to use unless you know what you're doing, try to keep to safe functions like AddTagBasedOnValue
func (NBTW *NBTWriter) AddTag(Val interface{}) {
	switch Val.(type) {
	case TEnd:
		NBTW.CurrentTag.AddTag(TEnd{})
	case TByte:
		NBTW.CurrentTag.AddTag(Val.(TByte))
	case TShort:
		NBTW.CurrentTag.AddTag(Val.(TShort))
	case TInt:
		NBTW.CurrentTag.AddTag(Val.(TInt))
	case TLong:
		NBTW.CurrentTag.AddTag(Val.(TLong))
	case TFloat:
		NBTW.CurrentTag.AddTag(Val.(TFloat))
	case TDouble:
		NBTW.CurrentTag.AddTag(Val.(TDouble))
	case TByteArray:
		BA := Val.(TByteArray)
		BA.length = len(BA.Value)
		NBTW.CurrentTag.AddTag(Val.(TByteArray))
	case TString:
		NBTW.CurrentTag.AddTag(Val.(TString))
	case TList:
		List := Val.(TList)
		List.length = len(List.Value)
		NBTW.CurrentTag.AddTag(List)
	case TCompound:
		TC := Val.(TCompound)
		TC.NumEntries = len(TC.Value)
		TC.PreviousTag = NBTW.CurrentTag
		TC.AddTag(TEnd{})
		NBTW.currentTagIndex++
		NBTW.CurrentTag = NBTW.CurrentTag.PreviousTag
		//NBTW.CurrentTag.Index++
		NBTW.CurrentTag.NumEntries = len(NBTW.CurrentTag.Value)
	case TIntArray:
		TIA := Val.(TIntArray)
		TIA.length = len(TIA.Value)
		NBTW.CurrentTag.AddTag(TIA)
	case TLongArray:
		TLA := Val.(TLongArray)
		TLA.length = len(TLA.Value)
		NBTW.CurrentTag.AddTag(TLA)
	}
}

//AddTagBasedOnValue - This creates and adds a tag that it based of the input given, give an int32 and it create TInt tag.
//This is limited to the primitives datatypes: byte, int16, int32, int64, float32, float64, string
//and map[string]interface{} that can hold a name to any primitive type (good for adding multiple values)
//Sucessful result will return no error and failure will return NBTUnknownType with the name and value that failed.
func (NBTW *NBTWriter) AddTagBasedOnValue(Name string, Val interface{}) error {
	if Name == "" {
		return NBTEmptyName
	}
	switch Val.(type) {
	case nil:
		return NBTnil
	case byte:
		NBTW.CurrentTag.AddTag(TByte{Name, Val.(byte)})
	case int16:
		NBTW.CurrentTag.AddTag(TShort{Name, Val.(int16)})
	case int32:
		NBTW.CurrentTag.AddTag(TInt{Name, Val.(int32)})
	case int64:
		NBTW.CurrentTag.AddTag(TLong{Name, Val.(int64)})
	case float32:
		NBTW.CurrentTag.AddTag(TFloat{Name, Val.(float32)})
	case float64:
		NBTW.CurrentTag.AddTag(TDouble{Name, Val.(float64)})
	case []byte:
		NBTW.CurrentTag.AddTag(TByteArray{Name, Val.([]byte), len(Val.([]byte)), 0})
	case string:
		NBTW.CurrentTag.AddTag(TString{Name, Val.(string)})
	case map[string]interface{}: //All types, iterated through, this will not accept compounds or lists since they cannot be deteced via primitives
		var err error
		for i, v := range Val.(map[string]interface{}) {
			Log.Info("key: ", i, " Value: ", v)
			e := NBTW.AddTagBasedOnValue(i, v)
			if e != nil {
				err = e
			}
		}
		return err
	case []int32:
		NBTW.CurrentTag.AddTag(TIntArray{Name, Val.([]int32), len(Val.([]int32)), 0})
	case []int64:
		NBTW.CurrentTag.AddTag(TLongArray{Name, Val.([]int64), len(Val.([]int64)), 0})
	default:
		return fmt.Errorf("NBTTypeUnknown: %s, Value: %v", Name, Val)
	}
	return nil
}

//Add multiple tags, this uses T{TagType} objects as tags require a name unless in a list
func (NBTW *NBTWriter) AddMultipleTags(Val []interface{}) {
	for _, v := range Val {
		NBTW.AddTag(v)
	}
}

//Add multiple tags, this uses primitive data types to create a tag
//byte, int16, int32, int64, float32, float64, string
//This function acts identical to using a map[string]interface in AddTagBasedOnValue
//Just that a map now doesn't have to be created
func (NBTW *NBTWriter) AddMultipleTagsBasedOnValue(Name []string, Val []interface{}) {
	for i, v := range Val {
		if Name[i] != "" {
			NBTW.AddTagBasedOnValue(Name[i], v)
		}
	}
}

func (NBTW *NBTWriter) AppendByteSlice(Data []byte) {
	NBTW.Data = append(NBTW.Data, Data...)
	NBTW.size += len(Data)
}

func (NBTW *NBTWriter) writeTag(Type byte, Name string) {
	NBTW.AppendByteSlice([]byte{Type})
	NBTW.writeString("", Name) //This works with the tests with 30% coverage I believe it's fine to do this
}

func (NBTW *NBTWriter) TestingShit() {
	NBTW.writeTag(TagString, "name")
	NBTW.writeString("", "Bananrama")
	NBTW.writeTag(TagEnd, "")
}
