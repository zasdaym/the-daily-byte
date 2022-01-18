package mergelinkedlists

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		a, b *linkedlist.Node
		want *linkedlist.Node
	}{
		{
			name: "first list smaller than second list",
			a:    &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3}}},
			b:    &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 6}}},
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 6}}}}}},
		},
		{
			name: "zigzag merge with same length",
			a:    &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 5}}},
			b:    &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 6}}},
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 6}}}}}},
		},
		{
			name: "first list longer than second list",
			a:    &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 7}}}},
			b:    &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 6}}},
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 6, Next: &linkedlist.Node{Value: 7}}}}}}},
		},
		{
			name: "second list longer than first list",
			a:    &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 5}}},
			b:    &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 6, Next: &linkedlist.Node{Value: 7}}}},
			want: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3, Next: &linkedlist.Node{Value: 4, Next: &linkedlist.Node{Value: 5, Next: &linkedlist.Node{Value: 6, Next: &linkedlist.Node{Value: 7}}}}}}},
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
