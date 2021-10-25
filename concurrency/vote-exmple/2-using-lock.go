package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var (
		count    int
		finished int
		mu       sync.Mutex
	)
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if requestVote() {
				count++
			}
			finished++
		}()
	}

	// bad for burning up CPU
	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()
		time.Sleep(time.Millisecond * 50) // reduce CPU use but still not right a time.Sleep
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
