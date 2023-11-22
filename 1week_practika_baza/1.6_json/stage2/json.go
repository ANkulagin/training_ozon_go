package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Request представляет структуру данных для JSON-запроса.
type Request struct {
	ID     int              `json:"id"`     // Идентификатор запроса
	Name   string           `json:"name"`   // Имя в запросе
	Cars   []Car            `json:"cars"`   // Список машин в запросе
	Params map[string]Param `json:"params"` // Параметры в запросе
}

// Car представляет структуру данных для информации о машине в запросе.
type Car struct {
	Plate string `json:"plate"` // Номер машины
	Brand string `json:"brand"` // Марка машины
}

// Param представляет структуру данных для параметра в запросе.
type Param struct {
	ValueID   int64  `json:"value_id"`   // Идентификатор значения параметра
	ValueName string `json:"value_name"` // Имя значения параметра
}

func main() {
	// Создаем экземпляр Request с данными.
	request := Request{
		ID:   123,
		Name: "Зубенко Михаил",
		Cars: []Car{},
		Params: map[string]Param{
			"occupation": {
				ValueID:   850,
				ValueName: "mafia",
			},
		},
	}

	// Преобразуем структуру данных в JSON.
	rawJSON, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим JSON-представление структуры данных.
	fmt.Println(string(rawJSON))
}
