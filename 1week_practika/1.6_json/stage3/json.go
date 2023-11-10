package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Command представляет команду с определенным типом и данными.
type Command struct {
	Type string          `json:"type"` // Тип команды
	Data json.RawMessage `json:"data"` // Сырые JSON-данные команды
}

// Request представляет массив команд.
type Request []Command

// SendMessageData представляет структуру данных для команды "send_message".
type SendMessageData struct {
	User int64  `json:"user"` // Идентификатор пользователя
	Text string `json:"text"` // Текст сообщения
}

// MakeOrderData представляет структуру данных для команды "make_order".
type MakeOrderData struct {
	Sku    int64 `json:"sku"`    // Артикул товара
	Amount int   `json:"amount"` // Количество товара
}

func main() {
	// JSON-строка с массивом команд.
	data := `
[
	{
		"type": "send_message",
		"data": {
			"user": 61254895,
			"text": "Hello!"
		}
	},
	{
		"type": "make_order",
		"data": {
			"sku": 12345678,
			"amount": 2
		}
	}
]
`

	// Распаковываем JSON-данные в массив команд.
	var request Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		log.Fatal(err)
	}

	// Итерируемся по каждой команде и обрабатываем ее в соответствии с типом.
	for _, command := range request {
		switch command.Type {
		case "send_message":
			var data SendMessageData
			err := json.Unmarshal(command.Data, &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("send_message: %#v\n", data)
		case "make_order":
			var data MakeOrderData
			err := json.Unmarshal(command.Data, &data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("make_order: %#v\n", data)
		}
	}
}
