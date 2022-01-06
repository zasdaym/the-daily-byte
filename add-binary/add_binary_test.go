package addbinary

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		want   string
	}{
		{
			name:   "different length without carry",
			first:  "1",
			second: "100",
			want:   "101",
		},
		{
			name:   "different length with carry",
			first:  "11",
			second: "1",
			want:   "100",
		},
		{
			name:   "same length",
			first:  "1",
			second: "0",
			want:   "1",
		},
		{
			name:   "long binaries with carry",
			first:  "10101010",
			second: "11001100",
			want:   "101110110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.first, tt.second); got != tt.want {
				t.Errorf("Sum(%s, %s) = %s, want %s", tt.first, tt.second, got, tt.want)
			}
		})
	}
}
