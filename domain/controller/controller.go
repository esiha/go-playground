package controller

import (
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
)

type Controller struct {
	mower        mower.Mower
	instructions []instruction.Instruction
	lawn         lawn.Lawn
}

func (c *Controller) Run() {
	for _, i := range c.instructions {
		switch i {
		case instruction.Advance:
			c.mower.Advance(c.lawn)
		case instruction.TurnRight:
			c.mower.TurnRight()
		case instruction.TurnLeft:
			c.mower.TurnLeft()
		default:
			panic("Not handled")
		}
	}
}
