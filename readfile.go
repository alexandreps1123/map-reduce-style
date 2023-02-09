package main

import (
	"os"
)

func readFile(path string) []byte {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
