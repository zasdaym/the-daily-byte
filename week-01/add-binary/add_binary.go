package addbinary

import "fmt"

func Sum(first, second string) string {
	if len(second) > len(first) {
		first, second = second, first
	}
	var result string
	var carry int
	for i := 0; i < len(first); i++ {
		var a, b int
		if first[len(first)-1-i] == '1' {
			a = 1
		}
		if i < len(second) && second[len(second)-1-i] == '1' {
			b = 1
		}
		sum := carry + a + b
		carry = sum / 2
		result = fmt.Sprintf("%d%s", sum%2, result)
	}
	if carry == 1 {
		result = fmt.Sprintf("%d%s", carry, result)
	}
	return result
}
