package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/lawn"
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

func (m *Mower) Advance(l lawn.Lawn) {

	nextPosition := m.position.Plus(m.direction.AsVector())
	if l.Contains(nextPosition) {
		m.position = nextPosition
	}

}

func (m *Mower) Position() point.Point {
	return m.position
}

func (m *Mower) Direction() direction.Direction {
	return m.direction
}
