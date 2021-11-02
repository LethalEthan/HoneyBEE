package packet

func (PM *PluginMessage_CB) Encode() PacketWriter {
	PM.Channel = "minecraft:brand"
	PM.Data = []byte("HoneyBEE!")
	PW := CreatePacketWriterWithCapacity(0x18, 32)
	PW.WriteIdentifier(PM.Channel)
	PW.WriteArray(PM.Data)
	Log.Debug("Created PM")
	return PW
}
