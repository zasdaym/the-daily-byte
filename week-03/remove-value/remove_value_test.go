package removevalue

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		head *linkedlist.Node
		n    int
		want *linkedlist.Node
	}{
		{
			name: "value exists",
			head: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3}}},
			n:    3,
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: nil}},
		},
		{
			name: "value doesn't exists",
			head: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3}}},
			n:    6,
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.head, tt.n); !got.Equal(tt.want) {
				t.Errorf("Remove(%s, %d) = %s, want %s", tt.head, tt.n, got, tt.want)
			}
		})
	}
}
