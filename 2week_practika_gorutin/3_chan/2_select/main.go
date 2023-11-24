package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для получения результата RPC-вызова
	chanForResp := make(chan int)

	// Запускаем RPC-вызов в горутине
	go RPCCall(chanForResp)

	// Используем select для выбора между результатом RPC-вызова и тайм-аутом
	select {
	case result := <-chanForResp:
		// Если пришел результат, выводим его
		fmt.Println(result)
	case <-time.After(time.Second * 2):
		// Если прошло 2 секунды (тайм-аут), выводим сообщение о том, что вызов превысил ожидание
		fmt.Println("RPC call timed out")
	}
}

// RPCCall имитирует RPC-вызов, блокируясь на час и затем отправляя случайное число в канал
func RPCCall(ch chan<- int) {
	// Засыпаем на час, имитируя долгий RPC-вызов
	time.Sleep(time.Hour * 24)

	// Отправляем случайное число в канал после окончания RPC-вызова

	ch <- 11
}
