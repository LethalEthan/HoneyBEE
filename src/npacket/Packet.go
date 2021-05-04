package npacket

type VarInt int32

type GeneralPacket struct {
	PacketSize   int32
	PacketID     int32
	PacketData   []byte
	State        int
	OptionalData interface{}
}

type PacketCodec interface {
	Encode() PacketWriter
	Decode()
	Create()
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
