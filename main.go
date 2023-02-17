package main

import (
	"os"
	"sync"

	"github.com/map-reduce-style/common"
	"github.com/map-reduce-style/mapreduce"
	"github.com/map-reduce-style/utils"
)

func main() {

	var wg sync.WaitGroup

	data := mapreduce.Partition(utils.ReadFile(os.Args[1]), common.NUMBER_LINES)

	// Mapper list
	splits := make(chan []common.Words)

	// Reduce map
	wordsFreqs := make(chan map[string]int)

	for _, aux := range data {
		wg.Add(1)

		go func(data string) {
			defer wg.Done()
			splits <- mapreduce.Map(data)
		}(aux)
	}

	go mapreduce.Reduce(splits, wordsFreqs)

	wg.Wait()
	close(splits)

	utils.SortAndPrint(<-wordsFreqs)
}
