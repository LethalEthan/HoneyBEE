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

	p.length = length
	//Protocol Version
	p.ProtocolVersion, err = reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	//Notchian Servers don't use this info
	p.ServerAddress, err = reader.ReadString()
	if err != nil {
		p.ServerAddress = ""
	}
	//Notchian Servers don't use this info
	p.ServerPort, err = reader.ReadUnsignedShort()
	if err != nil {
		p.ServerPort = 25565
	}
	//1, status | 2, login
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
