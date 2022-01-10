package twosum

func TwoSum(nums []int, target int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[target-num] {
			return true
		}
		seen[num] = true
	}
	return false
}
