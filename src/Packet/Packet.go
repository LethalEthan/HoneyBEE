package Packet

//Create - creates a packet that will be used later, stored in a channel
func Create(ID int32) chan PacketWriter {
	c := make(chan PacketWriter, 1)
	f := new(PacketWriter)
	f.packetID = ID
	c <- *f
	return c
}
