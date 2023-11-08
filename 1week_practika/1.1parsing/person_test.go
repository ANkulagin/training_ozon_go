package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_parseLine_ShouldFillPersonFields проверяет, что функция parseLine правильно заполняет поля структуры Person из строки.
func Test_parseLine_ShouldFillPersonFields(t *testing.T) {
	// Создаем тестовую строку.
	line := "Name:Сергей, Amount:10"

	// Вызываем функцию parseLine для анализа строки.
	person, err := parseLine(line)

	// Проверяем, что ошибки нет и поля Person заполнены правильно.
	assert.NoError(t, err)
	assert.Equal(t,
		Person{
			Name:   "Сергей",
			Amount: 10,
		},
		person,
	)
}

// Test_ParseLine_ShouldFillPersonFields_WhenAmountIsNegative проверяет, что функция parseLine правильно обрабатывает отрицательные суммы.
func Test_ParseLine_ShouldFillPersonFields_WhenAmountIsNegative(t *testing.T) {
	line := "Name:Сергей, Amount:-10"

	person, err := parseLine(line)

	assert.NoError(t, err)
	assert.Equal(t, Person{
		Name:   "Сергей",
		Amount: -10,
	}, person)
}

// Test_ParseLine_ShouldReturnError_WhenNameIsAbsent проверяет, что функция parseLine возвращает ошибку, если отсутствует имя.
func Test_ParseLine_ShouldReturnError_WhenNameIsAbsent(t *testing.T) {
	line := "Name:, Amount:10"

	_, err := parseLine(line)

	assert.Error(t, err)
}

// Test_ParseLine_ShouldReturnError_WhenAmountIsNotANumber проверяет, что функция parseLine возвращает ошибку, если сумма не является числом.
func Test_ParseLine_ShouldReturnError_WhenAmountIsNotANumber(t *testing.T) {
	line := "Name:Сергей, Amount:1asd0"

	_, err := parseLine(line)

	assert.Error(t, err)
}

// Test_ParseLine_ShouldReturnError_WhenAmountIsEmpty проверяет, что функция parseLine возвращает ошибку, если сумма пуста.
func Test_ParseLine_ShouldReturnError_WhenAmountIsEmpty(t *testing.T) {
	line := "Name:Сергей, Amount:"

	_, err := parseLine(line)

	assert.Error(t, err)
}

// Test_ParseLine_ShouldReturnError_WhenAmountTooBig проверяет, что функция parseLine возвращает ошибку, если сумма слишком велика.
func Test_ParseLine_ShouldReturnError_WhenAmountTooBig(t *testing.T) {
	line := "Name:Сергей, Amount:1111111111111111111111111111111111111111111111111111111"

	_, err := parseLine(line)

	assert.Equal(t, errCannotParseAmount, err)
}

// Test_ParseReader_WrongLineShouldGiveError проверяет, что функция ParseReader возвращает ошибку при некорректной строке.
func Test_ParseReader_WrongLineShouldGiveError(t *testing.T) {
	// Создаем данные с ошибочной строкой.
	data := "Nam, Amount:12\n" +
		"Name:Петр, Amount:65\n"
	buf := bytes.NewBufferString(data)

	// Вызываем функцию ParseReader для анализа данных.
	_, err := ParseReader(buf)

	// Проверяем, что возвращается ошибка.
	assert.Error(t, err)
}

// Test_ParseReader_CorrentLineShouldGivePersonsList проверяет, что функция ParseReader возвращает правильный список Persons при корректных данных.
func Test_ParseReader_CorrentLineShouldGivePersonsList(t *testing.T) {
	// Создаем корректные данные с двумя строками.
	data := "Name:Иван, Amount:12\n" +
		"Name:Петр, Amount:65"
	buf := bytes.NewBufferString(data)

	// Вызываем функцию ParseReader для анализа данных.
	persons, err := ParseReader(buf)

	// Проверяем, что ошибки нет и возвращается правильный список Persons.
	assert.NoError(t, err)
	assert.Equal(t,
		[]Person{
			{
				Name:   "Иван",
				Amount: 12,
			},
			{
				Name:   "Петр",
				Amount: 65,
			},
		},
		persons,
	)
}
