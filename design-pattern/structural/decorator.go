package main

import (
	"log"
	"os"
	"time"
)

// type fiboFunc func() func() int

// TODO import from module
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

type sumFunc func([]int) int

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func logger(fn sumFunc, logger *log.Logger) sumFunc {
	return func(nums []int) int {
		log.Printf("time:%v", time.Now())
		return fn(nums)
	}
}

func main() {
	nums := []int{1, 2, 3}
	f := logger(sum, log.New(os.Stdout, "test", 1))

	f(nums)
}
