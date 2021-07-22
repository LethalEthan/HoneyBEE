package npacket

///
///Serverbound
///

type (
	//Play 0x26
	UpdateCommandBlock_SB struct {
		Location Position
		Command  string
		Mode     int32
		Flags    byte
	}

	//Play 0x27
	UpdateCommandBlockMinecart_SB struct {
		EntityID    int32
		Command     string
		TrackOutput string
	}
)
