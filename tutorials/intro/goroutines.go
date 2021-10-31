package main

import (
	"fmt"
	"time"
)

func echo(s string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s\n", s)
		time.Sleep(time.Second)
	}
}

func main() {
	echo("calling")

	// invoke echo in goroutine
	go echo("[go] calling")

	// can also start a goroutine for an anonymous function
	go func() {
		fmt.Println("go func calling")
	}()

	// these two function call are running in separate goroutines now

	time.Sleep(time.Second * 10)

	fmt.Println("done.")
}
