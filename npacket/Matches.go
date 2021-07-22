package npacket

import "HoneyGO/jsonstruct"

type Matches struct {
	Match      string
	HasToolTip bool
	ToolTip    jsonstruct.ChatComponent
}
