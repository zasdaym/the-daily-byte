package removeadjacentduplicates

func RemoveAdjacentDuplicates(s string) string {
	var chars []rune
	for _, c := range s {
		if len(chars) > 0 && c == chars[len(chars)-1] {
			chars = chars[:len(chars)-1]
			continue
		}
		chars = append(chars, c)
	}
	return string(chars)
}
