package leetcode

// 该类用两个栈来实现一个队列
type MyQueue struct {
	inStack  Stack
	outStack Stack
}

/** Initialize your data structure here. */
func NewMyQueue() MyQueue {
	return MyQueue{
		inStack: Stack{
			nums: []int{},
			pop:  -1,
		},
		outStack: Stack{
			nums: []int{},
			pop:  -1,
		},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack.nums = append(q.inStack.nums, x)
	q.inStack.pop++
}

func (q *MyQueue) Pop() int {
	// 入栈为空时
	if q.Peek() == -1 {
		return -1
	}
	// 出栈为空时，将入栈0-pop的元素倒序压入
	if q.outStack.pop == -1 {
		nums := q.inStack.nums[0 : q.inStack.pop+1]
		for i := len(nums) - 1; i >= 0; i-- {
			q.outStack.nums = append(q.outStack.nums, nums[i])
			q.outStack.pop++
		}
		q.inStack.nums = []int{}
		q.inStack.pop = -1
	}
	num := q.outStack.nums[q.outStack.pop]
	q.outStack.nums = q.outStack.nums[0:q.outStack.pop]
	q.outStack.pop--
	return num
}

func (q *MyQueue) Peek() int {
	if q.outStack.pop != -1 {
		return q.outStack.nums[q.outStack.pop]
	}
	if q.inStack.pop != -1 {
		return q.inStack.nums[0]
	}
	return -1
}

func (q *MyQueue) Empty() bool {
	return q.inStack.pop == -1 && q.outStack.pop == -1
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
