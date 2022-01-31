package findvalue

import "testing"

func TestFindValue(t *testing.T) {
	tests := []struct {
		name    string
		root    *Node
		value   int
		wantNil bool
	}{
		{
			name: "Value exists",
			root: &Node{
				Value: 10,
				Left: &Node{
					Value: 5,
					Left: &Node{
						Value: 3,
					},
					Right: &Node{
						Value: 7,
					},
				},
				Right: &Node{
					Value: 15,
					Left: &Node{
						Value: 12,
					},
					Right: &Node{
						Value: 20,
					},
				},
			},
			value:   7,
			wantNil: false,
		},
		{
			name: "Value does not exist",
			root: &Node{
				Value: 10,
				Left: &Node{
					Value: 5,
					Left: &Node{
						Value: 3,
					},
					Right: &Node{
						Value: 7,
					},
				},
				Right: &Node{
					Value: 15,
					Left: &Node{
						Value: 12,
					},
					Right: &Node{
						Value: 20,
					},
				},
			},
			value:   18,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindValue(tt.root, tt.value)
			if got == nil && !tt.wantNil {
				t.Fatalf("FindValue(%v, %d) = nil, want not nil", tt.root, tt.value)
			}
			if got != nil && tt.wantNil {
				t.Fatalf("FindValue(%v, %d) = not nil, want nil", tt.root, tt.value)
			}
			if got != nil && got.Value != tt.value {
				t.Fatalf("FindValue(%v, %d) = %v, want %v", tt.root, tt.value, got.Value, tt.value)
			}
		})
	}
}
