package nbt

import (
	"fmt"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyBEE")

type NBTEncoder struct {
	data            []byte
	rootCompound    *Compound
	currentCompound *Compound
	// output          bytes.Buffer
}

func CreateNBTEncoder() NBTEncoder {
	NBTE := *new(NBTEncoder)
	NBTE.data = make([]byte, 0, 4096)
	NBTE.rootCompound = new(Compound)
	NBTE.rootCompound.value = make([]interface{}, 0, 32)
	NBTE.currentCompound = NBTE.rootCompound
	return NBTE
}

func (NBTE *NBTEncoder) Reset() {
	NBTE.data = NBTE.data[:0]
	NBTE.rootCompound.value = NBTE.rootCompound.value[:0]
	NBTE.currentCompound = NBTE.rootCompound
}

func (NBTE *NBTEncoder) GetData() []byte {
	return NBTE.data
}

func (NBTE *NBTEncoder) GetObjects() []interface{} {
	return NBTE.rootCompound.value
}

func (NBTE *NBTEncoder) SetRootTag(C *Compound) {
	NBTE.rootCompound = C
}

func (NBTE *NBTEncoder) GetRootTag() *Compound {
	return NBTE.rootCompound
}

func (NBTE *NBTEncoder) SetCurrentTag(C *Compound) {
	NBTE.currentCompound = C
}

func (NBTE *NBTEncoder) GetCurrentTag() *Compound {
	return NBTE.currentCompound
}

func (NBTE *NBTEncoder) Encode() []byte {
	// Log.Debug("Encode function")
	NBTE.data = NBTE.data[:0]
	if NBTE.currentCompound.previousTag == nil {
		NBTE.EncodeCompound(NBTE.rootCompound)
	} else {
		Log.Debug("Cannot encode, tags are not fully ended!")
		fmt.Println("Cannot encode, tags are not fully ended!")
	}
	return NBTE.data
}

func (NBTE *NBTEncoder) AddTag(v interface{}) {
	NBTE.currentCompound.value = append(NBTE.currentCompound.value, v)
}

func (NBTE *NBTEncoder) AddMultipleTags(val []interface{}) {
	NBTE.currentCompound.value = append(NBTE.currentCompound.value, val...)
}
