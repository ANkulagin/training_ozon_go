package main

import (
	"fmt"
	"sync"
)

func main() {
	// counter - количество горутин, которые нужно запустить
	counter := 20

	// wg - WaitGroup для синхронизации завершения горутин
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на значение counter
	wg.Add(counter)

	// Запускаем цикл для создания и запуска 20 горутин
	for i := 0; i < counter; i++ {
		// Запускаем анонимную горутину с параметром i
		go func(in int) {
			// Уменьшаем счетчик WaitGroup при завершении горутины
			defer wg.Done()

			// Выводим квадрат числа i
			fmt.Println(in * in)
		}(i)
	}

	// Ждем, пока все горутины завершатся
	wg.Wait()
}
