package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "empty words",
			words: nil,
			want:  "",
		},
		{
			name:  "single word",
			words: []string{"Pneumonoultramicroscopicsilicovolcanoconiosis"},
			want:  "Pneumonoultramicroscopicsilicovolcanoconiosis",
		},
		{
			name: "partial common prefix",
			words: []string{
				"colorado",
				"color",
				"cold",
			},
			want: "col",
		},
		{
			name: "duplicate words",
			words: []string{
				"color",
				"color",
				"color",
			},
			want: "color",
		},
		{
			name: "no common prefix",
			words: []string{
				"a",
				"b",
				"c",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.words); got != tt.want {
				t.Errorf("LongestCommonPrefix(%v) = %s, want %s", tt.words, got, tt.want)
			}
		})
	}
}
