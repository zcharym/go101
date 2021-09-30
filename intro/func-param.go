package main

import (
	"fmt"
)

// 关于函参是否被修改的问题
// 值类型：复制为副本，对原值没有影响
// 引用类型：浅拷贝，拷贝原值自身（如指针，cap，len）而不会拷贝底层引用的数据

func main() {
	complexArr := [3][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	fmt.Println("complexArr", complexArr) // [[a b c] [d e f] []]

	fmt.Println("opWithSlice", opWithSlice(complexArr)) // [[a b c] [a e f] []]

	fmt.Println("complexArr", complexArr) //  [[a b c] [a e f] []]

	fmt.Println("opWithArr", opWithArr(complexArr)) // [[a b c] [a e f] [g h i]]

	fmt.Println("complexArr", complexArr) // [[a b c] [a e f] []]
}

func opWithSlice(arr [3][]string) [3][]string {
	arr[1][0] = "a"
	return arr
}

func opWithArr(arr [3][]string) [3][]string {
	arr[2] = []string{"g", "h", "i"}
	return arr
}
