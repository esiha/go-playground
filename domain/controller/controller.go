package controller

import (
	"go-playground/domain/lawn"
	"go-playground/domain/mower"
)

type Controller struct {
	mowers []mower.Mower
	lawn   lawn.Lawn
}

func New(mowers []mower.Mower, l lawn.Lawn) Controller {
	return Controller{mowers, l}
}

func (c *Controller) Run() {
	for i := range c.mowers {
		c.mowers[i].RunOn(c.lawn)
	}
}

func (c *Controller) MowersPositions() []mower.Position {
	positions := make([]mower.Position, len(c.mowers))
	for i, m := range c.mowers {
		positions[i] = m.CurrentPosition()
	}
	return positions
}
