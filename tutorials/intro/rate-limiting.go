package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Second)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}
