package player

import "Packet"

//var log = logging.MustGetLogger("HoneyGO")

type SetDifficulty struct {
	Difficulty uint8 //0:Peaceful, 1:easy, 2:normal, 3:hard
	DiffLock   bool  //Difficulty Lock
}

func CreateSetDiff(Conn *ClientConnection) {
	SD := &SetDifficulty{0, false} //new(SetDifficulty)
	writer := Packet.CreatePacketWriter(0x0E)
	writer.WriteUnsignedByte(SD.Difficulty)
	writer.WriteBoolean(SD.DiffLock)
	SendData(Conn, writer)
}
