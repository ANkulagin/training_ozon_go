package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Person представляет информацию о человеке.
type Person struct {
	Name   string // Имя человека.
	Amount int    // Сумма, связанная с этим человеком.
}

func main() {
	// Вызываем функцию sumAmountsFromFile, чтобы вывести агрегированные суммы из файла.
	fmt.Println(sumAmountsFromFile("./data1.txt"))
}

// ParseFile читает файл и возвращает срез структур Person, представляющих информацию о людях.
func ParseFile(filename string) ([]Person, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	defer f.Close()

	return ParseReader(f)
}

// ParseReader читает данные из Reader и возвращает срез структур Person, представляющих информацию о людях.
func ParseReader(rd io.Reader) ([]Person, error) {
	data, err := io.ReadAll(rd)
	if err != nil {
		return nil, errors.Wrap(err, "reading data")
	}

	lines := strings.Split(string(data), "\n")
	persons := make([]Person, 0, len(lines))
	for _, line := range lines {
		person, err := parseLine(line)
		if err != nil {
			var myErr incorrectLineError
			if errors.As(err, &myErr) {
				return nil, errors.New("incorrect line:" + myErr.line)
			}
			return nil, errors.Wrap(err, "parsing person")
		}
		persons = append(persons, person)
	}

	return persons, nil
}

// sumAmountsFromFile читает файл и возвращает карту, содержащую суммы для каждого уникального имени.
func sumAmountsFromFile(filename string) (map[string]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	defer f.Close()

	return sumAmountsFromReader(f)
}

// sumAmountsFromReader читает данные из Reader и возвращает карту, содержащую суммы для каждого уникального имени.
func sumAmountsFromReader(r io.Reader) (map[string]int, error) {
	buf := bufio.NewReader(r)
	aggregations := make(map[string]int)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrap(err, "reading line")
		}

		line = strings.TrimSuffix(line, "\n")

		person, err := parseLine(line)
		if err != nil {
			return nil, errors.Wrap(err, "parsing line")
		}

		fmt.Println(person.Name, person.Amount)
		aggregations[person.Name] += person.Amount
	}
	return aggregations, nil
}

// lineRe - регулярное выражение для разбора строки с данными о человеке.
var lineRe = regexp.MustCompile(`^Name:([^,]+), Amount:(-?\d+)$`)

// errCannotParseAmount - ошибка, возникающая при невозможности преобразования суммы в число.
var errCannotParseAmount = errors.New("cannot parse amount")

// incorrectLineError - пользовательская ошибка для обозначения некорректной строки.
type incorrectLineError struct {
	line string
}

// Error возвращает строку с описанием ошибки некорректной строки.
func (e incorrectLineError) Error() string {
	return "incorrect line"
}

// parseLine разбирает строку и возвращает структуру Person, представляющую информацию о человеке.
func parseLine(line string) (Person, error) {
	// Поиск совпадений в строке с использованием регулярного выражения.
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, incorrectLineError{
			line: line,
		}
	}

	// Извлечение имени и суммы из совпадений.
	name := matches[1]
	amountStr := matches[2]

	// Преобразование суммы в число.
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return Person{}, errCannotParseAmount
	}

	return Person{
		Name:   name,
		Amount: amount,
	}, nil
}
