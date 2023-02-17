package mapreduce

import "github.com/map-reduce-style/common"

func Reduce(mapedList chan []common.Words, sendFinalValue chan map[string]int) {
	countWords(mapedList, sendFinalValue)
}

func countWords(mapedList chan []common.Words, sendFinalValue chan map[string]int) {

	final := make(map[string]int)

	for list := range mapedList {
		for _, value := range list {
			final[value.Word] += value.Value
		}
	}

	sendFinalValue <- final
}
