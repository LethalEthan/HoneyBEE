package npacket

import "HoneyBEE/jsonstruct"

type Icon struct {
	Type           int32
	X              byte
	Z              byte
	Direction      byte
	HasDisplayName bool
	DisplayName    jsonstruct.ChatComponent
}
