package server

import (
	"Packet"
	"net"
	"strings"
	"sync"
	"time"
)

//ClientConnection - Simple struct for Client Connections
type ClientConnection struct {
	Conn      net.Conn
	State     int
	isClosed  bool
	Direction string
	PMAP      sync.Map
	ID        int
	Packet    []Packet.PacketWriter
}

var ClientConnectionMap = make(map[string]*ClientConnection)
var ClientConnectionMutex = sync.RWMutex{}

//CreateClientConnection - Creates a client Connection, checks whether Connection previously existed and Restores the Connection and terminates due to a suspected failure
func CreateClientConnection(Conn net.Conn, State int) *ClientConnection {
	//Splittity Sploot
	RemoteAddress := strings.Split(Conn.RemoteAddr().String(), ":")[0] //Store remote address
	//Check if connection exists
	ClientConnectionMutex.RLock()
	tmp, ConnectionExists := ClientConnectionMap[RemoteAddress]
	ClientConnectionMutex.RUnlock()
	if ConnectionExists {
		Log.Debug("Client Connection Restored, closing")
		tmp.Conn.Close()
		Conn.Close()
		//tmp.Conn = Conn
		//tmp.Conn.Close()
		// ClientConnectionMutex.Lock()
		// ClientConnectionMap[RemoteAddress] = tmp
		// ClientConnectionMutex.Unlock()
		//Check PCM for players~connection
		if val, PCM := GetPCM(tmp.Conn); /*PlayerConnMap[tmp.Conn]*/ PCM {
			Log.Warning("PCM: Deleted:, ", val)
			PlayerConnMutex.Lock()
			delete(PlayerConnMap, tmp.Conn)
			PlayerConnMutex.Unlock()
			go Disconnect(val)
		} else {
			Log.Warning("No player mapped to connection, ignoring")
			PlayerConnMutex.Lock()
			delete(PlayerConnMap, tmp.Conn)
			PlayerConnMutex.Unlock()
		}
	} else {
		Log.Debug("Client Connection Created", &Conn)
		Connection := new(ClientConnection)
		Connection.Conn = Conn
		Connection.State = State
		ClientConnectionMutex.Lock()
		ClientConnectionMap[RemoteAddress] = Connection
		ClientConnectionMutex.Unlock()
	}
	ClientConnectionMutex.RLock()
	defer ClientConnectionMutex.RUnlock()
	return ClientConnectionMap[RemoteAddress]
}

func RecieveEvent(Conn ClientConnection) {

}

//CloseClientConnection - Destroys a Connection ly by terminating the Connection and deleting the Connection from the ClientConnectionMap
func CloseClientConnection(Connection *ClientConnection) {
	Log.Debug("Closing client connection")
	RemoteAddress := strings.Split(Connection.Conn.RemoteAddr().String(), ":")[0]
	Connection.isClosed = true
	Connection.Conn.Close()
	ClientConnectionMutex.Lock()
	delete(ClientConnectionMap, RemoteAddress)
	ClientConnectionMutex.Unlock()
	//Check PlayerConnectionMap
	if val, PCM := GetPCM(Connection.Conn); /*PlayerConnMap[Connection.Conn]*/ PCM {
		Log.Warning("PCM: Deleted:, ", val)
		PlayerConnMutex.Lock()
		delete(PlayerConnMap, Connection.Conn)
		PlayerConnMutex.Unlock()
		//delete(ConnPlayerMap, val)
	} /*else {
		Log.Warning("No player mapped to connection, ignoring")
	}*/
	Log.Debug("Client Connection destroyed")
	Connection = nil
}

//KeepAlive - Extends the deadline before the Connection is closed
func (c *ClientConnection) KeepAlive() {
	c.Conn.SetDeadline(time.Now().Add(time.Second * 5))
}

//var PMAP sync.Map
func (C *ClientConnection) Send(Packet *Packet.PacketWriter) {
	C.PMAP.Store(C.ID, Packet)
	C.ID++
}

// func Sender() {
// 	for key, val := range
// }

// func (CCM *ClientConnectionMap) Get(key string) (*ClientConnection, bool) {
//
// }
