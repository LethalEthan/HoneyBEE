package nbt

import "testing"

//var TestData = []byte{10, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 8, 0, 4, 110, 97, 109, 101, 0, 9, 66, 97, 110, 97, 110, 114, 97, 109, 97, 0}

//var NBTR = CreateNBTReader(TestData)

func BenchmarkNBTReader(b *testing.B) {
	NBTW := CreateNBTWriter("hello world")
	NBTW.AddTag(CreateStringTag("hello", "bruh"))
	NBTW.AddTag(CreateStringTag("bruh", "hello"))
	NBTW.EndCompoundTag()
	NBTW.Encode()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NBTR, err := CreateNBTReader(NBTW.Data)
		if err != nil {
			panic(err)
		}
		_, err = NBTR.AutoRead()
		if err != nil {
			panic(err)
		}
	}
}

func TestAutoRead(t *testing.T) {
	t.Run("AutoRead", func(t *testing.T) {
		NBTR, err := CreateNBTReader(TestData)
		if err != nil {
			panic(err)
		}
		Bruh, err := NBTR.AutoRead()
		if err != nil {
			panic(err)
		}
		_ = Bruh
		return
	})
}

func TestReadTagType(T *testing.T) {
	T.Run("ReadTagType", func(t *testing.T) {
		NBTR, err := CreateNBTReader(TestData)
		if err != nil {
			panic(err)
		}
		Type, Name, err := NBTR.readTagType()
		if err != nil {
			panic(err)
		}
		if Type != 10 {
			T.Error(NBTType)
		}
		if Name != "hello world" {
			T.Error(NBTName)
		}
	})
}

func TestReadString(T *testing.T) {
	TestString := CreateNBTWriter("hello world")
	TestString.AddTag(CreateStringTag("test", "Hello Bitches"))
	TestString.AddTag(CreateIntTag("test2", 0))
	TestString.EndCompoundTag()
	TestString.Encode()
	T.Run("TestReadString", func(t *testing.T) {
		NBTR, err := CreateNBTReader(TestString.Data)
		if err != nil {
			panic(err)
		}
		Type, Name, err := NBTR.readTagType()
		if err != nil {
			panic(err)
		}
		if Type != TagCompound {
			panic(NBTType)
		}
		if Name != "hello world" {
			panic(NBTName)
		}
		Type, Name, err = NBTR.readTagType()
		if err != nil {
			T.Error(err)
		}
		if Type != TagString {
			T.Error(NBTType)
		}
		if Name != "test" {
			T.Error(NBTValue)
		}
		String, err := NBTR.readString()
		if err != nil {
			T.Error(err)
		}
		if String != "Hello Bitches" {
			T.Error(NBTValue)
		}
	})
}

func TestReadShort(T *testing.T) {
	TestShort := CreateNBTWriter("hello world")
	TestShort.AddTag(CreateShortTag("bing", 2048))
	TestShort.AddTag(CreateShortTag("bong", 1024))
	TestShort.Encode()
	T.Run("TestReadShort", func(T *testing.T) {
		NBTR, err := CreateNBTReader(TestShort.Data)
		if err != nil {
			panic(err)
		}
		//TestCompoundTag
		Type, Name, err := NBTR.readTagType()
		if err != nil {
			panic(err)
		}
		if Type != 10 {
			T.Error(NBTType)
		}
		if Name != "hello world" {
			T.Error(NBTName)
		}
		//TestShortTag
		Type, Name, err = NBTR.readTagType()
		if err != nil {
			panic(err)
		}
		if Type != TagShort {
			T.Error(NBTType)
		}
		if Name != "bing" {
			T.Error(NBTValue)
		}
		Short, err := NBTR.readShort()
		if err != nil {
			panic(err)
		}
		if Short != 2048 {
			T.Error(NBTValue)
		}
	})
}

func ResetNBTR(NBTR *NBTReader) {
	NBTR.seeker = 0
	NBTR.data = TestData
	NBTR.end = len(NBTR.data)
}

func TestSeek(T *testing.T) {
	TestShort := CreateNBTWriter("hello world")
	TestShort.AddTag(CreateShortTag("bing", 2048))
	TestShort.AddTag(CreateShortTag("bong", 1024))
	TestShort.Encode()
	NBTR, err := CreateNBTReader(TestShort.Data)
	if err != nil {
		panic(err)
	}
	BA, err := NBTR.readWithEOFSeek(100000)
	if err == nil {
		panic(NBTExpectedError)
	}
	if len(BA) != 1 {
		panic(NBTExpectedError)
	}
}

func TestCNBTR(T *testing.T) {
	NBTR, err := CreateNBTReader([]byte{0, 0})
	if err == nil {
		T.Error(NBTExpectedError)
	}
	if NBTR != nil {
		T.Error(NBTExpectedError)
	}
}
