package utils

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/map-reduce-style/common"
)

func ReadFile(path string) []byte {

	file, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if file.Size() == 0 {
		log.Fatal("Empty file")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func SortAndPrint(mapToOder map[string]int) {
	keys := make([]string, 0, len(mapToOder))
	if len(mapToOder) == 0 {
		log.Fatal("No output to print")
	}

	for key := range mapToOder {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mapToOder[keys[i]] > mapToOder[keys[j]]
	})

	for i := 0; i < common.QUANTITY_PRINT; i++ {
		fmt.Printf("%v - %v\n", keys[i], mapToOder[keys[i]])
	}
}
