package queue

import "github.com/zchary-ma/impl-ds/common"

type LinkedListQueue struct {
	front, rear *common.Node
}

func NewLinkedListQueue() *LinkedListQueue {
	q := &LinkedListQueue{}
	q.rear = &common.Node{}
	q.front = q.rear
	return q
}

func (l *LinkedListQueue) EnQueue(value common.ElementType) {
	node := &common.Node{Val: value, Next: nil}
	l.front.Next = node
	l.rear = node
}

func (l *LinkedListQueue) DeQueue() (common.ElementType, bool) {
	if l.front == l.rear {
		return -1, false
	}
	node := l.front.Next
	l.front.Next = node.Next
	if l.rear == node {
		l.rear = l.front
	}
	return node.Val, true
}

func (l *LinkedListQueue) Front() common.ElementType {
	panic("implement me")
}

func (l *LinkedListQueue) Full() bool {
	panic("implement me")
}

func (l *LinkedListQueue) Empty() bool {
	if l.front == l.rear {
		return true
	}
	return false
}

func (l *LinkedListQueue) Size() int {
	var count int
	for p := l.front; p != nil; p = p.Next {
		count++
	}
	return count
}

func (l *LinkedListQueue) Clear() {
	panic("implement me")
}

func (l *LinkedListQueue) Values() []common.ElementType {
	var temp []common.ElementType
	for p := l.front; p != nil; p = p.Next {
		temp = append(temp, p.Val)
	}
	return temp
}
