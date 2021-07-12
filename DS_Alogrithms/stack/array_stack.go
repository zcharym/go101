package stack

import (
	"go101/DS_Alogrithms/list"
)

type ArrayStack struct {
	list *list.SqList
}

func New() *ArrayStack {
	return &ArrayStack{list: list.New()}
}

func (s *ArrayStack) Push(elem interface{}) {
	s.list.Add(elem)
}

func (s *ArrayStack) Pop() (value interface{}, ok bool) {
	value, ok = s.list.Get(s.list.Size() - 1)
	s.list.Remove(s.list.Size() - 1)
	return
}

func (s *ArrayStack) Peek() (value interface{}, ok bool) {
	return s.list.Get(s.list.Size() - 1)
}

func (s *ArrayStack) Empty() bool {
	return s.list.Empty()
}

func (s *ArrayStack) Size() int {
	return s.list.Size()
}

func (s *ArrayStack) Clear() {
	s.list.Clear()
}

func (s *ArrayStack) Values() []interface{} {
	panic("implement me")
}
