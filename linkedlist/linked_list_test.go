package linkedlist_test

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestNodeString(t *testing.T) {
	tests := []struct {
		node *linkedlist.Node[int]
		want string
	}{
		{
			node: &linkedlist.Node[int]{Value: 1, Next: &linkedlist.Node[int]{Value: 2, Next: &linkedlist.Node[int]{Value: 3}}},
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
		a, b *linkedlist.Node[int]
		want bool
	}{
		{
			name: "similar lists",
			a:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, nil}}},
			b:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, nil}}},
			want: true,
		},
		{
			name: "same length with different value",
			a:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, nil}}},
			b:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{4, &linkedlist.Node[int]{5, nil}}},
			want: false,
		},
		{
			name: "different length",
			a:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, &linkedlist.Node[int]{4, nil}}}},
			b:    &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, nil}}},
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
		want   *linkedlist.Node[int]
	}{
		{
			values: []int{1, 2, 3},
			want:   &linkedlist.Node[int]{1, &linkedlist.Node[int]{2, &linkedlist.Node[int]{3, nil}}},
		},
	}

	for _, tt := range tests {
		if got := linkedlist.New(tt.values); !got.Equal(tt.want) {
			t.Errorf("New(%v) = %s, want %s", tt.values, got, tt.want)
		}
	}
}
