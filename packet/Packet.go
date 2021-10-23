package packet

import (
	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyBEE")

type GeneralPacket struct {
	PacketSize   int32
	PacketID     int32
	PacketData   []byte
	State        int
	OptionalData interface{}
}

//Currently weighing if it's worth to use an interface or just use struct methods, currently just gonna use struct methods for ease

type PacketCodec interface {
	Encode() *PacketWriter
	Decode() *PacketReader //*GeneralPacket
}

func Encode(p PacketCodec) {
	p.Encode()
}

func Decode(p PacketCodec) {
	p.Decode()
}

func (PW *PacketWriter) Send(Conn *gnet.Conn) {
}

// func (GP *GeneralPacket) Create() {
// 	switch GP.PacketID {
// 	case 0x00:
// 		P := new(Handshake_0x00)
// 	}
// }

// func (HP *Handshake_0x00) Encode() {
//
// }
//
// func (HP *Handshake_0x00) Decode() {
//
// }
