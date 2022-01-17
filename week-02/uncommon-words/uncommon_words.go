package uncommonwords

import (
	"sort"
	"strings"
)

func UncommonWords(first, second string) []string {
	firstWords := make(map[string]bool)
	for _, word := range strings.Fields(first) {
		firstWords[word] = true
	}
	secondWords := make(map[string]bool)
	for _, word := range strings.Fields(second) {
		secondWords[word] = true
	}

	var result []string
	for word := range firstWords {
		if !secondWords[word] {
			result = append(result, word)
		}
	}
	for word := range secondWords {
		if !firstWords[word] {
			result = append(result, word)
		}
	}
	sort.Strings(result)
	return result
}
