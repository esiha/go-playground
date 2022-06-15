package direction

import "go-playground/domain/point"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Left90Degrees() Direction {
	switch d {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	default:
		panic("unhandled direction")
	}
}

func (d Direction) Right90Degrees() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic("unhandled direction")
	}
}

func (d Direction) AsVector() point.Point {
	switch d {
	case North:
		return point.New(0, 1)
	case East:
		return point.New(1, 0)
	case West:
		return point.New(-1, 0)
	case South:
		return point.New(0, -1)
	default:
		panic("unhandled direction")
	}
}
