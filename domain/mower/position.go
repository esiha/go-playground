package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/point"
)

type Position struct {
	point     point.Point
	direction direction.Direction
}

func NewPosition(p point.Point, d direction.Direction) Position {
	return Position{p, d}
}

func (pos *Position) Left90Degrees() Position {
	return Position{pos.point, pos.direction.Left90Degrees()}
}

func (pos *Position) Right90Degrees() Position {
	return Position{pos.point, pos.direction.Right90Degrees()}
}

func (pos *Position) TranslateOnce() Position {
	return Position{pos.point.Plus(pos.direction.AsVector()), pos.direction}
}
