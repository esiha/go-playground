package controller

import (
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
	"go-playground/domain/point"
	"testing"
)

func TestController_Run_oneMower(t *testing.T) {
	m := mower.New(mower.NewPosition(point.New(0, 0), direction.North))
	i := []instruction.Instruction{
		instruction.Advance, instruction.TurnRight, instruction.Advance, instruction.TurnLeft,
	}
	l := lawn.Rectangular(point.New(5, 5))

	c := Controller{
		mower:        m,
		instructions: i,
		lawn:         l,
	}

	c.Run()

	expectedPoint := point.New(1, 1)
	expectedDirection := direction.North
	actualPoint := c.mower.Position()
	actualDirection := c.mower.Direction()
	if actualPoint != expectedPoint || actualDirection != expectedDirection {
		t.Fatalf("c'est cassé, expectedPoint=%v, actualPoint=%v, expectedDirection=%v, actualDirection=%v",
			expectedPoint, actualPoint, expectedDirection, actualDirection,
		)
	}
}

func TestController_Run_remainsInLawn(t *testing.T) {
	m := mower.New(mower.NewPosition(point.New(4, 4), direction.North))
	i := []instruction.Instruction{
		instruction.Advance, instruction.Advance, instruction.TurnRight, instruction.Advance, instruction.Advance,
	}
	l := lawn.Rectangular(point.New(5, 5))

	c := Controller{
		mower:        m,
		instructions: i,
		lawn:         l,
	}

	c.Run()

	expectedPoint := point.New(5, 5)
	expectedDirection := direction.East
	actualPoint := c.mower.Position()
	actualDirection := c.mower.Direction()
	if actualPoint != expectedPoint || actualDirection != expectedDirection {
		t.Fatalf("c'est cassé, expectedPoint=%v, actualPoint=%v, expectedDirection=%v, actualDirection=%v",
			expectedPoint, actualPoint, expectedDirection, actualDirection,
		)
	}
}
