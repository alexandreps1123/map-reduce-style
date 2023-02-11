package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Words struct {
	Word  string
	Value int
}

func Partition(data []byte, n int) []string {

	var newData []string

	ln := strings.Split(string(data), "\n")

	for i := 0; i < len(ln); i = i + n {
		if i+n < len(ln) {
			newData = append(newData, strings.Join(ln[i:i+n], "\n"))
		} else {
			newData = append(newData, strings.Join(ln[i:], "\n"))
		}
	}

	return newData
}

func SplitWords(data string) []Words {
	var result []Words
	var word Words
	for _, w := range removeStopWords(scan(data)) {
		word.Word = w
		word.Value = 1

		result = append(result, word)
	}

	return result
}

func CountWords(mapList chan []Words, sendFinalValue chan map[string]int) {

	final := make(map[string]int)

	for list := range mapList {
		for _, value := range list {
			final[value.Word] += value.Value
		}
	}

	fmt.Println(final)

	sendFinalValue <- final
}

func scan(strData string) []string {
	var result []string

	words := strings.Split(strData, " ")
	pattern := regexp.MustCompile("[^a-zA-Z0-9_]+")

	for _, word := range words {
		word = pattern.ReplaceAllString(word, " ")
		// word = strings.TrimSpace(word)
		// word = strings.Trim(word, "\n")
		// word = strings.ToLower(word)

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

	sw := listStopWord()

	for _, w := range data {
		if !isStopWord(w, sw) {
			words = append(words, w)
		}
	}

	return words
}

func listStopWord() []string {
	bs := readFile("stop_words.txt")
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
