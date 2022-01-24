package validatecharacters

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func Validate(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var unclosed []rune
	for _, c := range s {
		pair, closer := pairs[c]
		if closer && len(unclosed) == 0 {
			return false
		}
		if !closer {
			unclosed = append(unclosed, c)
			continue
		}
		if pair != unclosed[len(unclosed)-1] {
			return false
		}
		unclosed = unclosed[:len(unclosed)-1]
	}
	return len(unclosed) == 0
}

// ({[]})
