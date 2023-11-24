package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
	state State
}

type State int

const (
	InitialState  State = iota
	FirstState    State = iota
	SecondState   State = iota
	FinishedState State = iota
)

// FirstProcessing обрабатывает первую часть работы, умножая каждое значение на Pi.
func FirstProcessing(jobs []job, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		jobs[i].value = int64(float64(jobs[i].value) * math.Pi)
		jobs[i].state = FirstState
	}
}

// SecondProcessing обрабатывает вторую часть работы, умножая каждое значение на E.
func SecondProcessing(jobs []job, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		jobs[i].value = int64(float64(jobs[i].value) * math.E)
		jobs[i].state = SecondState
	}
}

// LastProcessing обрабатывает последнюю часть работы, деля каждое значение на случайное число от 0 до 9.
func LastProcessing(jobs []job, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range jobs {
		jobs[i].value = int64(float64(jobs[i].value) / float64(rand.Intn(10)))
		jobs[i].state = FinishedState

	}
}

func main() {
	length := 5_000_000
	jobs := make([]job, length)
	for i := 0; i < length; i++ {
		jobs[i].value = int64(i)
	}

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)

	// Разделяем массив работы на три части и обрабатываем каждую часть в своей горутине.
	go FirstProcessing(jobs[:length/2], &wg)
	go SecondProcessing(jobs[length/2:length], &wg)
	go LastProcessing(jobs, &wg)
	for i := 0; i < length; i++ {
		fmt.Println(jobs[i].value)
	}
	// Ждем завершения всех горутин.
	wg.Wait()

	finished := time.Since(start)
	fmt.Println(finished)
}
