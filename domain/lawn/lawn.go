package lawn

import "go-playground/domain/point"

var bottomLeftCorner = point.New(0, 0)

type Lawn struct {
	topRightCorner point.Point
}

func Rectangular(topRightCorner point.Point) Lawn {
	return Lawn{
		topRightCorner: topRightCorner,
	}
}

func (l Lawn) Contains(p point.Point) bool {
	return l.isInBottomLeftCornerOfTopRightCorner(p) &&
		l.isInTopRightCornerOfBottomLeftCorner(p)
}

func (l Lawn) isInBottomLeftCornerOfTopRightCorner(p point.Point) bool {
	return p.X() <= l.topRightCorner.X() && p.Y() <= l.topRightCorner.Y()
}

func (l Lawn) isInTopRightCornerOfBottomLeftCorner(p point.Point) bool {
	return p.X() >= bottomLeftCorner.X() && p.Y() >= bottomLeftCorner.Y()
}
