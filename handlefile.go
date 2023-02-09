package main

import (
	"fmt"
	"regexp"
	"strings"
)

func partition(data []byte, n int) []string {

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
	bs := readFile("stop_words.txt")

	sw := strings.Split(string(bs), ",")

	ascii_lowercase := strings.Split("abcdefghijklmnopqrstuvwxyz", "")

	for _, a := range ascii_lowercase {
		sw = append(sw, a)
	}

	for _, d := range data {
		var flag bool = true

		for _, s := range sw {
			if d == s {
				flag = false
			}
		}

		if flag {
			words = append(words, d)
		}
	}

	return words
}

func splitWords(data string) {

	var result [][2]string

	fmt.Println(len(result))

	for i, w := range removeStopWords(scan(data)) {
		result[i][0] = w
		result[i][1] = "1"
	}

	fmt.Println(len(result))
}
