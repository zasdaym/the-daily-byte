package spotthedifference

func Difference(a, b string) rune {
	if len(a) == len(b) {
		return ' '
	}
	count := make(map[rune]int)
	count[rune(b[len(b)-1])]++
	for i := range a {
		count[rune(a[i])]--
		count[rune(b[i])]++
	}
	var result rune
	for k, v := range count {
		if v > 0 {
			result = k
		}
	}
	return result
}
