package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGophers(i, c) // goroutine orders execute randomly
	}
	time.Sleep(time.Second * 4) // -> use chan to receive stop signal

	timeout := time.After(time.Second * 1)
	for j := 0; j < 5; j++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher", gopherId, "finished sleeping.")
		case <-timeout:
			fmt.Println("timeout")
			return
		}

	}
}

func sleepyGophers(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)+1000) * time.Millisecond)
	c <- id
}
