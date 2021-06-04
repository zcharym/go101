package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	p := new(int)
	q := new(int)
	fmt.Println(p == q)
	fmt.Println(p, q)

	var x, y int
	fmt.Println(x == y)
	fmt.Println(x, y)
}
