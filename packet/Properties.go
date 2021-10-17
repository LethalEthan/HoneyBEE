package packet

type Property struct {
	Key               Identifier
	Value             float64
	NumberOfModifiers int32
	Modifiers         []ModifierFormat
}

type ModifierFormat struct {
	Operation byte
	Value     int32
}
