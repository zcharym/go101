package main

import (
	"fmt"
	"time"
)

func main() {
	channels := make([]chan string, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan string) // NOTE
		go process(channels[i])
	}

	for idx, channel := range channels {
		<-channel
		fmt.Println("routine", idx)
	}

}

func process(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "finished"
}
