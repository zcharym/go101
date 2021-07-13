package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGophers(i, c) // goroutine orders execute randomly
	}
	time.Sleep(time.Second * 4) // -> use chan to receive stop signal

	for j := 0; j < 5; j++ {
		gopherId := <-c
		fmt.Println("gopher", gopherId, "finished sleeping.")
	}
}

func sleepyGophers(id int, c chan int) {
	time.Sleep(time.Second * 3)
	fmt.Printf("...[%d]snore...\n", id)
	c <- id
}
