package validpalindromewithremoval

import "testing"

func TestPalindromeWithRemoval(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "non-palindrome",
			s:    "abccab",
			want: false,
		},
		{
			name: "palindrome without removal",
			s:    "abcba",
			want: true,
		},
		{
			name: "palindrome with removal",
			s:    "foobof",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromeWithRemoval(tt.s); got != tt.want {
				t.Errorf("IsValidPalindromeWithRemoval(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
