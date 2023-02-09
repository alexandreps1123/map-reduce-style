package main

import "time"

func main() {

	// splits := make(map[string]int)

	// divide o texto em pedacos menores de 200 linhas
	data := partition(readFile("republic2.txt"), 200)

	for _, aux := range data {
		splitWords(aux)
	}

	time.Sleep(time.Second * 5)
}
