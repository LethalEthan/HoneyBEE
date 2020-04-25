package player

import "Packet"

func SetHotbarSlot(Conn *ClientConnection) {
	writer := Packet.CreatePacketWriter(0x40)
	writer.WriteByte(0)
	SendData(Conn, writer)
}
