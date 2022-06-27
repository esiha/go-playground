package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/lawn"
	"go-playground/domain/point"
)

type Mower struct {
	position Position
}

func New(position Position) Mower {
	return Mower{position}
}

func (m *Mower) TurnLeft() {
	m.position = m.position.Left90Degrees()
}

func (m *Mower) TurnRight() {
	m.position = m.position.Right90Degrees()
}

func (m *Mower) Advance(l lawn.Lawn) {

	nextPosition := m.position.TranslateOnce()
	if l.Contains(nextPosition.point) {
		m.position = nextPosition
	}

}

func (m *Mower) Position() point.Point {
	return m.position.point
}

func (m *Mower) Direction() direction.Direction {
	return m.position.direction
}
