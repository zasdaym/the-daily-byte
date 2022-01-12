package validanagram

import "reflect"

func ValidAnagram(a, b string) bool {
	return reflect.DeepEqual(countChars(a), countChars(b))
}

func countChars(s string) map[rune]int {
	result := make(map[rune]int)
	for _, c := range s {
		result[c]++
	}
	return result
}
