package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct{}

var SingleInstance *Single

func getInstance() *Single {
	if SingleInstance == nil {
		once.Do(func() {
			fmt.Println("creating instance")
			SingleInstance = &Single{}
		})
	} else {
		fmt.Println("instance already created")
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
