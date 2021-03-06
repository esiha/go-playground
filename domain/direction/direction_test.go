package direction

import "testing"

func TestDirection_Left90Degrees(t *testing.T) {
	tests := []struct {
		name     string
		input    Direction
		expected Direction
	}{
		{
			"North -> West",
			North,
			West,
		},
		{
			"East -> North",
			East,
			North,
		},
		{
			"South -> East",
			South,
			East,
		},
		{
			"West -> South",
			West,
			South,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.Left90Degrees()
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestDirection_Right90Degrees(t *testing.T) {
	tests := []struct {
		name     string
		input    Direction
		expected Direction
	}{
		{
			"North -> East",
			North,
			East,
		},
		{
			"East -> South",
			East,
			South,
		},
		{
			"South -> West",
			South,
			West,
		},
		{
			"West -> North",
			West,
			North,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.Right90Degrees()
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
