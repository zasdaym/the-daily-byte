package intersectionofnumbers

func Intersection(a, b []int) []int {
	nums := make(map[int]bool)
	for _, v := range a {
		nums[v] = true
	}
	result := make([]int, 0)
	for _, v := range b {
		if nums[v] {
			result = append(result, v)
		}
	}
	return result
}
