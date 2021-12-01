package packet

import "testing"

func TestReadRestOfByteArrayNoSeek(T *testing.T) {
	PR := CreatePacketReader([]byte{0, 4, 3, 5, 8, 2, 5, 6, 3, 1, 3, 6, 8, 0, 6, 3, 4, 8, 9, 3, 5, 8})
	PR.seek(4)
	D := PR.ReadRestOfByteArrayNoSeek()
	for i := range D {
		Expected := []byte{8, 2, 5, 6, 3, 1, 3, 6, 8, 0, 6, 3, 4, 8, 9, 3, 5, 8}
		if D[i] != Expected[i] {
			T.Error(ReaderValue)
		}
	}
}

func TestReadLongArray(T *testing.T) {
	PR := CreatePacketReader([]byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10})
	if LA, err := PR.ReadLongArray(4); err == nil {
		_ = LA
	} else {
		T.Error(ReaderValue)
	}
}

func TestLongArrayFaulty(T *testing.T) {
	PR := CreatePacketReader([]byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10})
	if LA, err := PR.ReadLongArray(100); err == nil {
		_ = LA
		T.Error(ReaderExpectedError)
	} else {
		T.Log("VarIntArray fault triggered")
	}
}

func TestVarIntArray(T *testing.T) {
	PR := CreatePacketReader([]byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10})
	if LA, err := PR.ReadVarIntArray(8); err == nil {
		_ = LA
	} else {
		T.Error(ReaderValue)
	}
}

func TestVarIntArrayFaulty(T *testing.T) {
	PR := CreatePacketReader([]byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10})
	if LA, err := PR.ReadVarIntArray(100); err == nil {
		_ = LA
		T.Error(ReaderExpectedError)
	} else {
		T.Log("VarIntArray fault triggered")
	}
}
