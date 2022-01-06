package correctcapitalization

import "unicode"

func CorrectCapitalization(s string) bool {
	if len(s) < 2 {
		return true
	}
	firstIsUpper := unicode.IsUpper(rune(s[0]))
	previousIsUpper := unicode.IsUpper(rune(s[1]))
	if !firstIsUpper && previousIsUpper {
		return false
	}
	for _, c := range s[2:] {
		if previousIsUpper != unicode.IsUpper(c) {
			return false
		}
	}
	return true
}
