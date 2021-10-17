package packet

import (
	"HoneyBEE/config"
	"HoneyBEE/jsonstruct"
	"HoneyBEE/utils"
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

// var (
// 	//StatusCache - Use a cache so we don't have to do any uneccesary allocations
// 	StatusCache     = make(map[int32]ServerStatus)
// 	StatusMutex     = sync.RWMutex{}
// 	StatusSemaphore = utils.CreateSemaphore(10) //allow 10 concurrent connections to the StatusCache
// )

var LastStatus ServerStatus

//CreateStatusObject - Create the server status object
func CreateStatusObject(MinecraftProtocolVersion int32, MinecraftVersion string) *ServerStatus {
	//Limit the range of protocols to prevent the cache being flooded by false requests in attempt to crash the server maliciously
	if LastStatus.Version.Protocol == MinecraftProtocolVersion {
		Log.Debug("Cache")
		return &LastStatus
	}
	if MinecraftProtocolVersion > 1200 || MinecraftProtocolVersion < 500 {
		Log.Info("Protocol OOB, setting to 1.17!")
		MinecraftProtocolVersion = utils.PrimaryMinecraftProtocolVersion
		MinecraftVersion = utils.PrimaryMinecraftVersion
	}
	//StatusSemaphore.SetData(StatusCache)
	// SC, bool := CheckStatusCache(MinecraftProtocolVersion, MinecraftVersion)
	// if bool == true && SC != nil {
	// 	return SC
	// }
	// if config.GConfig.Server.DEBUG {
	// 	if bool == false || SC == nil {
	// 		Log.Debug("Cache miss")
	// 	}
	// }
	status := new(ServerStatus)
	status.Version = StatusVersion{Name: MinecraftVersion, Protocol: MinecraftProtocolVersion}
	status.Players = StatusPlayers{MaxPlayers: MPlayers, OnlinePlayers: OPlayers}
	Extra := make([]jsonstruct.StatusObject, 1)
	if config.GConfig.DEBUGOPTS.Maintenance {
		Extra[0] = jsonstruct.StatusObject{Text: "!", Bold: true, Color: "gold"}
		status.Description = jsonstruct.StatusObject{Text: "Server under maintenance", Bold: true, Color: "red", Extra: Extra}
	} else {
		Extra[0] = jsonstruct.StatusObject{Text: "GO!", Bold: true, Color: "gold"}
		status.Description = jsonstruct.StatusObject{Text: "Honey", Bold: true, Color: "yellow", Extra: Extra}
	}
	//PutStatusInCache(*status, MinecraftProtocolVersion)
	return status
}

// func PutStatusInCache(SS ServerStatus, MCP int32) {
// 	//StatusSemaphore.FlushAndSemaphore() //FlushSemaphore so new data can served
// 	StatusMutex.Lock()
// 	StatusCache[MCP] = SS
// 	StatusSemaphore.FlushAndSetSemaphoreNEW(StatusCache)
// 	StatusMutex.Unlock()
// 	//StatusSemaphore.SetData(StatusCache) //Set new data
// }

// func CheckStatusCache(MCP int32, MCV string) (*ServerStatus, bool) {
// 	SC := StatusSemaphore.GetData()
// 	if SC == nil {
// 		Log.Critical("Semaphore data interface is nil!")
// 		return nil, false
// 	}
// 	MB := SC.(map[int32]ServerStatus)
// 	//StatusMutex.RLock()
// 	SS, B := MB[MCP]
// 	//StatusMutex.RUnlock()
// 	if B != true {
// 		return nil, false
// 	}
//
// 	if SS.Players.OnlinePlayers != OPlayers || SS.Players.MaxPlayers != MPlayers {
// 		//SS := CreateStatusObject(MCP, MCV)
// 		return nil, false
// 	}
// 	return &SS, true
// }

//TBD properly
var (
	MPlayers int32 = 420
	OPlayers int32 = 69
)

func OnDisconnect() {
	OPlayers--
}
