package npacket

///
///ClientBound
///

type UUID string

//Play_0x00_CB -
type (
	Play_0x00_CB struct {
		EntityID   int32
		ObjectUUID UUID
		Type       int32
		X          float64
		Y          float64
		Z          float64
		Pitch      byte
		Yaw        byte
		Data       int
		VelocityX  int16
		VelocityY  int16
		VelocityZ  int16
	}

	//Play_0x01_CB - Spawn Experience Orb
	Play_0x01_CB struct {
		EntityID int32
		X        float64
		Y        float64
		Z        float64
		Count    int16
	}

	//Play_0x02_CB - Spawn Living Entity
	Play_0x02_CB struct {
		EntityID   int32
		EntityUUID UUID
		Type       int32
		X          float64
		Y          float64
		Z          float64
		Yaw        byte
		Pitch      byte
		HeadPitch  byte
		VelocityX  int16
		VelocityY  int16
		VelocityZ  int16
	}

	//Play_0x03_CB - Spawn Painting
	Play_0x03_CB struct {
		EntityID   int32
		EntityUUID UUID
		Motive     int32
		Location   int64
		Direction  byte
	}

	//Play_0x04_CB - Spawn Player
	Play_0x04_CB struct {
		EntityID   int32
		PlayerUUID UUID
		X          float64
		Y          float64
		Z          float64
		Yaw        byte
		Pitch      byte
	}
	//Play_0x05_CB
	Play_0x05_CB struct {
		EntityID  int32
		Animation byte
	}
	//Play_0x06_CB
	Play_0x06_CB struct {
		Count int32
		//Finish me
	}
	//Play_0x07_CB
	Play_0x07_CB struct {
		Location   int64
		Block      int32
		Status     int32
		Successful bool
	}
	//Play0x08_CB
	Play_0x08_CB struct {
		EntityID     int32
		Location     int64
		DestroyStage byte
	}
)
