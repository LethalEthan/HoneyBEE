package jsonstruct

type StatusObject struct {
	Text          string         `json:"text"`
	Bold          bool           `json:"bold,omitempty"`
	Italic        bool           `json:"italic,omitempty"`
	Underlined    bool           `json:"underlined,omitempty"`
	Strikethrough bool           `json:"strikethrough,omitempty"`
	Obfuscated    bool           `json:"obfuscated,omitempty"`
	Color         string         `json:"color,omitempty"`
	Extra         []StatusObject `json:"extra,omitempty"`
}

type SharedComponent struct {
	Bold          bool
	Italic        bool
	Underlined    bool
	Strikethrough bool
	Obfuscated    bool
	Colour        string
	Insertion     string
	ClickEvent    *ClickEventComponent
	HoverEvent    *HoverEventComponnet
	Extra         []*ChatComponent
}

type ChatComponent struct {
	Shared *SharedComponent
	Text   string `json:"text"`
}

type ClickEventComponent struct {
}

type HoverEventComponnet struct {
}
