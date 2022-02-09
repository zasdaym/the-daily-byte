package visiblevalues

import (
	"reflect"
	"testing"
)

func TestVisibleValues(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "symmetric tree",
			root: &TreeNode{
				Value: 4,
				Left:  &TreeNode{Value: 2},
				Right: &TreeNode{Value: 7},
			},
			want: []int{4, 2},
		},
		{
			name: "asymmetric tree",
			root: &TreeNode{
				Value: 7,
				Left: &TreeNode{
					Value: 4,
					Left:  &TreeNode{Value: 1},
					Right: &TreeNode{Value: 4},
				},
				Right: &TreeNode{
					Value: 9,
					Left:  &TreeNode{Value: 8},
					Right: &TreeNode{
						Value: 9,
						Right: &TreeNode{Value: 9},
					},
				},
			},
			want: []int{7, 4, 1, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VisibleValues(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VisibleValues(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
