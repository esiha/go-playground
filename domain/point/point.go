package point

type Point struct {
	x int
	y int
}

func New(x int, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Plus(other Point) Point {
	return New(p.x+other.x, p.y+other.y)
}

func (p Point) X() int {
	return p.x
}
func (p Point) Y() int {
	return p.y
}
