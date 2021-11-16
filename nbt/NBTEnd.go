package nbt

type End struct{}

func (NBTE *NBTEncoder) EncodeEnd() {
	NBTE.data = append(NBTE.data, TagEnd)
}
