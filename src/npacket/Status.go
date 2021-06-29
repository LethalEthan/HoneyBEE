package npacket

import (
	"server"

	"github.com/pquerna/ffjson/ffjson"
)

///
///CLientBound S->C
///

//Status_0x00_CB - Response
type Stat_Response struct {
	Packet          *GeneralPacket
	Response        *server.ServerStatus
	ProtocolVersion int32
}

//Status_0x01_CB - Pong
type Stat_Pong struct {
	Packet *GeneralPacket
	Pong   int64
}

///
///ServerBound C->S
///

//Status_0x00_SB - Request
type Stat_Request struct {
	Packet *GeneralPacket
}

//Status_0x01_SB - Ping
type Stat_Ping struct {
	Packet *GeneralPacket
	Ping   int64
}

func (SP *Stat_Ping) Decode() {
	PR := CreatePacketReader(SP.Packet.PacketData)
	var err error
	SP.Ping, err = PR.ReadLong()
	if err != nil {
		panic(err)
	}
}

func (SP *Stat_Response) Encode() *PacketWriter {
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

func (SP *Stat_Pong) Encode(PingData int64) *PacketWriter {
	//SP.Pong = SP.Packet.OptionalData.(int64)
	writer := CreatePacketWriter(0x01)
	writer.WriteLong(SP.Pong)
	return writer
}
