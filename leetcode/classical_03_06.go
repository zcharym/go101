package leetcode

import (
	"container/list"
)

type AnimalShelf struct {
	dogQueue *list.List
	catQueue *list.List
}

func NewAnimalShelf() AnimalShelf {
	return AnimalShelf{
		dogQueue: list.New(),
		catQueue: list.New(),
	}
}

func (s *AnimalShelf) Enqueue(animal []int) {
	// if cat
	if animal[1] == 0 {
		s.catQueue.PushBack(animal[0])
		// if dog
	} else {
		s.dogQueue.PushBack(animal[0])
	}
}

func (s *AnimalShelf) DequeueAny() []int {
	if s.dogQueue.Len() == 0 && s.catQueue.Len() == 0 {
		return []int{-1, -1}
	}
	if s.dogQueue.Len() == 0 {
		return s.DequeueCat()
	}
	if s.catQueue.Len() == 0 {
		return s.DequeueDog()
	}

	if s.dogQueue.Front().Value.(int) > s.catQueue.Front().Value.(int) {
		return s.DequeueCat()
	} else {
		return s.DequeueDog()
	}

}

func (s *AnimalShelf) DequeueDog() []int {
	if s.dogQueue.Len() > 0 {
		node := s.dogQueue.Front()
		s.dogQueue.Remove(node)
		return []int{node.Value.(int), 1}
	}
	return []int{-1, -1}
}

func (s *AnimalShelf) DequeueCat() []int {
	if s.catQueue.Len() > 0 {
		node := s.catQueue.Front()
		s.catQueue.Remove(node)
		return []int{node.Value.(int), 0}
	}
	return []int{-1, -1}
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
