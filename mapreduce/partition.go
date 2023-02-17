package mapreduce

import "strings"

func Partition(fileContent []byte, n int) []string {

	var pData []string

	ln := strings.Split(string(fileContent), "\n")

	for i := 0; i < len(ln); i = i + n {
		pData = append(pData, contentToAppend(i, n, ln))
	}

	return pData
}

func contentToAppend(i int, n int, ln []string) string {
	if i+n < len(ln) {
		return strings.Join(ln[i:i+n], "\n")
	} else {
		return strings.Join(ln[i:], "\n")
	}
}
