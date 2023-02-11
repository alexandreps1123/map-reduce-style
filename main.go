package main

import (
	"fmt"
	"sync"
)

func main() {

	data := Partition(readFile("republic2.txt"), 200)

	// Mapper list
	splits := make(chan []Words)

	// Reduce List
	wordsFreqs := make(chan map[string]int)

	var wg sync.WaitGroup

	wg.Add(len(data))

	for _, aux := range data {
		go func(data string) {
			defer wg.Done()
			splits <- Map(data)
		}(aux)
	}

	go Reduce(splits, wordsFreqs)

	wg.Wait()
	close(splits)

	fmt.Println(<-wordsFreqs)
}
