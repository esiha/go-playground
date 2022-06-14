package direction

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
