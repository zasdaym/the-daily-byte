package validatecharacters

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "valid without nested characters",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "valid nested characters",
			s:    "([()]{})",
			want: true,
		},
		{
			name: "closer without opener",
			s:    "))]]}}",
			want: false,
		},
		{
			name: "invalid pair",
			s:    "([]}",
			want: false,
		},
		{
			name: "odd length string",
			s:    "[]]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.s); got != tt.want {
				t.Errorf("Validate(%s) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
