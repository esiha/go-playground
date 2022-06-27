package mower

import (
	"go-playground/domain/direction"
	"go-playground/domain/point"
	"testing"
)

func TestPosition_Left90Degrees(t *testing.T) {
	p := point.New(3, 2)
	d := direction.North
	pos := NewPosition(p, d)

	expected := NewPosition(p, d.Left90Degrees())

	if actual := pos.Left90Degrees(); actual != expected {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestPosition_Right90Degrees(t *testing.T) {
	p := point.New(3, 2)
	d := direction.East
	pos := NewPosition(p, d)

	expected := NewPosition(p, d.Right90Degrees())

	if actual := pos.Right90Degrees(); actual != expected {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestPosition_TranslateOnce(t *testing.T) {
	p := point.New(3, 2)
	d := direction.South
	pos := NewPosition(p, d)

	expected := NewPosition(p.Plus(d.AsVector()), d)

	if actual := pos.TranslateOnce(); actual != expected {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
