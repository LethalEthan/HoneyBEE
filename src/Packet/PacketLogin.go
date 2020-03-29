package Packet

type VarInt int32
type LoginPacket struct {
	SharedSecretLength int32
	SharedSecret       []byte
	VerifyTokenLength  VarInt
	VerifyToken        []byte
	t                  error
}

func LoginPacketCreate(length int32, reader *PacketReader) (*LoginPacket, error) {
	lp := new(LoginPacket)
	lp.SharedSecretLength, lp.t = reader.ReadVarInt()
	lp.SharedSecret, lp.t = reader.ReadArray()
	return lp, lp.t
}
