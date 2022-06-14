package point

import "testing"

func TestPoint_Plus(t *testing.T) {
	a := New(3, 4)
	b := New(2, 8)

	actual := a.Plus(b)
	expected := New(5, 12)

	if actual != expected {
		t.Errorf("actual='%v', expected='%v'", actual, expected)
	}
}
