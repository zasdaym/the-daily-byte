package reversestring

func Reverse(s string) string {
	chars := []rune(s)
	left, right := 0, len(chars)-1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
	return string(chars)
}
