package controller

import (
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
	"go-playground/domain/point"
	"reflect"
	"testing"
)

func TestController_Run(t *testing.T) {
	c := New(
		[]mower.Mower{
			mower.New(
				mower.NewPosition(point.New(0, 0), direction.North),
				[]instruction.Instruction{
					instruction.Advance, instruction.TurnRight, instruction.Advance, instruction.TurnLeft,
				},
			),
			mower.New(
				mower.NewPosition(point.New(4, 4), direction.North),
				[]instruction.Instruction{
					instruction.Advance, instruction.Advance, instruction.TurnRight, instruction.Advance, instruction.Advance,
				},
			),
		},
		lawn.Rectangular(point.New(5, 5)),
	)

	c.Run()

	expected := []mower.Position{
		mower.NewPosition(point.New(1, 1), direction.North),
		mower.NewPosition(point.New(5, 5), direction.East),
	}

	actual := c.MowersPositions()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
