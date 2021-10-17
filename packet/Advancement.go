package packet

import "HoneyBEE/jsonstruct"

type Advancement struct {
	HasParent        bool
	ParentID         *Identifier
	HasDisplay       bool
	DisplayData      *AdvancementDisplay
	NumberOfCriteria int32
	Criteria         struct {
		Key   []Identifier
		Value byte //void
	}
	ArrayLength  int32
	Requirements struct {
		ArrayLength2 int32
		Requirement  []string
	}
}
type AdvancementDisplay struct {
	Title             jsonstruct.ChatComponent
	Description       jsonstruct.ChatComponent
	Icon              Slot
	FrameType         int32
	Flags             int32
	BackgroundTexture *Identifier
	X                 float32
	Y                 float32
}
