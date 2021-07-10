package npacket

import "jsonstruct"

type Matches struct {
	Match      string
	HasToolTip bool
	ToolTip    jsonstruct.ChatComponent
}
