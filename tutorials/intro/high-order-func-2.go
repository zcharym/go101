package main

import (
	"errors"
	"fmt"
)

// 另一类为将其他函数作为返回结果的函数
type calculatorFunc func(int, int) (int, error)
type operation func(int, int) int

func main() {
	x, y := 3, 4
	op := func(x, y int) int {
		return x + y
	}
	add := genCalculator(op)
	result, _ := add(x, y)
	fmt.Println(result)
}

func genCalculator(op operation) calculatorFunc {

	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("op is empty")
		}
		return op(x, y), nil
	}
}
