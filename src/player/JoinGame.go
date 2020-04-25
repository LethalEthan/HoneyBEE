package player

import (
	"Packet"
	"net"
	"runtime"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("HoneyGO")

const (
	Nether    = -1
	Overworld = 0
	End       = 1
)

var Dimension = Overworld

//
type ClientConnection struct {
	Conn     net.Conn
	State    int
	isClosed bool
}

//GameJoin - Structure of the JoinGame packet
type GameJoin struct {
	EntityID            int32  //Players EntityID
	GameMode            uint8  //0: Survival, 1: Creative, 2: Adventure, 3: Spectator. Bit 3 (0x8) is the hardcore flag.
	Dimension           int    //See connstants above
	HashedSeed          int64  //First 8 bytes of the SHA-256 hash of world seed
	MaxPlayers          uint8  //Used to be used but according to wiki.vg it's no longer used
	LevelType           string //Max 16 length: default, flat, largeBiomes etc
	ViewDistance        VarInt //RenderDistance (2-32)
	ReducedDebugInfo    bool
	EnableRespawnScreen bool //Set false when doImmediateRespawn gamerule is true
}

func CreateGameJoin(Conn *ClientConnection) { //, C chan bool) {
	Conn.Conn.SetDeadline(time.Now().Add(time.Second * 5)) //KeepAlive
	GJ := &GameJoin{2, Creative, 0, 12345, 20, "default", 16, false, true}
	log.Debug("GJ:", GJ)
	//No easy way to do this without this mess, a packet system re-write will be done in the future
	writer := Packet.CreatePacketWriter(0x26)
	writer.WriteInt(GJ.EntityID)
	writer.WriteUnsignedByte(GJ.GameMode)
	writer.WriteInt(int32(GJ.Dimension))
	writer.WriteLong(GJ.HashedSeed)
	writer.WriteUnsignedByte(0)
	writer.WriteString("default")
	writer.WriteVarInt(16)
	writer.WriteBoolean(false)
	writer.WriteBoolean(true)
	SendData(Conn, writer)
	log.Debug("GameJoin Packet sent, Sending SetDifficulty packet")
	log.Debug("GOR:", runtime.NumGoroutine())
	go CreateSetDiff(Conn) //Creates SetDifficultyPacket
	//C <- true
}

func SendData(Connection *ClientConnection, writer *Packet.PacketWriter) {
	Connection.Conn.Write(writer.GetPacket())
}

func (Conn *ClientConnection) KeepAlive() {
	Conn.Conn.SetDeadline(time.Now().Add(time.Second * 10))
}

// func fetchtype(t *GameJoin) {
// 	fmt.Print(reflect.TypeOf(t.EntityID))
//
// }
