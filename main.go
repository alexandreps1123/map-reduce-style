package main

import (
	"sync"
)

var NUMBER_LINES = 200

func main() {

	var wg sync.WaitGroup

	data := Partition(ReadFile("republic.txt"), NUMBER_LINES)

	// Mapper list
	splits := make(chan []Words)

	// Reduce map
	wordsFreqs := make(chan map[string]int)

	for _, aux := range data {
		wg.Add(1)

		go func(data string) {
			defer wg.Done()
			splits <- Map(data)
		}(aux)
	}

	go Reduce(splits, wordsFreqs)

	wg.Wait()
	close(splits)

	SortAndPrint(<-wordsFreqs)
}
