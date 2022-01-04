package validpalindrome

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "level",
			want: true,
		},
		{
			in:   "algorithm",
			want: false,
		},
		{
			in:   "A man, a plan, a canal: Panama.",
			want: true,
		},
	}

	for _, tt := range tests {
		if got := Palindrome(tt.in); got != tt.want {
			t.Errorf("Palindrome(%s) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
