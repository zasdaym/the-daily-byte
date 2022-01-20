package removenthtolastnode

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestRemoveNthLastNode(t *testing.T) {
	tests := []struct {
		head *linkedlist.Node
		n    int
		want *linkedlist.Node
	}{
		{
			head: linkedlist.New([]int{1, 2, 3, 4, 5}),
			n:    1,
			want: linkedlist.New([]int{1, 2, 3, 4}),
		},
		{
			head: linkedlist.New([]int{1, 2, 3, 4, 5}),
			n:    4,
			want: linkedlist.New([]int{1, 3, 4, 5}),
		},
		{
			head: linkedlist.New([]int{1, 2, 3, 4, 5}),
			n:    5,
			want: linkedlist.New([]int{2, 3, 4, 5}),
		},
		{
			head: linkedlist.New([]int{1, 2, 3, 4, 5}),
			n:    6,
			want: linkedlist.New([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, tt := range tests {
		if got := RemoveNthLastNode(tt.head, tt.n); !got.Equal(tt.want) {
			t.Errorf("RemoveNthLastNode(%s, %d) = %s, want %s", tt.head, tt.n, got, tt.want)
		}
	}
}
