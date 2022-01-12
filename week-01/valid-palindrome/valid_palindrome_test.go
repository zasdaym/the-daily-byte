package validpalindrome

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "lowercase palindrome",
			in:   "level",
			want: true,
		},
		{
			name: "not palindrome",
			in:   "algorithm",
			want: false,
		},
		{
			name: "palindrome with non-letter chars",
			in:   "A man, a plan, a canal: Panama.",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome(tt.in); got != tt.want {
				t.Errorf("Palindrome(%s) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
