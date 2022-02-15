package calculatedepth

import "testing"

func TestCalculateDepth(t *testing.T) {

	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "symmetric tree",
			root: &TreeNode{
				Value: 9,
				Left: &TreeNode{
					Value: 1,
				},
				Right: &TreeNode{
					Value: 2,
				},
			},
			want: 2,
		},
		{
			name: "asymmetric tree",
			root: &TreeNode{
				Value: 5,
				Left: &TreeNode{
					Value: 1,
				},
				Right: &TreeNode{
					Value: 29,
					Left: &TreeNode{
						Value: 4,
					},
					Right: &TreeNode{
						Value: 13,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDepth(tt.root); got != tt.want {
				t.Errorf("calculateDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
