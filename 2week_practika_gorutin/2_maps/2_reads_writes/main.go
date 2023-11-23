package main

import (
	"fmt"
	"sync"
)

func main() {
	// Количество горутин для записи
	writers := 1000
	// Количество горутин для чтения
	reads := 1000

	// Хранилище данных
	storage := make(map[int]int)
	// Мьютекс для синхронизации доступа к хранилищу
	var mu sync.RWMutex
	// WaitGroup для отслеживания завершения всех горутин
	wg := sync.WaitGroup{}

	// Горутины для записи
	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go func(i int) {
			defer wg.Done()
			// Блокировка мьютекса перед записью
			mu.Lock()
			storage[i] = i
			// Разблокировка мьютекса после записи
			mu.Unlock()
		}(i)
	}

	// Горутины для чтения
	wg.Add(reads)
	for i := 0; i < reads; i++ {
		go func(i int) {
			defer wg.Done()
			// Блокировка мьютекса перед чтением
			mu.RLock()
			_ = storage[i]
			// Разблокировка мьютекса после чтения
			mu.RUnlock()
		}(i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результата
	fmt.Println(storage)
}
