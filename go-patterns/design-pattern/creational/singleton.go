package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type Single struct{}

var SingleInstance *Single

func getInstance() *Single {
	if SingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingleInstance == nil {
			fmt.Println("Creating single instance now.")
			SingleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return SingleInstance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getInstance()
		}()
	}
	wg.Wait()
}
