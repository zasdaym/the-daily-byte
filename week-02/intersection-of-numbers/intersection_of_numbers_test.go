package intersectionofnumbers

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{
			name: "with intersection",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{2, 3, 7, 8},
			want: []int{2, 3},
		},
		{
			name: "no intersection",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{7, 12, 15},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection(%v, %v) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
