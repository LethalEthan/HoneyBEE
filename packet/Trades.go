package packet

type Trades struct {
	InputItem1      []Slot
	OutputItem      []Slot
	HasSecondItem   bool
	InputItem2      []Slot
	TradeDisabled   bool
	NumberTradeUses int32
	MaxNumTradeUses int32
	XP              int32
	SpecialPrice    int32
	PriceMultiplier float32
	Demand          int32
}
