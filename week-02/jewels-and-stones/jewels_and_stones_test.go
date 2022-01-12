package jewelsandstones

import "testing"

func TestJewelsStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{
			name:   "stone jewel exists",
			jewels: "abc",
			stones: "ab",
			want:   2,
		},
		{
			name:   "stone jewel doesn't exist",
			jewels: "AYOPD",
			stones: "ayopd",
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JewelsStones(tt.jewels, tt.stones); got != tt.want {
				t.Errorf("JewelsStones(%s, %s) = %d, want %d", tt.jewels, tt.stones, got, tt.want)
			}
		})
	}
}
