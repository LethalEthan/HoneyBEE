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
	Bold          bool                 `json:"bold,omitempty"`
	Italic        bool                 `json:"italic,omitempty"`
	Underlined    bool                 `json:"underlined,omitempty"`
	Strikethrough bool                 `json:"strikethrough,omitempty"`
	Obfuscated    bool                 `json:"obfuscated,omitempty"`
	Colour        string               `json:"color,omitempty"`
	Insertion     string               `json:"insertion,omitempty"`
	ClickEvent    *ClickEventComponent `json:"clickevent,omitempty"`
	HoverEvent    *HoverEventComponnet `json:"hoverevent,omitempty"`
	Extra         []*ChatComponent     `json:"extra,omitempty"`
}

type ChatComponent struct {
	//Shared *SharedComponent
	Text          string               `json:"text"`
	Bold          bool                 `json:"bold,omitempty"`
	Italic        bool                 `json:"italic,omitempty"`
	Underlined    bool                 `json:"underlined,omitempty"`
	Strikethrough bool                 `json:"strikethrough,omitempty"`
	Obfuscated    bool                 `json:"obfuscated,omitempty"`
	Colour        string               `json:"color,omitempty"`
	Insertion     string               `json:"insertion,omitempty"`
	ClickEvent    *ClickEventComponent `json:"clickevent,omitempty"`
	HoverEvent    *HoverEventComponnet `json:"hoverevent,omitempty"`
	Extra         []*ChatComponent     `json:"extra,omitempty"`
}

type ClickEventComponent struct {
}

type HoverEventComponnet struct {
}
