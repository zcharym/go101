package main

import (
	"fmt"
)

func main() {
	i := test()
	fmt.Println(i)
}

func test() (i int) {
	defer fmt.Printf("one\n")
	defer fmt.Printf("two\n")
	defer func() {
		i++
	}()
	return
}

// output:
// two
// one
// 1
