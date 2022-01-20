package findmiddleelement

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestMiddle(t *testing.T) {
	tests := []struct {
		name string
		head *linkedlist.Node
		want int
	}{
		{
			name: "odd number elements",
			head: linkedlist.New([]int{1, 2, 3}),
			want: 2,
		},
		{
			name: "even number elements",
			head: linkedlist.New([]int{1, 2, 3, 4}),
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Middle(tt.head); got != tt.want {
				t.Errorf("Middle(%s) = %d, want %d", tt.head, got, tt.want)
			}
		})
	}
}
