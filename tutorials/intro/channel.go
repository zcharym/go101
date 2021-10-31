// Channels are the pipes that connect concurrent goroutines.
package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)      // unbuffered
	msgQueue := make(chan string, 10) // buffered chan

	msgQueue <- "1"
	msgQueue <- "2"
	msgQueue <- "3"

	go func() {
		time.Sleep(time.Second * 1)
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
	fmt.Println(<-msgQueue)
	fmt.Println(<-msgQueue)
	fmt.Println(<-msgQueue)
}
