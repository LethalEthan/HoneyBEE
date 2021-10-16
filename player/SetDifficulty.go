package player

import (
	"HoneyBEE/Packet"
)

//var log = logging.MustGetLogger("HoneyBEE")

//SetDifficulty - Packet struct for SetDifficulty
type SetDifficulty struct {
	Difficulty uint8 //0:Peaceful, 1:easy, 2:normal, 3:hard
	DiffLock   bool  //Difficulty Lock
}

//CreateSetDiff - Create SetDifficulty Packet and send
func CreateSetDiff(Conn *ClientConnection) {
	//Conn.KeepAlive()
	//Log := logging.MustGetLogger("HoneyBEE")
	//Log.Debug("Packet Play, 0x0E Created")
	SD := &SetDifficulty{0, true} //new(SetDifficulty)
	writer := Packet.CreatePacketWriter(0x0E)
	writer.WriteUnsignedByte(SD.Difficulty)
	writer.WriteBoolean(SD.DiffLock)
	//wait := <-C
	SendData(Conn, writer)
	//Log.Debug("Difficulty set sent")
	//CreatePlayerAbilities(Conn)
}
