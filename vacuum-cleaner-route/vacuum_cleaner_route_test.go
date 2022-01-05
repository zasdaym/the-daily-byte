package vacuumcleanerroute

import "testing"

func TestIsCompleteRoute(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in: "LR", want: true,
		},
		{
			in: "URURD", want: false,
		},
		{
			in: "RUULLDRD", want: true,
		},
	}

	for _, tt := range tests {
		if got := IsCompleteRoute(tt.in); got != tt.want {
			t.Errorf("IsCompleteRoute(%s) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
