package npacket

import (
	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var err error

var Log = logging.MustGetLogger("HoneyGO")

type GeneralPacket struct {
	PacketSize   int32
	PacketID     int32
	PacketData   []byte
	State        int
	OptionalData interface{}
}

//Currently weighing if it's worth to use an interface or just use struct methods, currently just gonna use struct methods for ease

// type PacketCodec interface {
// 	Encode() *PacketWriter
// 	Decode() //*GeneralPacket
// 	//Create()
// }

func (PW *PacketWriter) Send(Conn *gnet.Conn) {
	return
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
