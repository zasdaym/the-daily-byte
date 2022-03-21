package mergelinkedlists

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		a, b *linkedlist.Node[int]
		want *linkedlist.Node[int]
	}{
		{
			name: "first list smaller than second list",
			a:    linkedlist.New([]int{1, 2, 3}),
			b:    linkedlist.New([]int{4, 5, 6}),
			want: linkedlist.New([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name: "zigzag merge with same length",
			a:    linkedlist.New([]int{1, 3, 5}),
			b:    linkedlist.New([]int{2, 4, 6}),
			want: linkedlist.New([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name: "first list longer than second list",
			a:    linkedlist.New([]int{1, 3, 5, 7}),
			b:    linkedlist.New([]int{2, 4, 6}),
			want: linkedlist.New([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			name: "second list longer than first list",
			a:    linkedlist.New([]int{1, 3, 5}),
			b:    linkedlist.New([]int{2, 4, 6, 7}),
			want: linkedlist.New([]int{1, 2, 3, 4, 5, 6, 7}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.a, tt.b); !got.Equal(tt.want) {
				t.Errorf("Merge(%s, %s) = %s, want %s", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
