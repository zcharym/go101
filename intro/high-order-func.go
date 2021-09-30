package main

import (
	"errors"
	"fmt"
)

// 关于高阶函数
// example：便携calculate函数来实现两个数整数间的加减乘除，并且整数和运算操作
// 都是由该函数的调用方给出

type operator func(int, int) int

func main() {
	x, y := 12, 5

	fmt.Println(Calculate(x, y, Add))
	fmt.Println(Calculate(x, y, Minus))
	fmt.Println(Calculate(x, y, Mutiply))
	fmt.Println(Calculate(x, y, Divide))
}

// Calculate 一类为接收其他的函数作为参数传入的函数
func Calculate(x, y int, op operator) (int, error) {
	if op == nil {
		return 0, errors.New("operator is empty")
	}
	return op(x, y), nil
}

func Divide(x, y int) int {
	return x / y
}

func Mutiply(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}

func Minus(x, y int) int {
	return x - y
}
