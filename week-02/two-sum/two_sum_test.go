package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "two sum doesn't exist",
			nums:   []int{1, 9, 5, 3, 2},
			target: 19,
			want:   false,
		},
		{
			name:   "two sum exists",
			nums:   []int{1, 3, 8, 2},
			target: 10,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.nums, tt.target); got != tt.want {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
