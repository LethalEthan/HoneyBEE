package npacket

import logging "github.com/op/go-logging"

var Log = logging.MustGetLogger("HoneyGO")

type GeneralPacket struct {
	PacketSize   int32
	PacketID     int32
	PacketData   []byte
	State        int
	OptionalData interface{}
}

type PacketCodec interface {
	Encode() *PacketWriter
	Decode()
	Create()
	Send()
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
