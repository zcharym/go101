package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var (
		count    int
		finished int
	)
	for i := 0; i < 10; i++ {
		go func() {
			if requestVote() {
				count++
			}
			finished++
		}()
	}

	for count < 5 || finished != 10 {
		// waite until the vote is finished or get 5 counts
	}

	if count < 5 {
		fmt.Println("candidate loses", count, finished)
	} else {
		fmt.Println("candidate wons", count, finished)
	}

}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Intn(11)%2 == 0
}
