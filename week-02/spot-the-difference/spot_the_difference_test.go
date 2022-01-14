package spotthedifference

import "testing"

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want rune
	}{
		{
			name: "no difference",
			a:    "coding",
			b:    "ingcod",
			want: ' ',
		},
		{
			name: "with difference",
			a:    "foobar",
			b:    "barfoot",
			want: 't',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.a, tt.b); got != tt.want {
				t.Errorf("Difference(%s, %s) = '%c', want %c", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
