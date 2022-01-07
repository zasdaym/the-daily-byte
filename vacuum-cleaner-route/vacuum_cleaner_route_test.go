package vacuumcleanerroute

import "testing"

func TestIsCompleteRoute(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "horizontal",
			in:   "LR",
			want: true,
		},
		{
			name: "non-complete route",
			in:   "URURD",
			want: false,
		},
		{
			name: "all direction",
			in:   "RUULLDRD", want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCompleteRoute(tt.in); got != tt.want {
				t.Errorf("IsCompleteRoute(%s) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
