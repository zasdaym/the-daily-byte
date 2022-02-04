package identicaltrees

import "testing"

func TestIdentical(t *testing.T) {
	tests := []struct {
		name string
		a, b *TreeNode
		want bool
	}{
		{
			name: "identical",
			a: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 3},
			},
			b: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 3},
			},
			want: true,
		},
		{
			name: "different value",
			a: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 3},
			},
			b: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 6},
			},
			want: false,
		},
		{
			name: "different structure",
			a: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{Value: 6},
			},
			b: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 1},
				Right: &TreeNode{
					Value: 6,
					Left:  &TreeNode{Value: 4},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Identical(tt.a, tt.b); got != tt.want {
				t.Errorf("Identical(%v, %v) = %v, want %v", tt.a, tt.a, got, tt.want)
			}
		})
	}
}
