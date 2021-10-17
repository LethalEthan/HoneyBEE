package packet

type Handshake_0x00 struct {
	Packet          *GeneralPacket
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (P *Handshake_0x00) Encode() *PacketWriter {
	return nil
}

func (P *Handshake_0x00) Decode() error {
	var err error
	PR := CreatePacketReader(P.Packet.PacketData)
	//fmt.Print("pl: ", P.Packet.PacketSize, "pd: ", P.Packet.PacketData)
	//var NR uint32
	P.ProtocolVersion, _, err = PR.ReadVarInt()
	if err != nil {
		return err
	}
	//print("PV: ", P.ProtocolVersion, "NR: ", NR)
	P.ServerAddress, err = PR.ReadString()
	if err != nil {
		return err
	}
	P.ServerPort, err = PR.ReadUnsignedShort()
	if err != nil {
		return err
	}
	P.NextState, _, err = PR.ReadVarInt()
	if err != nil {
		return err
	}
	return nil
	//print("DECODED: ", "PV: ", P.ProtocolVersion, " SA: ", P.ServerAddress, " SP: ", P.ServerPort, " NS: ", P.NextState)
}
