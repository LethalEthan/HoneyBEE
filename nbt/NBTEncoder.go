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
}

func CreateNBTEncoder() NBTEncoder {
	NBTE := *new(NBTEncoder)
	NBTE.data = make([]byte, 0, 2048)
	NBTE.rootCompound = new(Compound)
	NBTE.rootCompound.value = make([]interface{}, 0, 16)
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

func (NBTE *NBTEncoder) Encode() {
	NBTE.data = NBTE.data[:0]
	if NBTE.currentCompound.previousTag == nil {
		NBTE.EncodeCompound(NBTE.rootCompound)
	} else {
		Log.Debug("Cannot encode, tags are not fully ended!")
		fmt.Println("Cannot encode, tags are not fully ended!")
	}
}

func (NBTE *NBTEncoder) AddTag(v interface{}) {
	NBTE.currentCompound.value = append(NBTE.currentCompound.value, v)
}
