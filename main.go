package main

import (
	"fmt"
	rand2 "math/rand"
)

func main() {

	// initiate array(1)
	var a [5]int
	fmt.Println("initial array:", a)
	a[4] = 100
	fmt.Println("set by anchor", a)
	fmt.Println("length of array", len(a))

	// initiate array(2)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// initiate array(3)
	d := make([]string, 4)
	d[0] = "hello"
	d[1] = "go"
	d[2] = "programming"
	d[3] = "language"

	fmt.Println("part of d", d[2:4])

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d array:", c)

	// range
	for i, num := range d {
		fmt.Println(i, num)
	}

	sum := sum(1, 4, -1)
	fmt.Println("sum is", sum)

	index := intSeq()
	fmt.Println(index())
	fmt.Println(index())
	fmt.Println(index())
	fmt.Println(index())

	fmt.Println(factorial(5))

}

func getRandomNumber() int {
	return rand2.Intn(5)
}

// Variadic Functions
func sum(sums ...int) int {
	fmt.Println("numbers are", sums)
	total := 0
	for _, num := range sums {
		total += num
	}
	return total
}

// anonymous function with closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i

	}
}

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
