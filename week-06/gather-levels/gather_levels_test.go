package gatherlevels

import (
	"reflect"
	"testing"
)

func TestGatherLevels(t *testing.T) {
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
			name: "simple tree",
			root: &TreeNode{
				Value: 4,
				Left:  &TreeNode{Value: 2},
				Right: &TreeNode{Value: 7},
			},
			want: [][]int{{4}, {2, 7}},
		},
		{
			name: "asymmetric tree",
			root: &TreeNode{
				Value: 2,
				Left:  &TreeNode{Value: 10},
				Right: &TreeNode{Value: 15, Right: &TreeNode{Value: 20}},
			},
			want: [][]int{{2}, {10, 15}, {20}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GatherLevels(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GatherLevels(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}
