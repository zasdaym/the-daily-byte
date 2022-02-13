package bottomsup

import (
	"reflect"
	"testing"
)

func TestBottomsUp(t *testing.T) {
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
			name: "symmetric tree",
			root: &TreeNode{
				Value: 5,
				Left: &TreeNode{
					Value: 2,
				},
				Right: &TreeNode{
					Value: 10,
				},
			},
			want: [][]int{{2, 10}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BottomsUp(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BottomsUp(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
