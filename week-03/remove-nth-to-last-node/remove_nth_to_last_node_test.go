package removenthtolastnode

import (
	"testing"

	ll "github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestRemoveNthLastNode(t *testing.T) {
	tests := []struct {
		head *ll.Node
		n    int
		want *ll.Node
	}{
		{
			head: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}}},
			n:    1,
			want: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, nil}}}},
		},
		{
			head: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}}},
			n:    4,
			want: &ll.Node{1, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}},
		},
		{
			head: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}}},
			n:    5,
			want: &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}},
		},
		{
			head: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}}},
			n:    6,
			want: &ll.Node{1, &ll.Node{2, &ll.Node{3, &ll.Node{4, &ll.Node{5, nil}}}}},
		},
	}

	for _, tt := range tests {
		if got := RemoveNthLastNode(tt.head, tt.n); !got.Equal(tt.want) {
			t.Errorf("RemoveNthLastNode(%s, %d) = %s, want %s", tt.head, tt.n, got, tt.want)
		}
	}
}
