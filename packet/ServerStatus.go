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
	// SamplePlayer  []string `json:"sample, omitempty"`
}

var LastStatus ServerStatus

//CreateStatusObject - Create the server status object
func CreateStatusObject(MinecraftProtocolVersion int32, MinecraftVersion string) *ServerStatus {
	//Limit the range of protocols to prevent the cache being flooded by false requests in attempt to crash the server maliciously
	// if LastStatus.Version.Protocol == MinecraftProtocolVersion {
	// 	Log.Debug("Cache")
	// 	return &LastStatus
	// }
	Log.Debug("STATUS CREATION")
	if MinecraftProtocolVersion > 1200 || MinecraftProtocolVersion < 500 {
		MinecraftProtocolVersion = utils.PrimaryMinecraftProtocolVersion
		MinecraftVersion = utils.PrimaryMinecraftVersion
	}
	status := new(ServerStatus)
	status.Version = StatusVersion{Name: MinecraftVersion, Protocol: MinecraftProtocolVersion}
	status.Players = StatusPlayers{MaxPlayers: MPlayers, OnlinePlayers: OPlayers}
	Extra := make([]jsonstruct.StatusObject, 1)
	if config.GConfig.DEBUGOPTS.Maintenance {
		Extra[0] = jsonstruct.StatusObject{Text: "!", Bold: true, Color: "gold"}
		status.Description = jsonstruct.StatusObject{Text: "Server under maintenance", Bold: true, Color: "red", Extra: Extra}
	} else {
		Extra[0] = jsonstruct.StatusObject{Text: "BEE!", Bold: true, Color: "gold"}
		status.Description = jsonstruct.StatusObject{Text: "Honey", Bold: true, Color: "yellow", Extra: Extra}
	}
	return status
}

//TBD properly
var (
	MPlayers int32 = 420
	OPlayers int32 = 69
)

func OnDisconnect() {
	OPlayers--
}
