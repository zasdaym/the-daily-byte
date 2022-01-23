package reverselist

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		head *linkedlist.Node
		want *linkedlist.Node
	}{
		{
			name: "single node",
			head: linkedlist.New([]int{1}),
			want: linkedlist.New([]int{1}),
		},
		{
			name: "multiple nodes",
			head: linkedlist.New([]int{1, 2, 3, 4, 5, 6}),
			want: linkedlist.New([]int{6, 5, 4, 3, 2, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.head); !got.Equal(tt.want) {
				t.Errorf("Reverse(%s) = %s, want %s", tt.head, got, tt.want)
			}
		})
	}
}
