package comparekeystrokes

import "testing"

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{
			name: "empty keystrokes",
			a:    "",
			b:    "",
			want: true,
		},
		{
			name: "similar result",
			a:    "ABC#",
			b:    "CD##AB",
			want: true,
		},
		{
			name: "backspace in beginning",
			a:    "#ABC#",
			b:    "CD##AB",
			want: true,
		},
		{
			name: "different result",
			a:    "cof#dim#ng",
			b:    "code",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.a, tt.b); got != tt.want {
				t.Errorf("Compare(%s, %s) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
