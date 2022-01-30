package callcounter

import (
	"testing"
)

func TestCallCounter(t *testing.T) {
	counter := New()
	tests := []struct {
		timestamp int
		want      int
	}{
		{
			timestamp: 1,
			want:      1,
		},
		{
			timestamp: 300,
			want:      2,
		},
		{
			timestamp: 3000,
			want:      3,
		},
		{
			timestamp: 3002,
			want:      3,
		},
		{
			timestamp: 7000,
			want:      1,
		},
	}

	for _, tt := range tests {
		if got := counter.Ping(tt.timestamp); got != tt.want {
			t.Errorf("counter.Ping(%d) = %d, want %d", tt.timestamp, got, tt.want)
		}
	}
}
