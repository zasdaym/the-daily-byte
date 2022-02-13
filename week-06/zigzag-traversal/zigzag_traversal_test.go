package zigzagtraversal

import (
	"reflect"
	"testing"
)

func TestZigzagTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "non-empty tree",
			root: &TreeNode{
				Value: 8,
				Left: &TreeNode{
					Value: 2,
				},
				Right: &TreeNode{
					Value: 29,
					Left: &TreeNode{
						Value: 3,
					},
					Right: &TreeNode{
						Value: 35,
					},
				},
			},
			want: [][]int{{8}, {29, 2}, {3, 35}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZigzagTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZigzagTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
