package chunk

import "fmt"

//Chunk Columns are 16*256*16
type ChunkColumn struct {
}

func ReadChunkColumn(Chunk ChunkPacket) {

}
func ReadChunkDataPacket(buffer []byte) {
	x := buffer
	y := buffer
	fmt.Print(x, y)
}

/*
func ReadChunkDataPacket(Buffer []byte) {
	x := (Buffer)
	z := (Buffer)
	full := ReadBool(ChunkSend.FullChunk)
	//Chunk chunk;
	if full {
		//chunk = new Chunk(x, z);
	} else {
		chunk = GetExistingChunk(x, z)
	}
	mask := ReadVarInt(data)
	size := ReadVarInt(data)
	ReadChunkColumn(chunk, full, mask, data.ReadByteArray(size))

	blockEntityCount := ReadVarInt(data)
	for i := 0; i < blockEntityCount; i++ {
		//CompoundTag tag = NBT.ReadCompoundTag(data);
		chunk.AddBlockEntity(tag.GetInt("x"), tag.GetInt("y"), tag.GetInt("z"), tag)
	}

	return chunk
}
*/
