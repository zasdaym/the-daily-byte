package findthemode

import "testing"

func TestFindMode(t *testing.T) {
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
			root: &TreeNode{
				Value: 1,
			},
			want: 1,
		},
		{
			name: "multi node",
			root: &TreeNode{
				Value: 7,
				Left: &TreeNode{
					Value: 4,
					Left: &TreeNode{
						Value: 1,
					},
					Right: &TreeNode{
						Value: 4,
					},
				},
				Right: &TreeNode{
					Value: 9,
					Left: &TreeNode{
						Value: 8,
					},
					Right: &TreeNode{
						Value: 9,
						Right: &TreeNode{
							Value: 9,
						},
					},
				},
			},
			want: 9,
		},
	}

	for _, tt := range tests {
		if got := FindMode(tt.root); got != tt.want {
			t.Errorf("FindMode(%v) = %d, want %d", tt.root, got, tt.want)
		}
	}
}
