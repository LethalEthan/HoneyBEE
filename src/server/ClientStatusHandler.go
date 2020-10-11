package server

import (
	"jsonstruct"
	"sync"
)

type ServerStatus struct {
	Version     StatusVersion           `json:"version"`
	Players     StatusPlayers           `json:"players"`
	Description jsonstruct.StatusObject `json:"description"`
	Favicon     string                  `json:"favicon,omitempty"`
}

type StatusVersion struct {
	Name     string `json:"name"`
	Protocol int32  `json:"protocol"`
}

type StatusPlayers struct {
	MaxPlayers    int32 `json:"max"`
	OnlinePlayers int32 `json:"online"`
	//SamplePlayer  []SPlayer `json:"sample, omitempty"`
}

var (
	//Use a cache so we don't have to do any uneccesary allocations
	StatusCache = make(map[int32]ServerStatus)
	StatusMutex = sync.RWMutex{}
)

//CreateStatusObject - Create the server status object
func CreateStatusObject(MinecraftProtocolVersion int32, MinecraftVersion string) *ServerStatus {
	SC, bool := CheckStatusCache(MinecraftProtocolVersion, MinecraftVersion)
	if bool == true && SC != nil {
		return SC
	}
	if Config.Server.DEBUG {
		if bool == false || SC == nil {
			Log.Debug("Cache miss")
		}
	}
	status := new(ServerStatus)
	status.Version = StatusVersion{Name: MinecraftVersion, Protocol: MinecraftProtocolVersion}
	status.Players = StatusPlayers{MaxPlayers: MPlayers, OnlinePlayers: OPlayers}
	Extra := make([]jsonstruct.StatusObject, 1)
	Extra[0] = jsonstruct.StatusObject{Text: "GO!", Bold: true, Color: "gold"}
	status.Description = jsonstruct.StatusObject{Text: "Honey", Bold: true, Color: "yellow", Extra: Extra}
	PutStatusInCache(*status, MinecraftProtocolVersion)
	return status
}

func PutStatusInCache(SS ServerStatus, MCP int32) {
	StatusMutex.Lock()
	StatusCache[MCP] = SS
	StatusMutex.Unlock()
}

func CheckStatusCache(MCP int32, MCV string) (*ServerStatus, bool) {
	StatusMutex.RLock()
	SS, B := StatusCache[MCP]
	StatusMutex.RUnlock()
	if B != true {
		return nil, false
	}
	if SS.Players.OnlinePlayers != OPlayers || SS.Players.MaxPlayers != MPlayers {
		//SS := CreateStatusObject(MCP, MCV)
		return nil, false
	}
	return &SS, true
}

//TBD properly
var (
	MPlayers int32 = 420
	OPlayers int32 = 69
)

func OnDisconnect() {
	OPlayers--
}
