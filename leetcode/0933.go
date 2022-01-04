package leetcode

type RecentCounter struct {
	arr []int
}

func NewRecentCounter() RecentCounter {
	return RecentCounter{
		arr: []int{},
	}
}

func (c *RecentCounter) Ping(t int) int {
	c.arr = append(c.arr, t)
	for c.arr[0] < t-3000 {
		c.arr = c.arr[1:len(c.arr)]
	}
	return len(c.arr)
}
