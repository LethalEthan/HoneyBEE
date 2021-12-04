package packet

import "HoneyBEE/utils"

///
///CLientBound S->C
///

//Status_0x00_CB - Response
type Stat_Response struct {
	Response        *ServerStatus
	ProtocolVersion int32
}

//Status_0x01_CB - Pong
type Stat_Pong struct {
	Pong int64
}

///
///ServerBound C->S
///

//Status_0x00_SB - Request
type Stat_Request struct{}

//Status_0x01_SB - Ping
type Stat_Ping struct {
	Ping int64
}

func (SP *Stat_Ping) Decode(PR *PacketReader) error {
	var err error
	SP.Ping, err = PR.ReadLong()
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (SP *Stat_Response) Encode(PW *PacketWriter) {
	PW.ResetData(0x00)
	SP.Response = CreateStatusObject(utils.PrimaryMinecraftProtocolVersion, utils.PrimaryMinecraftVersion)
	Log.Debug("ClientProtocolVersion: ", SP.ProtocolVersion, "ServerProtocolVersion: ", utils.PrimaryMinecraftProtocolVersion)
	marshaledStatus, err := SP.Response.MarshalJSON() //Sends status via json
	if err != nil {
		Log.Error(err)
		return
	}
	PW.WriteString(string(marshaledStatus))
}

func (SP *Stat_Pong) Encode(PW *PacketWriter) {
	PW.ResetData(0x01)
	PW.WriteLong(SP.Pong)
}
