package server

import (
	"net"
	"strings"
	"time"
)

//ClientConnection - Simple struct for Client Connections
type ClientConnection struct {
	Conn      net.Conn
	State     int
	isClosed  bool
	Direction string
}

var ClientConnectionMap = make(map[string]*ClientConnection)

//CreateClientConnection - Creates a client Connection, checks whether Connection previously existed and Restores the Connection and terminates due to a suspected failure
func CreateClientConnection(Conn net.Conn, State int) *ClientConnection {
	//Splittity Sploot
	RemoteAddress := strings.Split(Conn.RemoteAddr().String(), ":")[0]
	//Check if connection exists
	_, ConnectionExists := ClientConnectionMap[RemoteAddress]
	if ConnectionExists {
		Log.Debug("Client Connection Restored, closing")
		tmp := ClientConnectionMap[RemoteAddress]
		tmp.Conn.Close()
		tmp.Conn = Conn
		tmp.Conn.Close()
		ClientConnectionMap[RemoteAddress] = tmp
		if val, PCM := PlayerConnMap[tmp.Conn]; PCM {
			Log.Warning("PCM: Deleted:, ", val)
			delete(PlayerConnMap, tmp.Conn)
			//delete(ConnPlayerMap, tmp.Conn)
			go Disconnect(val)
		} else {
			Log.Warning("No player mapped to connection, ignoring")
		}
	} else {
		Log.Debug("Client Connection Created", &Conn)
		Connection := new(ClientConnection)
		Connection.Conn = Conn
		Connection.State = State
		ClientConnectionMap[RemoteAddress] = Connection
	}

	return ClientConnectionMap[RemoteAddress]
}

//CloseClientConnection - Destroys a Connection safely by terminating the Connection and deleting the Connection from the ClientConnectionMap
func CloseClientConnection(Connection *ClientConnection) {
	Log.Debug("Client Connection destroyed")
	RemoteAddress := strings.Split(Connection.Conn.RemoteAddr().String(), ":")[0]
	Connection.Conn.Close()
	Connection.isClosed = true
	delete(ClientConnectionMap, RemoteAddress)
	if val, PCM := PlayerConnMap[Connection.Conn]; PCM {
		Log.Warning("PCM: Deleted:, ", val)
		delete(PlayerConnMap, Connection.Conn)
		//delete(ConnPlayerMap, val)
	} else {
		Log.Warning("No player mapped to connection, ignoring")
	}
}

//KeepAlive - Extends the deadline before the Connection is closed
func (c *ClientConnection) KeepAlive() {
	c.Conn.SetDeadline(time.Now().Add(time.Second * 5))
}
