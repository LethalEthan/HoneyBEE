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

//CreateClientConnection - Creates a client Connection, checks whether Connection previously existed and Restores the Connection and terminates due to a suspected failure
func CreateClientConnection(Conn net.Conn, State int) *ClientConnection {
	//Initialise Map
	if ClientConnectionMap == nil {
		ClientConnectionMap = make(map[string]*ClientConnection)
	}
	//Splittity Sploot
	RemoteAddress := strings.Split(Conn.RemoteAddr().String(), ":")[0]

	_, ConnectionExists := ClientConnectionMap[RemoteAddress]

	if ConnectionExists {
		Log.Debug("Client Connection Restored :), therefore you shall be... TERMINATED")
		tmp := ClientConnectionMap[RemoteAddress]
		tmp.Conn.Close()
		tmp.Conn = Conn
		ClientConnectionMap[RemoteAddress] = tmp

	} else {
		Log.Debug("Client Connection Created")
		Connection := new(ClientConnection)
		Connection.Conn = Conn
		Connection.State = State
		ClientConnectionMap[RemoteAddress] = Connection
	}

	return ClientConnectionMap[RemoteAddress]
}

//DestroyClientConnection - Destroys a Connection safely by terminating the Connection and deleting the Connection from the ClientConnectionMap
func DestroyClientConnection(Connection *ClientConnection) {
	if ClientConnectionMap == nil {
		ClientConnectionMap = make(map[string]*ClientConnection)
	}
	Log.Debug("Client Connection destroyed")
	RemoteAddress := strings.Split(Connection.Conn.RemoteAddr().String(), ":")[0]
	Connection.Conn.Close()
	Connection.isClosed = true
	delete(ClientConnectionMap, RemoteAddress)
}

//KeepAlive - Extends the deadline before the Connection is closed
func (c *ClientConnection) KeepAlive() {
	c.Conn.SetDeadline(time.Now().Add(time.Second * 5))
}
