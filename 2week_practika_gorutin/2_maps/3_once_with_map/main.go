package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	// Создаем мьютекс для безопасной работы с хранилищем и канал для уникальных идентификаторов
	storage := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000
	doubles := make([]int, 0, capacity)

	// Генерируем случайные числа и добавляем их в слайс
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}

	// Создаем буферизованный канал для уникальных идентификаторов
	uniqueIDs := make(chan int, capacity)

	// Запускаем горутины для проверки уникальности идентификаторов
	for i := 0; i < capacity; i++ {
		go func(i int) {
			// Захватываем мьютекс для безопасного доступа к хранилищу
			mu.Lock()
			defer mu.Unlock()

			// Проверяем, есть ли идентификатор в хранилище
			if _, ok := storage[doubles[i]]; !ok {
				// Если нет, добавляем его в хранилище и отправляем в канал
				storage[doubles[i]] = struct{}{}
				uniqueIDs <- doubles[i]
			}
		}(i)
	}

	close(uniqueIDs) // Закрываем канал, так как больше не будем в него писать
	fmt.Printf("len of ids: %d \n", len(uniqueIDs))
	fmt.Println(uniqueIDs)
}
