package nbt

import (
	"fmt"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyBEE")
var TestData = []byte{10, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 8, 0, 4, 110, 97, 109, 101, 0, 9, 66, 97, 110, 97, 110, 114, 97, 109, 97, 0}

type NBTWriter struct {
	Name            string
	Data            []byte
	Result          []*TCompound //NBT tags usually only use root compound tag but this allows for it to have multiple
	CurrentTag      *TCompound
	size            int
	currentTagIndex int
	numElements     int
	numCompounds    int
	totalNumEntries int
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

/*AddTag - Adds a tag to the current compound tag, requires a T{TagType} object e.g. TShort
This can be dangerous to use unless you know what you're doing, try to keep to safe functions like AddTagBasedOnValue
Compound tags are automatically ended when added do not end the tag youself, the encoder will ignore it anyway*/
func (NBTW *NBTWriter) AddTag(Val interface{}) error {
	switch Val := Val.(type) {
	case TEnd:
		NBTW.CurrentTag.AddTag(TEnd{})
	case TByte:
		NBTW.CurrentTag.AddTag(Val)
	case TShort:
		NBTW.CurrentTag.AddTag(Val)
	case TInt:
		NBTW.CurrentTag.AddTag(Val)
	case TLong:
		NBTW.CurrentTag.AddTag(Val)
	case TFloat:
		NBTW.CurrentTag.AddTag(Val)
	case TDouble:
		NBTW.CurrentTag.AddTag(Val)
	case TByteArray:
		BA := Val
		BA.length = len(BA.Value)
		NBTW.CurrentTag.AddTag(Val)
	case TString:
		NBTW.CurrentTag.AddTag(Val)
	case TList:
		List := Val
		List.length = len(List.Value)
		NBTW.CurrentTag.AddTag(List)
	case TCompound:
		TC := Val
		TC.NumEntries = len(TC.Value)
		TC.PreviousTag = NBTW.CurrentTag
		TC.AddTag(TEnd{})
		//NBTW.currentTagIndex++
		//NBTW.CurrentTag = NBTW.CurrentTag.PreviousTag
		NBTW.totalNumEntries += len(TC.Value)
		//NBTW.CurrentTag.Index++
		NBTW.CurrentTag.NumEntries++ //= len(TC.Value)
	case TIntArray:
		TIA := Val
		TIA.length = len(TIA.Value)
		NBTW.CurrentTag.AddTag(TIA)
	case TLongArray:
		TLA := Val
		TLA.length = len(TLA.Value)
		NBTW.CurrentTag.AddTag(TLA)
	default:
		return NBTUnknownType
	}
	return nil
}

//AddTagBasedOnValue - This creates and adds a tag that it based of the input given, give an int32 and it create TInt tag.
//This is limited to the primitives datatypes: byte, int16, int32, int64, float32, float64, string
//and map[string]interface{} that can hold a name to any primitive type (good for adding multiple values)
//Successful result will return no error and failure will return NBTUnknownType with the name and value that failed.
func (NBTW *NBTWriter) AddTagBasedOnValue(Name string, Val interface{}) error {
	if Name == "" {
		return NBTEmptyName
	}
	switch Val := Val.(type) {
	case nil:
		return NBTnil
	case byte:
		NBTW.CurrentTag.AddTag(TByte{Name, Val})
	case int16:
		NBTW.CurrentTag.AddTag(TShort{Name, Val})
	case int32:
		NBTW.CurrentTag.AddTag(TInt{Name, Val})
	case int64:
		NBTW.CurrentTag.AddTag(TLong{Name, Val})
	case float32:
		NBTW.CurrentTag.AddTag(TFloat{Name, Val})
	case float64:
		NBTW.CurrentTag.AddTag(TDouble{Name, Val})
	case []byte:
		NBTW.CurrentTag.AddTag(TByteArray{Name, Val, len(Val), 0})
	case string:
		NBTW.CurrentTag.AddTag(TString{Name, Val})
	case map[string]interface{}: //All types, iterated through, this will not accept compounds or lists since they cannot be detected via primitives
		var err error
		for i, v := range Val {
			Log.Info("key: ", i, " Value: ", v)
			e := NBTW.AddTagBasedOnValue(i, v)
			if e != nil {
				err = e
			}
		}
		return err
	case []int32:
		NBTW.CurrentTag.AddTag(TIntArray{Name, Val, len(Val), 0})
	case []int64:
		NBTW.CurrentTag.AddTag(TLongArray{Name, Val, len(Val), 0})
	default:
		return fmt.Errorf("NBTTypeUnknown: %s, Value: %v", Name, Val)
	}
	return nil
}

// AddMultipleTags this uses T{TagType} objects as tags require a name unless in a list
func (NBTW *NBTWriter) AddMultipleTags(Val []interface{}) {
	for _, v := range Val {
		NBTW.AddTag(v)
	}
}

/* AddMultipleTagsBasedOnValue this uses primitive data types to create a tag
byte, int16, int32, int64, float32, float64, string
This function acts identical to using a map[string]interface in AddTagBasedOnValue
Just that a map now doesn't have to be created*/
func (NBTW *NBTWriter) AddMultipleTagsBasedOnValue(Name []string, Val []interface{}) error {
	var Err error
	for i, v := range Val {
		err := NBTW.AddTagBasedOnValue(Name[i], v)
		if err != nil {
			Err = err //Update Err to err, this ensures it doesn't get replaced with nil
		}
	}
	return Err
}

func (NBTW *NBTWriter) AppendByteSlice(Data []byte) {
	//if cap(NBTW.Data)-len(Data) >= 10 {
	NBTW.Data = append(NBTW.Data, Data...)
	NBTW.size += len(Data)
	//}
	// NBTW.Data = append(NBTW.Data, make([]byte, 0, len(Data)+256)...)
	// NBTW.Data = append(NBTW.Data, Data...)
	// NBTW.size += len(Data)
}

func (NBTW *NBTWriter) writeTag(Type byte, Name string) {
	NBTW.AppendByteSlice([]byte{Type})
	NBTW.writeString("", Name) //This works with the tests with 30% coverage I believe it's fine to do this
}
