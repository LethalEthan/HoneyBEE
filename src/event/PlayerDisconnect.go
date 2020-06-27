package event

import (
	"player"

	logging "github.com/op/go-logging"
)

var Log *logging.Logger

//Event - Bleh
type Event interface {
	PlayerDisconnect() *player.PlayerObject
}

//PlayerDisconnect - Handle Player Disconnect
func PlayerDisconnect() {

}
