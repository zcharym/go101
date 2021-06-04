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
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Area()
		expected := 100.0
		if got != expected {
			t.Errorf("got '%.2f' want '%.2f'", got, expected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		expected := math.Pi * 100.0
		if got != expected {
			t.Errorf("got '%.2f' want '%.2f'", got, expected)
		}
	})

}
