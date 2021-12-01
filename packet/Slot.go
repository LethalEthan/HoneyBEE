package packet

type Slot struct {
	Present   bool
	ItemID    int32
	ItemCount int32
	NBT       NBT
}

type Slots struct {
	SlotNumber int16
	SlotData   int16
}
