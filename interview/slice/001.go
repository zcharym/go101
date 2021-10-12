package main

import (
	"fmt"
)

func main() {
	var array [10]int
	var slice = array[5:6]
	fmt.Println("slice len:", len(slice)) // 1
	fmt.Println("slice cap:", cap(slice)) // 5 （为总长-切取的起始长度）
}
