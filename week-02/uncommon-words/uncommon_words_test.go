package uncommonwords

import (
	"reflect"
	"testing"
)

func TestUncommonWords(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want []string
	}{
		{
			name: "all words uncommon",
			a:    "the quick",
			b:    "brown fox",
			want: []string{"the", "quick", "brown", "fox"},
		},
		{
			name: "some common words",
			a:    "the tortoise beat the hare",
			b:    "the tortoise lost to the hare",
			want: []string{"beat", "lost", "to"},
		},
		{
			name: "no uncommon words",
			a:    "the daily byte",
			b:    "byte the daily",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UncommonWords(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UncommonWords(%s, %s) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
