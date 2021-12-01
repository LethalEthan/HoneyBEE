package packet

func (PM *PluginMessage_CB) Encode(PW *PacketWriter, Channel Identifier, Data []byte) {
	PW.ResetData(0x18)
	PW.WriteIdentifier(PM.Channel)
	PW.WriteArray(PM.Data)
	Log.Debug("Created PM")
}
