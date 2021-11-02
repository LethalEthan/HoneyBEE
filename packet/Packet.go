package packet

import (
	logging "github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyBEE")

type GeneralPacket struct {
	PacketSize   int32
	PacketID     int32
	PacketReader *PacketReader
	OptionalData interface{}
}

//Currently weighing if it's worth to use an interface or just use struct methods, currently just gonna use struct methods for ease

type PacketCodec interface {
	Encode() ([]byte, error)
	Decode() error
}

func Encode(p PacketCodec) {
	p.Encode()
}

func Decode(p PacketCodec) {
	p.Decode()
}
