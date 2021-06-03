package main

import "fmt"

func main() {
	user := "Alex"
	fmt.Println(greeting(user))
}

func greeting(name string) string {
	return "hello," + name
}
