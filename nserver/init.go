package nserver

import (
	"HoneyGO/world"
)

var GCPShutdown = make(chan bool)

//Bri-ish init bruv
func Init() {
	// go server.StatusSemaphore.Start()
	// server.StatusSemaphore.FlushAndSetSemaphore(server.StatusCache)
	// server.CurrentStatus = server.CreateStatusObject(utils.PrimaryMinecraftProtocolVersion, utils.PrimaryMinecraftVersion)
	// player.Init()
	go world.Init()
	//go player.GCPlayer(GCPShutdown)
	if DEBUG {
		Log.Debug("Server initialised")
	}
}
