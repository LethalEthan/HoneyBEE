package packet

type Handshake_0x00 struct {
	Packet          *GeneralPacket
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (P *Handshake_0x00) Encode() ([]byte, error) {
	return nil, nil
}

func (P *Handshake_0x00) Decode() error {
	var err error
	PR := P.Packet.PacketReader
	if P.ProtocolVersion, _, err = PR.ReadVarInt(); err != nil {
		return err
	}
	if P.ServerAddress, err = PR.ReadString(); err != nil {
		return err
	}
	if P.ServerPort, err = PR.ReadUnsignedShort(); err != nil {
		return err
	}
	if P.NextState, _, err = PR.ReadVarInt(); err != nil {
		return err
	}
	return nil
}
