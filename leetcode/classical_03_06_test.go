package leetcode

import (
	"fmt"
	"testing"
)

func TestAnimalShelf(t *testing.T) {
	obj := NewAnimalShelf()
	obj.Enqueue([]int{0, 0})
	obj.Enqueue([]int{1, 0})
	a := obj.DequeueAny()
	b := obj.DequeueDog()
	c := obj.DequeueCat()
	fmt.Println(a, b, c)
}
