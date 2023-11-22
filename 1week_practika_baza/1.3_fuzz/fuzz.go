package main

func main() {

}

// reverseString принимает строку input и возвращает ее инвертированную версию.
func reverseString(input string) string {
	// Если входная строка пуста, возвращаем пустую строку.
	if len(input) == 0 {
		return ""
	}

	// Создаем буфер для хранения инвертированной строки.
	output := make([]byte, 0, len(input))

	// Итерируемся по символам входной строки в обратном порядке.
	for i := len(input) - 1; i >= 0; i-- {
		// Добавляем текущий символ в начало буфера.
		output = append(output, input[i])
	}

	// Преобразуем буфер в строку и возвращаем ее.
	return string(output)
}
