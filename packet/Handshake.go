package packet

type Handshake_0x00 struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (P *Handshake_0x00) Decode(PR *PacketReader) error {
	var err error
	if P.ProtocolVersion, _, err = PR.ReadVarInt(); err != nil {
		return err
	}
	if P.ServerAddress, err = PR.ReadString(); err != nil {
		return err
	}
	if P.ServerPort, err = PR.ReadUShort(); err != nil {
		return err
	}
	if P.NextState, _, err = PR.ReadVarInt(); err != nil {
		return err
	}
	return nil
}
