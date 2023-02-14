package main

import (
	"regexp"
	"strings"
)

func Partition(fileContent []byte, n int) []string {

	var pData []string

	ln := strings.Split(string(fileContent), "\n")

	for i := 0; i < len(ln); i = i + n {
		pData = append(pData, stringToAppend(i, n, ln))
	}

	return pData
}

func Map(data string) []Words {
	return splitWords(data)
}

func Reduce(mapedList chan []Words, sendFinalValue chan map[string]int) {
	countWords(mapedList, sendFinalValue)
}

func splitWords(data string) []Words {
	var result []Words
	var word Words

	for _, w := range removeStopWords(scan(data)) {
		word.Word, word.Value = w, 1

		result = append(result, word)
	}

	return result
}

func countWords(mapedList chan []Words, sendFinalValue chan map[string]int) {

	final := make(map[string]int)

	for list := range mapedList {
		for _, value := range list {
			final[value.Word] += value.Value
		}
	}

	sendFinalValue <- final
}

func scan(strData string) []string {
	var result []string

	words := strings.Split(strData, " ")
	pattern := regexp.MustCompile("[^a-zA-Z0-9_]+")

	for _, word := range words {
		word = pattern.ReplaceAllString(word, " ")

		for _, aux := range strings.Split(word, " ") {
			aux = strings.TrimSpace(aux)
			aux = strings.Trim(aux, "\n")
			aux = strings.ToLower(aux)

			if aux != "" {
				result = append(result, aux)
			}
		}
	}

	return result
}

func removeStopWords(data []string) []string {
	var words []string

	sw := stopWordsList()

	for _, w := range data {
		if !isStopWord(w, sw) {
			words = append(words, w)
		}
	}

	return words
}

func stopWordsList() []string {
	bs := ReadFile("stop_words.txt")
	sw := strings.Split(string(bs), ",")
	ascii_lowercase := strings.Split("abcdefghijklmnopqrstuvwxyz", "")

	for _, a := range ascii_lowercase {
		sw = append(sw, a)
	}

	return sw
}

func isStopWord(w string, sw []string) bool {
	for _, s := range sw {
		if w == strings.ToLower(s) {
			return true
		}
	}

	return false
}

func stringToAppend(i int, n int, ln []string) string {
	if i+n < len(ln) {
		return strings.Join(ln[i:i+n], "\n")
	} else {
		return strings.Join(ln[i:], "\n")
	}
}
