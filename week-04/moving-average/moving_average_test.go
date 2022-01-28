package movingaverage

import (
	"testing"
)

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 3,
		},
		{
			n:    5,
			want: 4,
		},
		{
			n:    7,
			want: 5,
		},
		{
			n:    6,
			want: 6,
		},
	}

	ma := New(3)
	for _, tt := range tests {
		if got := ma.Next(tt.n); got != tt.want {
			t.Errorf("ma.Next(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
