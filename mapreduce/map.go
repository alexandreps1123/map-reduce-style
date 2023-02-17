package mapreduce

import (
	"regexp"
	"strings"

	"github.com/map-reduce-style/common"
	"github.com/map-reduce-style/utils"
)

func Map(data string) []common.Words {
	return splitWords(data)
}

func splitWords(data string) []common.Words {
	var result []common.Words
	var word common.Words

	for _, w := range removeStopWords(scan(data)) {
		word.Word, word.Value = w, 1

		result = append(result, word)
	}

	return result
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

	sw := stopWordsSlice()

	for _, w := range data {
		if !isStopWord(w, sw) {
			words = append(words, w)
		}
	}

	return words
}

func stopWordsSlice() []string {
	bs := utils.ReadFile("stop_words.txt")
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
