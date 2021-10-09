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

func (this *StackOfPlates) Push(val int) {
	//fmt.Println("PUSH", val)
	if this.stackSize == 0 {
		return
	}
	if this.idx == len(this.stacks) {
		this.stacks = append(this.stacks, &Stack{
			nums: make([]int, this.stackSize),
			pop:  0,
		})
	}
	curStack := this.stacks[this.idx]
	curStack.nums[curStack.pop] = val
	curStack.pop++
	if curStack.pop == this.stackSize {
		this.idx++
	}
}

func (this *StackOfPlates) Pop() int {
	//fmt.Println("POP")
	if this.stackSize == 0 {
		return -1
	}
	i := this.idx
	for ; i >= 0; i-- {
		if i == len(this.stacks) {
			continue
		}
		if this.stacks[i].pop != 0 {
			break
		}
	}
	if i < 0 {
		this.idx = 0
		this.stacks = this.stacks[:0]
		return -1
	}
	targetStack := this.stacks[i]
	ret := targetStack.nums[targetStack.pop-1]
	targetStack.pop--
	if targetStack.pop == 0 {
		this.stacks = this.stacks[:i]
		if i > 0 && this.stacks[i-1].pop < this.stackSize {
			this.idx = i - 1
		} else {
			this.idx = i
		}
	}
	return ret
}

func (this *StackOfPlates) PopAt(index int) int {
	//fmt.Println("POP AT", index)
	if this.stackSize == 0 {
		return -1
	}
	n := len(this.stacks)
	if index < 0 || index >= n {
		return -1
	}
	curStack := this.stacks[index]
	ret := curStack.nums[curStack.pop-1]
	curStack.pop--
	if curStack.pop == 0 {
		if index < this.idx {
			this.idx--
			this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
		} else {
			this.stacks = this.stacks[:index]
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
