package longestcommonprefix

func LongestCommonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		return words[0]
	}

	var end int
	limit := minLength(words)
	for i := 0; i < limit; i++ {
		for j := 0; j < len(words)-1; j++ {
			if words[j][i] != words[j+1][i] {
				return words[0][:end]
			}
		}
		end++
	}
	return words[0]
}

func minLength(words []string) int {
	result := len(words[0])
	for _, word := range words[1:] {
		if len(word) < result {
			result = len(word)
		}
	}
	return result
}
