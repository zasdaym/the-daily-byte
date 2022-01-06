package correctcapitalization

import "testing"

func TestCorrectCapitalization(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "a",
			want: true,
		},
		{
			in:   "US",
			want: true,
		},
		{
			in:   "uS",
			want: false,
		},
		{
			in:   "USA",
			want: true,
		},
		{
			in:   "Calvin",
			want: true,
		},
		{
			in:   "compUter",
			want: false,
		},
		{
			in:   "coding",
			want: true,
		},
	}

	for _, tt := range tests {
		if got := CorrectCapitalization(tt.in); got != tt.want {
			t.Errorf("IsCorrect(%s) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
