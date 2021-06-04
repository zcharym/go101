package learning_tests

import (
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
