package Packet

type HandshakePacket struct {
	length          int32
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func HandshakePacketCreate(length int32, reader *PacketReader) (*HandshakePacket, error) {
	p := new(HandshakePacket)
	var err error
	//Set Packet length
	p.length = length
	//
	p.ProtocolVersion, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	}

	p.ServerAddress, err = reader.ReadString()
	if err != nil {
		return nil, err
	}

	p.ServerPort, err = reader.ReadUnsignedShort()
	if err != nil {
		return nil, err
	}

	p.NextState, err = reader.ReadVarInt()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (p HandshakePacket) Length() int32 {
	return p.length
}

func (p HandshakePacket) PacketID() int32 {
	return 0x00
}
