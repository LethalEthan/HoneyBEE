package npacket

import "HoneyBEE/jsonstruct"

type Matches struct {
	Match      string
	HasToolTip bool
	ToolTip    jsonstruct.ChatComponent
}
