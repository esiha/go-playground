package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/instruction"
	"go-playground/domain/lawn"
	"go-playground/domain/point"
)

type Mower struct {
	position     Position
	instructions []instruction.Instruction
}

func New(position Position, instructions []instruction.Instruction) Mower {
	return Mower{position, instructions}
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

func (m *Mower) RunOn(l lawn.Lawn) {
	for _, i := range m.instructions {
		nextPosition := m.nextPosition(i)
		if l.Contains(nextPosition.point) {
			m.position = nextPosition
		}
	}
}

func (m *Mower) nextPosition(i instruction.Instruction) Position {
	switch i {
	case instruction.Advance:
		return m.position.TranslateOnce()
	case instruction.TurnRight:
		return m.position.Right90Degrees()
	case instruction.TurnLeft:
		return m.position.Left90Degrees()
	default:
		panic("Not handled")
	}
}
