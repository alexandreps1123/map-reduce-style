package mapreduce

import "github.com/map-reduce-style/common"

func Reduce(mapedSlice chan []common.Words, sendFinalValue chan map[string]int) {
	countWords(mapedSlice, sendFinalValue)
}

func countWords(mapedSlice chan []common.Words, sendFinalValue chan map[string]int) {

	final := make(map[string]int)

	for slice := range mapedSlice {
		for _, value := range slice {
			final[value.Word] += value.Value
		}
	}

	sendFinalValue <- final
}
