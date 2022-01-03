package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "Cat",
			want: "taC",
		},
		{
			in:   "The Daily Byte",
			want: "etyB yliaD ehT",
		},
		{
			in:   "civic",
			want: "civic",
		},
	}
	for _, tt := range tests {
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%s) = %s, want %s", tt.in, got, tt.want)
		}
	}
}
