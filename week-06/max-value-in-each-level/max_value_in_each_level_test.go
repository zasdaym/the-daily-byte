package maxvalueineachlevel

import (
	"reflect"
	"testing"
)

func TestMaxValuePerLevel(t *testing.T) {
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
			name: "multi-level tree",
			root: &TreeNode{
				Value: 10,
				Left: &TreeNode{
					Value: 2,
				},
				Right: &TreeNode{
					Value: 15,
					Right: &TreeNode{Value: 20},
				},
			},
			want: []int{10, 15, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxValuePerLevel(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxValuePerLevel(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
