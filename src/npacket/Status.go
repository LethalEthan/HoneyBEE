package npacket

import (
	"server"
)

type Status_0x00_CB struct {
	Packet   *GeneralPacket
	Response *server.ServerStatus
}

type Status_0x01_CB struct {
	Packet *GeneralPacket
	Pong   int64
}

type Status_0x00_SB struct {
	Packet *GeneralPacket
}

type Status_0x01_SB struct {
	Packet *GeneralPacket
	Ping   int64
}

func (SP Status_0x01_SB) Decode() {
	PR := CreatePacketReader(SP.Packet.PacketData)
	var err error
	SP.Ping, err = PR.ReadLong()
	if err != nil {
		panic(err)
	}
}

func (SP Status_0x01_CB) Encode() PacketWriter {
	// Pong =
	// SP.Pong
	//PW := Packet.CreatePacketWriter(0x01)
	//PW.WriteLong(SP.Packet.OptionalData.(int64))
	SP.Pong = SP.Packet.OptionalData.(int64)
	PW := CreatePacketWriter(0x01)
	PW.WriteLong(SP.Pong)
	return *PW
}
