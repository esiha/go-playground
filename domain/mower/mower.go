package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/point"
)

type Mower struct {
	position  point.Point
	direction direction.Direction
}

func New(p point.Point, d direction.Direction) Mower {
	return Mower{
		p,
		d,
	}
}

func (m *Mower) TurnLeft() {
	m.direction = m.direction.Left90Degrees()
}

func (m *Mower) TurnRight() {
	m.direction = m.direction.Right90Degrees()
}

func (m *Mower) Advance() {
	m.position = m.position.Plus(m.direction.AsVector())
}
