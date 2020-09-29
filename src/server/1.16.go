package server

import (
	"Packet"

	"github.com/pquerna/ffjson/ffjson"
)

func Handle_MC1_16(Connection *ClientConnection, PH PacketHeader) {
	Log.Info("Connection handler for MC 1.16 initiated")
	CurrentStatus = CreateStatusObject(735, "1.16")
	if publicKey == nil || privateKey == nil {
		panic("Keys have been thanos snapped")
	}
	for !Connection.isClosed {
		var err error
		PH.packet, PH.packetSize, PH.packetID, err = readPacketHeader(Connection)
		if err != nil {
			CloseClientConnection(Connection)
			Log.Error("Connection Terminated: " + err.Error())
			return
		}
		//DEBUG: output debug info
		if DEBUG {
			DisplayPacketInfo(PH, Connection)
		}
		//Create Packet Reader
		reader := Packet.CreatePacketReader(PH.packet)
		//Packet Handling
		switch Connection.State {
		case STATUS: //Handle Status Request
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--// Request Status
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x00)
						marshaledStatus, err := ffjson.Marshal(CurrentStatus) //Sends status via json
						if err != nil {
							Log.Error(err.Error())
							CloseClientConnection(Connection)
							return
						}
						writer.WriteString(string(marshaledStatus))
						SendData(Connection, writer)
					}
				case 0x01:
					{
						//--Packet 0x01 S->C Start--// Ping
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x01)
						Log.Debug("Status State, packetID 0x01")
						mirror, _ := reader.ReadLong()
						Log.Debug("Mirror:", mirror)
						writer.WriteLong(mirror)
						SendData(Connection, writer)
						CloseClientConnection(Connection)
						//--Packet 0x01 End--//
					}
				}
			}
		case LOGIN: //Handle Login
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 C->S Start--// Login Start (Player Username)
						Log.Debug("Login State, packetID 0x00")
						Connection.KeepAlive()
						playername, _ = reader.ReadString()
						//--Packet 0x01 S->C --// Encryption Request
						CreateEncryptionRequest(Connection)
					}
				case 0x01:
					{
						//--Packet 0x01 S->C Start--//
						//EncryptionResponse
						//ClientSharedSecret, err :=
						HandleEncryptionResponse(PH, Connection)
					}
				case 0x02:
					{
						Log.Debug("Login State, packet 0x02")
					}
				case 0x05:
					{
						Log.Debug("Packet 0x05")
					}
				}
			}
			//Play will be handled by another package/function
		case PLAY:
			{
				switch PH.packetID {
				case 0x00:
					{
						Log.Debug("Play State, packet 0x00")
					}
				case 0x01:
					{
						Log.Debug("Play State, Packet 0x01")
					}
				case 0x02:
					{
						Log.Debug("Play State, Packet 0x02")
					}
				case 0x03:
					{
						Log.Debug("Play State, Packet 0x03")
					}
				case 0x04:
					{
						Log.Debug("Play State, Packet 0x04")
					}
				case 0x05:
					{
						Log.Debug("Play State, Packet 0x05")
					}
				default:
					for {
						Log.Fatal("Play Packet recieved")
					}
				}
			}
		default:
			Log.Debug("oo")
		}
	}
}
