package main

import (
	"errors"
	"fmt"
)

func foo(arg int) (int, error) {
	if arg < 10 {
		return -1, errors.New("input number is less than 10.")
	}
	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("[%d - %s]", e.arg, e.prob)
}

func bar(arg int) (int, error) {
	if arg < 10 {
		return -1, &argError{arg, "can't work with number less than 10"}
	}
	return arg, nil
}

func main() {
	for _, i := range []int{8, 12} {
		if r, e := foo(i); e != nil {
			fmt.Println("foo fails", e, "arg:", r)
		} else {
			fmt.Println("foo succeeds", e, "arg:", r)
		}
	}

	for _, j := range []int{8, 12} {
		if r, e := bar(j); e != nil {
			fmt.Println("foo fails", e, "arg:", r)
		} else {
			fmt.Println("foo succeeds", e, "arg:", r)
		}
	}
}
