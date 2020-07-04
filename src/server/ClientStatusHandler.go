package server

import "jsonstruct"

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

//Implement RGB for 1.16 later and support multiple client/protocol versions
func CreateStatusObject() *ServerStatus {
	status := new(ServerStatus)
	//Ref: Server.go: constants for status.Version
	status.Version = StatusVersion{Name: MinecraftVersion, Protocol: MinecraftProtocolVersion}
	status.Players = StatusPlayers{MaxPlayers: MPlayers, OnlinePlayers: OPlayers}
	Extra := make([]jsonstruct.StatusObject, 1)
	Extra[0] = jsonstruct.StatusObject{Text: "GO!", Bold: true, Color: "gold"}
	status.Description = jsonstruct.StatusObject{Text: "Honey", Bold: true, Color: "yellow", Extra: Extra}
	return status
}

//TBD properly
var (
	MPlayers int32 = 420
	OPlayers int32 = 69
)

func OnDisconnect() {

}
