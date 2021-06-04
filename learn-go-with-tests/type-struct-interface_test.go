package learning_tests

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	expected := 40.0
	if got != expected {
		t.Errorf("got '%.2f' want '%.2f'", got, expected)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got '%.2f' expected '%.2f'", got, expected)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		checkArea(t, rect, 100.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, math.Pi*100.0)
	})

}
