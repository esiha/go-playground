package lawn

import (
	"go-playground/domain/point"
	"testing"
)

func TestLawn_Contains(t *testing.T) {
	topRightCorner := point.New(5, 5)
	lawn := Lawn{topRightCorner: topRightCorner}

	tests := []struct {
		name     string
		input    point.Point
		expected bool
	}{
		{
			name:     "topRightCorner",
			input:    topRightCorner,
			expected: true,
		},
		{
			name:     "x too big",
			input:    point.New(6, 5),
			expected: false,
		},
		{
			name:     "x small enough",
			input:    point.New(4, 5),
			expected: true,
		},
		{
			name:     "y too big",
			input:    point.New(5, 6),
			expected: false,
		},
		{
			name:     "y small enough",
			input:    point.New(5, 4),
			expected: true,
		},
		{
			name:     "x right in the spot",
			input:    point.New(0, 4),
			expected: true,
		},
		{
			name:     "x too small",
			input:    point.New(-1, 4),
			expected: false,
		},
		{
			name:     "y right in the spot",
			input:    point.New(4, 0),
			expected: true,
		},
		{
			name:     "y too small",
			input:    point.New(4, -1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lawn.Contains(tt.input)
			if actual != tt.expected {
				t.Fatalf("Contains(%v) = %v, expected = %v", tt.input, actual, tt.expected)
			}
		})
	}
}
