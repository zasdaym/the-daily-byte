package firstuniquecharacter

import "testing"

func TestFirstUniqueCharacter(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "single unique char",
			s:    "abcdabd",
			want: 2,
		},
		{
			name: "multiple unique chars",
			s:    "developer",
			want: 0,
		},
		{
			name: "no unique char",
			s:    "abbbcac",
			want: -1,
		},
	}

	for _, tt := range tests {
		if got := FirstUniqueCharacter(tt.s); got != tt.want {
			t.Run(tt.name, func(t *testing.T) {
				t.Errorf("FirstUniqueCharacter(%s) = %d, want %d", tt.s, got, tt.want)
			})
		}
	}
}
