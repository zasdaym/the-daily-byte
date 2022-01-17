package mergelinkedlists

import (
	"testing"
)

func TestNodeString(t *testing.T) {
	tests := []struct {
		node *Node
		want string
	}{
		{
			node: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}},
			want: "->1->2->3",
		},
	}

	for _, tt := range tests {
		if got := tt.node.String(); got != tt.want {
			t.Errorf("%s, want %s", got, tt.want)
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		a, b *Node
		want *Node
	}{
		{
			name: "first list smaller than second list",
			a:    &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}},
			b:    &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}},
			want: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}}}},
		},
		{
			name: "zigzag merge with same length",
			a:    &Node{Value: 1, Next: &Node{Value: 3, Next: &Node{Value: 5}}},
			b:    &Node{Value: 2, Next: &Node{Value: 4, Next: &Node{Value: 6}}},
			want: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}}}},
		},
		{
			name: "first list longer than second list",
			a:    &Node{Value: 1, Next: &Node{Value: 3, Next: &Node{Value: 5, Next: &Node{Value: 7}}}},
			b:    &Node{Value: 2, Next: &Node{Value: 4, Next: &Node{Value: 6}}},
			want: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 7}}}}}}},
		},
		{
			name: "second list longer than first list",
			a:    &Node{Value: 1, Next: &Node{Value: 3, Next: &Node{Value: 5}}},
			b:    &Node{Value: 2, Next: &Node{Value: 4, Next: &Node{Value: 6, Next: &Node{Value: 7}}}},
			want: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 7}}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.a, tt.b); !IsEqual(got, tt.want) {
				t.Errorf("Merge(%s, %s) = %s, want %s", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func IsEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return IsEqual(a.Next, b.Next)
}
