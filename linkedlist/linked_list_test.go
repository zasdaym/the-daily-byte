package linkedlist_test

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestNodeString(t *testing.T) {
	tests := []struct {
		node *linkedlist.Node
		want string
	}{
		{
			node: &linkedlist.Node{Value: 1, Next: &linkedlist.Node{Value: 2, Next: &linkedlist.Node{Value: 3}}},
			want: "->1->2->3",
		},
	}

	for i, tt := range tests {
		if got := tt.node.String(); got != tt.want {
			t.Errorf("#%d = %s, want %s", i, got, tt.want)
		}
	}
}

func TestNodeEqual(t *testing.T) {
	tests := []struct {
		name string
		a, b *linkedlist.Node
		want bool
	}{
		{
			name: "similar lists",
			a:    &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, nil}}},
			b:    &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, nil}}},
			want: true,
		},
		{
			name: "same length with different value",
			a:    &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, nil}}},
			b:    &linkedlist.Node{1, &linkedlist.Node{4, &linkedlist.Node{5, nil}}},
			want: false,
		},
		{
			name: "different length",
			a:    &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, &linkedlist.Node{4, nil}}}},
			b:    &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, nil}}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equal(tt.b); got != tt.want {
				t.Errorf("%s == %s is %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		values []int
		want   *linkedlist.Node
	}{
		{
			values: []int{1, 2, 3},
			want:   &linkedlist.Node{1, &linkedlist.Node{2, &linkedlist.Node{3, nil}}},
		},
	}

	for _, tt := range tests {
		if got := linkedlist.New(tt.values); !got.Equal(tt.want) {
			t.Errorf("New(%v) = %s, want %s", tt.values, got, tt.want)
		}
	}
}
