package main

import (
	"fmt"
	"os"
	"sort"
)

type Words struct {
	Word  string
	Value int
}

var QUANTITY_PRINT = 25

func ReadFile(path string) []byte {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}

func SortAndPrint(mapToOder map[string]int) {
	keys := make([]string, 0, len(mapToOder))

	for key := range mapToOder {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mapToOder[keys[i]] > mapToOder[keys[j]]
	})

	for i := 0; i < QUANTITY_PRINT; i++ {
		fmt.Printf("%v - %v\n", keys[i], mapToOder[keys[i]])
	}
}
