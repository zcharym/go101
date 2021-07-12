package main

import (
	"fmt"
)

type Turtle struct {
	x, y float32
}

func main() {
	t := Turtle{x: 0, y: 0}
	t.moveY(2)
	t.moveY(2)
	t.moveY(2)
	t.moveY(2)
	t.moveX(3)

	fmt.Printf("%v\n", t.position())

}

func (t *Turtle) moveX(offset float32) {
	t.x += 1.0 * offset
}

func (t *Turtle) moveY(offset float32) {
	t.y += 1.0 * offset
}

func (t *Turtle) position() []float32 {
	return []float32{t.x, t.y}
}
