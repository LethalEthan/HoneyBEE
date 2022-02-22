package world

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
)

func GetRegionsFromStorage() {
	// Read dir
	de, err := os.ReadDir("overworld")
	if err != nil {
		err = os.Mkdir("overworld", fs.FileMode(os.O_RDWR)) // create dir on error, kinda ass backwards but idk the specific error to compare to
		if err != nil {
			panic(err) // panic if creation failed
		}
	}
	de, err = os.ReadDir("overworld")
	if err != nil {
		panic(err)
	}
	r, err := regexp.Compile(`[R]+([\-]?[\d]{0,10})+[\,]+([\-]?[\d]{0,10})`)
	if err != nil {
		panic(err)
	}
	Log.Debug(de)
	test := make([]byte, 0, 80000)
	for i := range de {
		if !de[i].IsDir() {
			test = append(test, []byte(de[i].Name())...)
			fmt.Println([]byte(de[i].Name()))
		}
	}
	indexes := r.FindAllIndex(test, -1)
	fmt.Println(indexes)
	for i := 0; i < len(indexes); i++ {
		fmt.Println(string(test[indexes[i][0]:indexes[i][1]]))
	}
}

type region struct {
	ID     RegionID
	Data   []ChunkColumn
	Loaded bool
	// ChunkModify chan chunk.Chunk - to account for possible concurrent modifications of blocks in chunks a channel per chunk will be used, many other things also need to be synchronised but some are lesser of importance.
	//Lock &sync.Mutex{}
}

type RegionID struct {
	X int
	Z int
}
