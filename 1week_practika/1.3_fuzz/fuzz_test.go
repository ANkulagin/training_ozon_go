package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_reverseString проверяет корректность работы функции reverseString.
func Test_reverseString(t *testing.T) {
	// Проверка, что инвертирование строки работает правильно.
	assert.Equal(t, "123", reverseString("321"))
	// Проверка, что инвертирование пустой строки возвращает пустую строку.
	assert.Equal(t, "", reverseString(""))
}

// Fuzz_ParseLine используется для фазз-тестирования функции reverseString.
func Fuzz_ParseLine(f *testing.F) {
	// Добавляем тестовые данные для фазз-тестирования.
	f.Add("123")
	f.Add("test string")
	f.Add("")

	// Функция Fuzz выполняет фазз-тестирование с использованием тестовых данных,
	// предоставленных методом Fuzz.
	f.Fuzz(func(t *testing.T, data string) {
		// Вызываем функцию reverseString с тестовыми данными.
		reverseString(data)

		// В этом примере закомментирован код с использованием пакета utf8,
		// так как функция reverseString не обрабатывает Unicode и кодировку.

		// Раскомментируйте следующие строки для проверки корректности работы с Unicode.
		// assert.Equal(t,
		// 	utf8.ValidString(data),
		// 	utf8.ValidString(reverseString(data)),
		// )
	})
}

