package messaging_pattern

import (
	"fmt"
)

type pipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, pipeFunctions ...pipeFunc) <-chan int {
	out := make(<-chan int)
	conv := func(nums []int) <-chan int {
		out := make(chan int)
		go func() {
			for num := range nums {
				out <- num
			}
			close(out)
		}()
		return out
	}

	for idx, _ := range pipeFunctions {
		out = pipeFunctions[idx](conv(nums))
	}
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for num := range in {
			sum += num
		}
		out <- sum
		close(out)
	}()
	return out
}

func echo(in <-chan int) <-chan int {
	go func() {
		for r := range in {
			fmt.Printf("%d\t", r)
		}
	}()
	return in
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
