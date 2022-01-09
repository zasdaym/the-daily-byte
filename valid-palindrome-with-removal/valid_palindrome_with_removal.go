package validpalindromewithremoval

func PalindromeWithRemoval(s string) bool {
	if palindrome(s) {
		return true
	}
	for i := 0; i < len(s); i++ {
		if palindrome(s[:i] + s[i+1:]) {
			return true
		}
	}
	return false
}

func palindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
