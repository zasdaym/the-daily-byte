package correctcapitalization

import "testing"

func TestCorrectCapitalization(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "single char",
			in:   "a",
			want: true,
		},
		{
			name: "all letter capitalized",
			in:   "USA",
			want: true,
		},
		{
			name: "first letter capitalized",
			in:   "Calvin",
			want: true,
		},
		{
			name: "no capitalization",
			in:   "coding",
			want: true,
		},
		{
			name: "half capitalized",
			in:   "CALVin",
			want: false,
		},
		{
			name: "wrong caps on second letter",
			in:   "cOmputer",
			want: false,
		},
		{
			name: "wrong caps in the middle",
			in:   "compUter",
			want: false,
		},
		{
			name: "wrong caps in the end",
			in:   "computeR",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectCapitalization(tt.in); got != tt.want {
				t.Errorf("IsCorrect(%s) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
