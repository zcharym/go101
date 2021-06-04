package learning_tests

import (
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
	rect := Rectangle{10.0, 10.0}
	got := Area(rect)
	expected := 100.0
	if got != expected {
		t.Errorf("got '%.2f' want '%.2f'", got, expected)
	}
}
