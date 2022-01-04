package leetcode

type ZigzagIterator struct {
	count int
	data  chan int
}

func NewZigzagIterator(v1, v2 []int) *ZigzagIterator {
	l1, l2 := len(v1), len(v2)
	if l1 == 0 && l2 == 0 {
		return nil
	}

	var zig = &ZigzagIterator{}
	zig.count = 2
	zig.data = make(chan int, l1+l2)

	for idx1, idx2 := 0, 0; idx1 < l1 || idx2 < l2; {
		if idx1 < l1 {
			zig.data <- v1[idx1]
			idx1++
		}
		if idx2 < l2 {
			zig.data <- v2[idx2]
			idx2++
		}
	}

	return zig
}

func (z *ZigzagIterator) next() int {
	if z.hasNext() {
		res := <-z.data
		return res
	}
	return -1
}

func (z *ZigzagIterator) hasNext() bool {
	if len(z.data) == 0 {
		return false
	}
	return true
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */
