package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGophers() // goroutine orders execute randomly
	}
	time.Sleep(time.Second * 4) // -> use chan to receive stop signal
}

func sleepyGophers() {
	time.Sleep(time.Second * 3)
	fmt.Println("...snore...")
}
