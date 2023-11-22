package main

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// Person представляет структуру данных для хранения информации о человеке.
type Person struct {
	Name   string // Имя человека.
	Amount int    // Количество (например, сумма денег).
}

// lineRe - это регулярное выражение для сопоставления строк, представляющих информацию о человеке.
var lineRe = regexp.MustCompile(`^Name:([^,]+), Amount:(-?\d+)\n?$`)

// errIncorrectLine - ошибка, возвращаемая, если строка не соответствует ожидаемому формату.
var errIncorrectLine = errors.New("the line is incorrect")

// errCannotParseAmount - ошибка, возвращаемая, если не удается преобразовать количество в число.
var errCannotParseAmount = errors.New("cannot parse amount")

// parseLineV1 разбирает строку в соответствии с регулярным выражением lineRe и создает структуру Person.
// Входные параметры:
//   - line: строка для разбора, представляющая информацию о человеке в формате "Name:John, Amount:42".
// Возвращаемые значения:
//   - Person: структура, представляющая информацию о человеке (имя и количество).
//   - error: ошибка, которая может быть errIncorrectLine, если строка не соответствует ожидаемому формату,
//            или errCannotParseAmount, если не удается преобразовать количество в число.
func parseLineV1(line string) (Person, error) {
	// Используем регулярное выражение для поиска совпадений в строке.
	matches := lineRe.FindStringSubmatch(line)
	// Проверяем, что было найдено как минимум два совпадения (включая всю строку и захватывающие группы).
	if len(matches) < 3 {
		// Если совпадений меньше двух, возвращаем ошибку errIncorrectLine, так как строка не соответствует ожидаемому формату.
		return Person{}, errIncorrectLine
	}

	// Извлекаем имя и количество из захватывающих групп совпадений.
	nameStr := matches[1]
	amountStr := matches[2]

	// Преобразуем количество из строки в целое число.
	amount, err := strconv.Atoi(amountStr)
	// Проверяем, произошла ли ошибка при преобразовании.
	if err != nil {
		// Если произошла ошибка, возвращаем ошибку errCannotParseAmount.
		return Person{}, errCannotParseAmount
	}

	// Создаем и возвращаем структуру Person с извлеченными данными.
	return Person{
		Name:   nameStr,
		Amount: amount,
	}, nil
}

// parseLineV2 также разбирает строку и создает структуру Person.
// В этой версии используется errors.Wrap для добавления контекста к ошибке, если преобразование не удалось.
func parseLineV2(line string) (Person, error) {
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, errors.New("incorrect line")
	}

	nameStr := matches[1]
	amountStr := matches[2]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Person{}, errors.Wrap(err, "incorrect amount")
	}

	return Person{
		Name:   nameStr,
		Amount: amount,
	}, nil
}
