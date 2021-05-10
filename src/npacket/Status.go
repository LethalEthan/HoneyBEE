package npacket

import (
	"server"

	"github.com/pquerna/ffjson/ffjson"
)

///
///CLientBound S->C
///

type Status_0x00_CB struct {
	Packet          *GeneralPacket
	Response        *server.ServerStatus
	ProtocolVersion int32
}

type Status_0x01_CB struct {
	Packet *GeneralPacket
	Pong   int64
}

///
///ServerBound C->S
///

type Status_0x00_SB struct {
	Packet *GeneralPacket
}

type Status_0x01_SB struct {
	Packet *GeneralPacket
	Ping   int64
}

func (SP *Status_0x01_SB) Decode() {
	PR := CreatePacketReader(SP.Packet.PacketData)
	var err error
	SP.Ping, err = PR.ReadLong()
	if err != nil {
		panic(err)
	}
}

func (SP *Status_0x00_CB) Encode() *PacketWriter {
	//if SP.Packet.PacketSize == 1 {
	writer := CreatePacketWriter(0x00)
	SP.Response = server.CreateStatusObject(SP.ProtocolVersion, "1.16.5")
	Log.Debug("ProtoVER: ", SP.ProtocolVersion)
	marshaledStatus, err := ffjson.Marshal(SP.Response) //Sends status via json
	if err != nil {
		Log.Error(err)
		return writer
	}
	writer.WriteString(string(marshaledStatus))
	return writer
}

func (SP *Status_0x01_CB) Encode() *PacketWriter {
	// Pong =
	// SP.Pong
	//PW := Packet.CreatePacketWriter(0x01)
	//PW.WriteLong(SP.Packet.OptionalData.(int64))
	SP.Pong = SP.Packet.OptionalData.(int64)
	writer := CreatePacketWriter(0x01)
	writer.WriteLong(SP.Pong)
	return writer
}
