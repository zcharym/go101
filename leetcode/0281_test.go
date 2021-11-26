package leetcode

import (
	"testing"
)

func TestNewZigzagIterator(t *testing.T) {
	z := NewZigzagIterator([]int{1, 2, 3}, []int{4, 5, 6})
	ch := make(chan int, 20)
	for z.hasNext() {
		ch <- z.next()
	}
	for {
		select {
		case v := <-ch:
			t.Log(v)
		default:
			return
		}
	}
}
