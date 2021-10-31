package main

// reference: https://stackoverflow.com/questions/25491370/fibonacci-closure-in-go
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}
