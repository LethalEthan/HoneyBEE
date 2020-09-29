package server

import (
	"Packet"
	"encoding/json"
	"player"
	"time"
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
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--// Request Status
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x00)
						marshaledStatus, err := json.Marshal(*CurrentStatus) //Sends status via json
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
						mememode(Connection)
						Connection.KeepAlive()
						playername, _ = reader.ReadString()
						//--Packet 0x01 S->C --// Encryption Request
						CreateEncryptionRequest(Connection)
					}
				case 0x01:
					{
						//--Packet 0x01 S->C Start--//
						//EncryptionResponse
						ClientSharedSecret, err := HandleEncryptionResponse(PH)
						if err != nil {
							CloseClientConnection(Connection)
							return
						}
						//--Authentication Stuff--//
						Auth, err := AuthPlayer(playername, ClientSharedSecret)
						if err != nil {
							Log.Error(err)
							CloseClientConnection(Connection)
						} else {
							Log.Debug(playername, "[", Auth, "]")
						}
						//--Packer 0x01 End--//

						//--Packet 0x02 S->C Start--//
						writer := Packet.CreatePacketWriter(0x02)
						Log.Debug("Playername: ", playername)
						writer.WriteString(Auth)
						writer.WriteString(playername)
						time.Sleep(5000000) //DEBUG:Add delay -- remove me later
						SendData(Connection, writer)

						///Entity ID Handling///
						SetPCMSafe(Connection.Conn, playername) //PlayerConnMap[Connection.Conn] = playername //link connection to player
						player.InitPlayer(playername, Auth /*, player.PlayerEntityMap[playername]*/, 1)
						player.GetPlayerByID(player.PlayerEntityMap[playername])
						EID, _ := player.GetPEMSafe(playername) //player.PlayerEntityMap[playername]
						SetCPMSafe(EID, Connection.Conn)        //ConnPlayerMap[EID] = Connection.Conn
						//--//
						Connection.State = PLAY
						PC := &player.ClientConnection{Connection.Conn, Connection.State, Connection.isClosed}
						player.CreateGameJoin(PC, player.PlayerEntityMap[playername])
						player.CreateSetDiff(PC)
						player.CreatePlayerAbilities(PC)
						Log.Debug("END")
						mememode(Connection)
						CloseClientConnection(Connection)
						Disconnect(playername)
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
