package sortedarraytobinarysearchtree

import "testing"

func TestNewBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{
			name: "balanced tree",
			nums: []int{1, 2, 3},
			want: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBST(tt.nums); !sameTree(got, tt.want) {
				t.Errorf("NewBST(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func sameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}
