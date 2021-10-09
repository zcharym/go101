package leetcode

import (
	"fmt"
	"testing"
)

func TestMyQueue1(t *testing.T) {
	obj := NewMyQueue()
	obj.Push(1)
	obj.Push(2)
	b := obj.Peek()
	a := obj.Pop()
	c := obj.Empty()
	fmt.Println(a, b, c)
}

func TestMyQueue2(t *testing.T) {
	obj := NewMyQueue()
	obj.Push(1)
	a := obj.Pop()
	c := obj.Empty()
	fmt.Println(a, c)
}

func TestMyQueue3(t *testing.T) {
	obj := NewMyQueue()
	obj.Push(1)
	obj.Push(2)
	a := obj.Pop()
	c := obj.Peek()
	fmt.Println(a, c)
}
