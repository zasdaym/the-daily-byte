package removeadjacentduplicates

import "testing"

func TestRemoveAdjacentDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no adjacent duplicate",
			input: "abcdefg",
			want:  "abcdefg",
		},
		{
			name:  "adjacent duplicates",
			input: "abccbefggfe",
			want:  "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAdjacentDuplicates(tt.input); got != tt.want {
				t.Errorf("RemoveAdjacentDuplicates(%s) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}
