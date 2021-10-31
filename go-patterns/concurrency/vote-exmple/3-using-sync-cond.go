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
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if requestVote() {
				count++
			}
			finished++
			cond.Broadcast()
		}()
	}
	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait()
	}

	if count < 5 {
		fmt.Println("candidate loses", count, finished)
	} else {
		fmt.Println("candidate wons", count, finished)
	}
	mu.Unlock()

}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Intn(11)%2 == 0
}

/**
conditional variable pattern:

mu.Lock()
// do something that might affect the condition
cond.Broadcast()
mu.Unlock()

----

mu.Lock()
while condition == false {
	cond.Wait()
}
// now condition is true, and we have the lock
mu.Unlock()

*/
