package gathernarytreelevels

import (
	"reflect"
	"testing"
)

func Test_gatherNaryTreeLevels(t *testing.T) {
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
			name: "ternary tree",
			root: &TreeNode{
				Value: 8,
				Children: []*TreeNode{
					{Value: 2},
					{Value: 3},
					{Value: 9},
				},
			},
			want: [][]int{{8}, {2, 3, 9}},
		},
		{
			name: "multi-level ternary tree",
			root: &TreeNode{
				Value: 2,
				Children: []*TreeNode{
					{
						Value: 1,
						Children: []*TreeNode{
							{Value: 8},
						},
					},
					{
						Value: 6,
						Children: []*TreeNode{
							{
								Value: 2,
								Children: []*TreeNode{
									{Value: 19},
									{Value: 12},
									{Value: 90},
								}},
						},
					},
					{
						Value: 9,
						Children: []*TreeNode{
							{Value: 2},
						},
					},
				},
			},
			want: [][]int{{2}, {1, 6, 9}, {8, 2, 2}, {19, 12, 90}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GatherNaryTreeLevels(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GatherNaryTreeLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
