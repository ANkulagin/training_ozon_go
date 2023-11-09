package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Request структура для представления JSON-запроса
type Request struct {
	ID     int              `json:"id"`    // Уникальный идентификатор запроса
	Name   string           `json:"name"`  // Имя пользователя
	Cars   []Car            `json:"cars"`  // Список автомобилей пользователя
	Params map[string]Param `json:"params"` // Параметры пользователя
}

// Car структура для представления информации об автомобиле
type Car struct {
	Plate string `json:"plate"` // Номерной знак автомобиля
	Brand string `json:"brand"` // Марка автомобиля
}

// Param структура для представления параметра пользователя
type Param struct {
	ValueID   int64  `json:"value_id"`   // Уникальный идентификатор значения параметра
	ValueName string `json:"value_name"` // Наименование значения параметра
}

func main() {
	// JSON-строка, представляющая запрос
	data := `
{
	"id": 123,
	"name": "Александр Пушкин",
	"cars": [
		{ "plate":"e123kx777", "brand": "Acura" },
		{ "plate":"a456mp777", "brand": "Ford"  }
	],
	"params": {
		"occupation": {"value_id":57, "value_name":"writer"         },
		"city":       {"value_id":5,  "value_name":"Санкт-Петербург"}
	}
}
`

	var request Request
	// Распаковываем JSON в структуру Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		log.Fatal(err) // В случае ошибки завершаем программу с выводом ошибки
	}

	// Выводим структуру Request на экран
	fmt.Printf("%#v\n", request)
}
