package leetcode

type TripleInOne struct {
	stack       []int
	firstStart  int
	firstEnd    int
	secondStart int
	secondEnd   int
	thirdStart  int
	thirdEnd    int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stack:       make([]int, stackSize*3),
		firstStart:  0,
		firstEnd:    0,
		secondStart: stackSize,
		secondEnd:   stackSize,
		thirdStart:  stackSize * 2,
		thirdEnd:    stackSize * 2,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {

	switch stackNum {
	case 0:
		if this.firstEnd < len(this.stack)/3 {
			this.stack[this.firstEnd] = value
			this.firstEnd++
		}
	case 1:
		if this.secondEnd < len(this.stack)*2/3 {
			this.stack[this.secondEnd] = value
			this.secondEnd++
		}
	case 2:
		if this.thirdEnd < len(this.stack) {
			this.stack[this.thirdEnd] = value
			this.thirdEnd++
		}

	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	switch stackNum {
	case 0:
		if this.firstEnd > this.firstStart {
			val := this.stack[this.firstEnd-1]
			this.stack[this.firstEnd-1] = 0
			this.firstEnd--
			return val
		}
	case 1:
		if this.secondEnd > this.secondStart {
			val := this.stack[this.secondEnd-1]
			this.stack[this.secondEnd-1] = 0
			this.secondEnd--
			return val
		}
	case 2:
		if this.thirdEnd > this.thirdStart {
			val := this.stack[this.thirdEnd-1]
			this.stack[this.thirdEnd-1] = 0
			this.thirdEnd--
			return val
		}
	}
	return -1

}

func (this *TripleInOne) Peek(stackNum int) int {

	switch stackNum {
	case 0:
		if this.firstEnd > this.firstStart {
			val := this.stack[this.firstEnd-1]
			return val
		}
	case 1:
		if this.secondEnd > this.secondStart {
			val := this.stack[this.secondEnd-1]
			return val
		}
	case 2:
		if this.thirdEnd > this.thirdStart {
			val := this.stack[this.thirdEnd-1]
			return val
		}
	}
	return -1

}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	switch stackNum {
	case 0:
		if this.firstStart == this.firstEnd {
			return true
		}
	case 1:
		if this.secondStart == this.secondEnd {
			return true
		}
	case 2:
		if this.thirdStart == this.thirdEnd {
			return true
		}
	}
	return false

}
