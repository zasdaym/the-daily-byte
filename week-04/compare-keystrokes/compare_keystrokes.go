package comparekeystrokes

const backspace = '#'

func Compare(a, b string) bool {
	var aChars []rune
	for _, c := range a {
		if c == backspace {
			if len(aChars) > 0 {
				aChars = aChars[:len(aChars)-1]
			}
			continue
		}
		aChars = append(aChars, c)
	}

	var bChars []rune
	for _, c := range b {
		if c == backspace {
			if len(bChars) > 0 {
				bChars = bChars[:len(bChars)-1]
			}
			continue
		}
		bChars = append(bChars, c)
	}
	return string(aChars) == string(bChars)
}
