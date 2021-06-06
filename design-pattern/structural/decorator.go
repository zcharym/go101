package go101_pattern

import (
	"log"
)

// type fiboFunc func() func() int
type sumFunc func([]int) int

// TODO import from module
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func logger(fn sumFunc, logger *log.Logger) sumFunc {
	return func(nums []int) int {
		return fn(nums)
	}
}
