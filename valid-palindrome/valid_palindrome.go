package validpalindrome

import (
	"strings"
	"unicode"
)

func Palindrome(s string) bool {
	normalized := normalize(s)
	left, right := 0, len(normalized)-1
	for left < right {
		if normalized[left] != normalized[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func normalize(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		sb.WriteRune(unicode.ToLower(c))
	}
	return sb.String()
}
