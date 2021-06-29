package npacket

type Slot struct {
	Present   bool
	ItemID    int32
	ItemCount int32
	NBT       NBT
}
