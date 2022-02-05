package minimumdifference

import "testing"

func TestMinDifference(t *testing.T) {
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
			name: "single node",
			root: &TreeNode{Value: 5},
			want: 5,
		},
		{
			name: "1 depth",
			root: &TreeNode{
				Value: 2,
				Left: &TreeNode{
					Value: 1,
				},
				Right: &TreeNode{
					Value: 3,
				},
			},
			want: 1,
		},
		{
			name: "2 depth",
			root: &TreeNode{
				Value: 29,
				Left: &TreeNode{
					Value: 17,
					Left: &TreeNode{
						Value: 1,
					},
				},
				Right: &TreeNode{
					Value: 50,
					Left: &TreeNode{
						Value: 42,
					},
					Right: &TreeNode{
						Value: 59,
					},
				},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDifference(tt.root); got != tt.want {
				t.Errorf("MinDifference(%v) = %d, want %d", tt.root, got, tt.want)
			}
		})
	}
}
