package convertbinarysearchtree

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *linkedlist.Node
	}{
		{
			name: "one level",
			root: &TreeNode{
				Value: 5,
				Left: &TreeNode{
					Value: 1,
				},
				Right: &TreeNode{
					Value: 6,
				},
			},
			want: linkedlist.New([]int{1, 5, 6}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.root); !got.Equal(tt.want) {
				t.Errorf("Convert(%v) = %s, want %s", tt.root, got, tt.want)
			}
		})
	}
}
