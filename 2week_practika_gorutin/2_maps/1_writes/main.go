package main

import (
	"fmt"
	"sync"
)

func main() {
	writers := 1000
	storage := make(map[int]int)
	var mu sync.Mutex
	wg := sync.WaitGroup{}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(storage)
}
