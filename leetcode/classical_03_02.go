package leetcode

type MinStack struct {
	min  []int
	data []int
	top  int
}

func NewMinStack() MinStack {
	return MinStack{
		min:  make([]int, 10001),
		data: make([]int, 10001),
		top:  -1,
	}
}

func (s *MinStack) Push(x int) {
	if s.top == -1 {
		s.top++
		s.min[s.top] = x
	} else {
		s.top++
		if s.min[s.top-1] > x {
			s.min[s.top] = x
		} else {
			s.min[s.top] = s.min[s.top-1]
		}
	}
	s.data[s.top] = x
}

func (s *MinStack) Pop() {
	s.top--
}

func (s *MinStack) Top() int {
	return s.data[s.top]
}

func (s *MinStack) GetMin() int {
	return s.min[s.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
