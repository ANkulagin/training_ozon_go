package main

import (
	"fmt"
	"sync"
)

func main() {
	var storage map[int]int
	wg := sync.WaitGroup{}
	writers := 1000
	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go func(i int) {
			defer wg.Done()
			storage[i] = i
		}(i)
	}
	wg.Wait()
	fmt.Println(storage)
}
