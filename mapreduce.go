package main

func Map(data string) []Words {
	return SplitWords(data)
}

func Reduce(mapList chan []Words, sendFinalValue chan map[string]int) {
	CountWords(mapList, sendFinalValue)
}
