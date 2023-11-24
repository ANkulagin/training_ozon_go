package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := merge(ch1, ch2)
	ch2 <- 1
	ch1 <- 2
	// Ваш код, который генерирует данные и отправляет их в ch1 и ch2

	close(ch1)
	close(ch2)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for val := range ch3 {
			fmt.Println(val)
		}
	}()

	// Ждем завершения работы горутин
	<-done
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Функция, которая передает данные из канала в выходной канал
	output := func(c <-chan int) {
		defer wg.Done()
		for val := range c {
			out <- val
		}
	}

	wg.Add(len(channels))

	// Запускаем горутины для каждого переданного канала
	for _, c := range channels {
		go output(c)
	}

	// Закрываем канал, когда все горутины завершат свою работу
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
