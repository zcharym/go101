package leetcode

type Stack struct {
	nums []int
	pop  int
}

type StackOfPlates struct {
	// 除了最后一个栈，其余的都应该是满的
	stacks []*Stack
	// 每个栈的容量上限
	stackSize int
	// 永远指向最后一个未满的栈
	idx int
}

func NewStackOfPlates(cap int) StackOfPlates {
	sp := StackOfPlates{
		stacks:    make([]*Stack, 0),
		stackSize: cap,
		idx:       0,
	}
	return sp
}

func (s *StackOfPlates) Push(val int) {
	if s.stackSize == 0 {
		return
	}
	if s.idx == len(s.stacks) {
		s.stacks = append(s.stacks, &Stack{
			nums: make([]int, s.stackSize),
			pop:  0,
		})
	}
	curStack := s.stacks[s.idx]
	curStack.nums[curStack.pop] = val
	curStack.pop++
	if curStack.pop == s.stackSize {
		s.idx++
	}
}

func (s *StackOfPlates) Pop() int {
	//fmt.Println("POP")
	if s.stackSize == 0 {
		return -1
	}
	i := s.idx
	for ; i >= 0; i-- {
		if i == len(s.stacks) {
			continue
		}
		if s.stacks[i].pop != 0 {
			break
		}
	}
	if i < 0 {
		s.idx = 0
		s.stacks = s.stacks[:0]
		return -1
	}
	targetStack := s.stacks[i]
	ret := targetStack.nums[targetStack.pop-1]
	targetStack.pop--
	if targetStack.pop == 0 {
		s.stacks = s.stacks[:i]
		if i > 0 && s.stacks[i-1].pop < s.stackSize {
			s.idx = i - 1
		} else {
			s.idx = i
		}
	}
	return ret
}

func (s *StackOfPlates) PopAt(index int) int {
	if s.stackSize == 0 {
		return -1
	}
	n := len(s.stacks)
	if index < 0 || index >= n {
		return -1
	}
	curStack := s.stacks[index]
	ret := curStack.nums[curStack.pop-1]
	curStack.pop--
	if curStack.pop == 0 {
		if index < s.idx {
			s.idx--
			s.stacks = append(s.stacks[:index], s.stacks[index+1:]...)
		} else {
			s.stacks = s.stacks[:index]
		}
	}
	return ret
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
