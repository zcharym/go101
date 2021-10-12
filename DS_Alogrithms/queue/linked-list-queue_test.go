package queue

import (
	"fmt"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.DeQueue()
	q.EnQueue(3)
	fmt.Println(q.Size())
	fmt.Println(q.Values())
}
