package event

import (
	logging "github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyGO")

type Player string

//Event - Bleh
type Event interface {
	PlayerDisconnect() //*player.PlayerObject
}

//PlayerDisconnect - Handle Player Disconnect
func (player Player) /*player.PlayerObject)*/ PlayerDisconnect() {
	Log.Info("Player: ", player, "Disconnected")
}
