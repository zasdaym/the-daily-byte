package greaterelements

import (
	"reflect"
	"testing"
)

func TestGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{
			name: "unsorted numbers",
			a:    []int{4, 1, 2},
			b:    []int{1, 3, 4, 2},
			want: []int{-1, 3, -1},
		},
		{
			name: "descending numbers",
			a:    []int{4, 1, 2},
			b:    []int{4, 3, 2, 1},
			want: []int{-1, -1, -1},
		},
		{
			name: "ascending numbers",
			a:    []int{4, 1, 2},
			b:    []int{1, 2, 3, 4},
			want: []int{-1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreaterElements(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterElements(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
