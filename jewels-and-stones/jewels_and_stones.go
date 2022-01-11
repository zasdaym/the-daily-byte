package jewelsandstones

func JewelsStones(jewels, stones string) int {
	jewelChars := make(map[rune]bool)
	for _, c := range jewels {
		jewelChars[c] = true
	}

	var result int
	for _, c := range stones {
		if jewelChars[c] {
			result++
		}
	}
	return result
}
