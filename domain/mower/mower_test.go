package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/point"
	"testing"
)

func TestMower_Run(t *testing.T) {
	l := lawn.Rectangular(point.New(4, 4))

	tests := []struct {
		name             string
		start            Position
		instructions     []instruction.Instruction
		expectedPosition Position
	}{
		{
			"Advance",
			NewPosition(point.New(3, 3), direction.North),
			[]instruction.Instruction{instruction.Advance},
			NewPosition(point.New(3, 4), direction.North),
		},
		{
			"TurnLeft",
			NewPosition(point.New(3, 3), direction.North),
			[]instruction.Instruction{instruction.TurnLeft},
			NewPosition(point.New(3, 3), direction.North.Left90Degrees()),
		},
		{
			"TurnRight",
			NewPosition(point.New(3, 3), direction.North),
			[]instruction.Instruction{instruction.TurnRight},
			NewPosition(point.New(3, 3), direction.North.Right90Degrees()),
		},
		{
			"Remain inside lawn",
			NewPosition(point.New(3, 3), direction.North),
			[]instruction.Instruction{instruction.TurnRight, instruction.Advance, instruction.Advance, instruction.TurnRight},
			NewPosition(point.New(4, 3), direction.North.Right90Degrees().Right90Degrees()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mower := New(tt.start, tt.instructions)
			mower.RunOn(l)
			actual := mower.CurrentPosition()
			if actual != tt.expectedPosition {
				t.Errorf("Expected %+v, got %+v", tt.expectedPosition, actual)
			}
		})
	}
}
