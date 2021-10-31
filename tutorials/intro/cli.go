package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			s += " " + arg
		}
		fmt.Println(s)
	} else {
		fmt.Println("no args take.")
	}
}
