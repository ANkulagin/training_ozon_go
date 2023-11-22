package main

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// main - основная функция программы
func main() {
    // Отправляем HTTP-запрос для получения сжатого файла CSV
    res, err := http.Get("http://192.168.5.110:8080/storage/temp/geoip.csv.bz2")
    if err != nil {
        log.Fatal(err) // Если произошла ошибка при выполнении запроса, завершаем программу с ошибкой
    }
    defer res.Body.Close() // Закрываем тело ответа после завершения работы функции

    // Проверяем, что HTTP-код ответа - OK (200)
    if res.StatusCode != http.StatusOK {
        log.Fatal("wrong status code:", res.StatusCode) // Если код ответа не OK, завершаем программу с ошибкой
    }

    // Выводим результат работы функции countZones на экран, передавая ей тело ответа и страну "RU"
    fmt.Println(countZones(res.Body, "RU"))
}


// startIpNum, endIpNum,     country,  region,  city,  postalCode,  latitude,  longitude, dmaCode, areaCode
// 1.0.0.0,    1.7.255.255,  "AU",     "",      "",    "",          -27.0000,  133.0000,   ,
// 1.9.0.0,    1.9.255.255,  "MY",     "",      "",    "",          2.5000,    112.5000,   ,
// 1.10.10.0,  1.10.10.255,  "AU",     "",      "",    "",          -27.0000,  133.0000,   ,

// Функция countZones считает количество строк в CSV-файле,
// где указана заданная страна.
// Возвращает количество строк и ошибку, если что-то идет не так.
func countZones(reader io.Reader, country string) (int, error) {
	// Инициализация читателей для разархивации и чтения CSV
	bzipReader := bzip2.NewReader(reader)
	csvReader := csv.NewReader(bzipReader)

	// Счетчик строк для указанной страны
	counter := 0

	// Счетчик общего количества обработанных строк
	rowId := 0

	// Цикл обработки строк CSV
	for {
		rowId++

		// Логгирование каждого миллиона обработанных строк
		if rowId%1_000_000 == 0 {
			log.Println("processed", rowId, "rows")
		}

		// Чтение строки из CSV
		row, err := csvReader.Read()

		// Проверка на конец файла
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, errors.Wrap(err, "reading csv")
		}

		// Пропуск строк с неправильным форматом данных
		if len(row) < 3 {
			continue
		}

		// Проверка, соответствует ли страна заданной стране
		if row[2] == country {
			counter++
		}
	}

	// Логгирование общего количества обработанных строк
	log.Println("total rows:", rowId)

	// Возвращение результата
	return counter, nil
}
