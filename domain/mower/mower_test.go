package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/point"
	"testing"
)

func TestMower_TurnLeft(t *testing.T) {
	north := direction.North
	mower := New(NewPosition(aPoint(), north), []instruction.Instruction{})

	mower.TurnLeft()

	expected := north.Left90Degrees()
	if mower.position.direction != expected {
		t.Fatalf("Expected mower direction to be %v, got %v", expected, mower.position.direction)
	}
}

func aPoint() point.Point {
	return point.New(3, 2)
}

func TestMower_TurnRight(t *testing.T) {
	west := direction.West
	mower := New(NewPosition(aPoint(), west), []instruction.Instruction{})

	mower.TurnRight()

	expected := west.Right90Degrees()
	if mower.position.direction != expected {
		t.Fatalf("Expected mower direction to be %v, got %v", expected, mower.position.direction)
	}
}

func TestMower_Advance(t *testing.T) {
	p := point.New(3, 3)

	tests := []struct {
		name      string
		direction direction.Direction
		expected  point.Point
	}{
		{
			"North",
			direction.North,
			point.New(3, 4),
		},
		{
			"East",
			direction.East,
			point.New(4, 3),
		},
		{
			"West",
			direction.West,
			point.New(2, 3),
		},
		{
			"South",
			direction.South,
			point.New(3, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mower := New(NewPosition(p, tt.direction), []instruction.Instruction{})
			mower.Advance(lawn.Rectangular(point.New(5, 5)))
			actual := mower.position.point
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

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
			actual := mower.position
			if actual != tt.expectedPosition {
				t.Errorf("Expected %+v, got %+v", tt.expectedPosition, actual)
			}
		})
	}
}
