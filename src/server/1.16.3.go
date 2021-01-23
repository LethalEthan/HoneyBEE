package server

import (
	"Packet"

	"github.com/pquerna/ffjson/ffjson"
)

func Handle_MC1_16_3(Connection *ClientConnection, PH PacketHeader) {
	Log.Info("Connection handler for MC 1.16.3 initiated")
	CurrentStatus = CreateStatusObject(753, "1.16.3")
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
			switch PH.packetID {
			case 0x00:
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
			case 0x01:
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
			default:
				Log.Critical("Unkown packet: ", PH.packet, "PHSize: ", PH.packetSize, "state: ", LOGIN)
				Log.Critical("Contains: ", PH.packet)
			}
		case LOGIN: //Handle Login
			switch PH.packetID {
			case 0x00:
				//--Packet 0x00 C->S Start--// Login Start (Player Username)
				Log.Debug("Login State, packetID 0x00")
				SendLoginDisconnect(Connection, "Currently there isn't enough play logic to continue :(")
				return
				//Connection.KeepAlive()
				playername, _ = reader.ReadString()
				//--Packet 0x01 S->C --// Encryption Request
				CreateEncryptionRequest(Connection)
			case 0x01:
				//--Packet 0x01 S->C Start--//
				//EncryptionResponse
				HandleEncryptionResponse(PH, Connection)
				return
			case 0x02:
				Log.Debug("Login State, packet 0x02")
			case 0x05:
				Log.Debug("Packet 0x05")
			default:
				Log.Critical("Unkown packet: ", PH.packet, "PHSize: ", PH.packetSize, "state: ", LOGIN)
				Log.Critical("Contains: ", PH.packet)
			}
		//Play will be handled by another package/function
		case PLAY:
			switch PH.packetID {
			case 0x00:
				Log.Debug("Play State, packet 0x00")
			case 0x01:
				Log.Debug("Play State, Packet 0x01")
			case 0x02:
				Log.Debug("Play State, Packet 0x02")
			case 0x03:
				Log.Debug("Play State, Packet 0x03")
			case 0x04:
				Log.Debug("Play State, Packet 0x04")
			case 0x05:
				Log.Debug("Play State, Packet 0x05")
			default:
				Log.Critical("Unkown packet: ", PH.packet, "PHSize: ", PH.packetSize, "state: ", PLAY)
				Log.Critical("Contains: ", PH.packet)
			}
		default:
			Log.Critical("Unkown packet: ", PH.packet, "PHSize: ", PH.packetSize, "state: Unknown")
			Log.Critical("Contains: ", PH.packet)
		}
	}
}

//Packet0x34 - Player look and position
func Packet0x34() *Packet.PacketWriter {
	PW := Packet.CreatePacketWriter(0x34)
	PW.WriteDouble(0.0)
	PW.WriteDouble(0.0)
	PW.WriteDouble(0.0)
	PW.WriteFloat(0.0)
	PW.WriteFloat(0.0)
	PW.WriteByte(0x00)
	PW.WriteVarInt(1)
	return PW
}
