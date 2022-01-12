package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "single word",
			in:   "Cat",
			want: "taC",
		},
		{
			name: "single sentence",
			in:   "The Daily Byte",
			want: "etyB yliaD ehT",
		},
		{
			name: "palindrome",
			in:   "civic",
			want: "civic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Errorf("Reverse(%s) = %s, want %s", tt.in, got, tt.want)
			}
		})
	}
}
