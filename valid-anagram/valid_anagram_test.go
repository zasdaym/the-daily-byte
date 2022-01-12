package validanagram

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{
			name: "not anagram",
			a:    "program",
			b:    "function",
			want: false,
		},
		{
			name: "anagram",
			a:    "listen",
			b:    "silent",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidAnagram(tt.a, tt.b); got != tt.want {
				t.Errorf("ValidAnagram(%s, %s) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
