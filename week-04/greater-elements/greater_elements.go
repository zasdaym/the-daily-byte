package greaterelements

func GreaterElements(subset, nums []int) []int {
	greaters := make(map[int]int)
	var stack []int
	for _, num := range nums {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			greaters[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	var result []int
	for _, v := range subset {
		greater, ok := greaters[v]
		if !ok {
			greater = -1
		}
		result = append(result, greater)
	}
	return result
}
