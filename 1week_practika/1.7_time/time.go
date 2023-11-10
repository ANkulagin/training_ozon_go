package main

import (
	"fmt"
	"time"
)

func main() {
	// Вызываем различные функции для демонстрации работы с временем.
	// Расскомментируйте нужные функции для запуска.
	// simple()
	// sleepAndDuration()
	// parseAndFormat()
	 timeMath()
	//test()
}

// simple демонстрирует получение текущего времени.
func simple() {
	t := time.Now()
	fmt.Println(t)
}

// sleepAndDuration демонстрирует работу с ожиданием и замером времени.
func sleepAndDuration() {
	t := time.Now()
	time.Sleep(time.Second * 2)
	elapsed := time.Since(t)

	fmt.Println(elapsed)
}

// parseAndFormat демонстрирует парсинг и форматирование времени,
// а также работу с часовыми поясами.
func parseAndFormat() {
	// Форматируем текущее время в заданном формате.
	fmt.Println(time.Now().Format("2006-01-02 PM 03:04:05 Z07:00"))

	// Задаем часовой пояс для Москвы.
	loc := time.FixedZone("Moscow", 3*60*60)
	// Парсим строку времени с указанием часового пояса.
	moscowTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-09-24 12:49:01", loc)
	fmt.Println(moscowTime, err)

	// Задаем часовой пояс для Нью-Йорка.
	targetLoc := time.FixedZone("NewYork", -4*60*60)
	// Преобразуем время из Московского в Нью-Йоркское.
	newYorkTime := moscowTime.In(targetLoc)
	fmt.Println(newYorkTime)

	// Выводим Unix-время для Москвы и Нью-Йорка.
	fmt.Println(moscowTime.Unix())
	fmt.Println(newYorkTime.Unix())

	// Выводим час в Нью-Йорке.
	fmt.Println(newYorkTime.Hour())
}

// timeMath демонстрирует математические операции с временем.
func timeMath() {
	now:= time.Now()

	// Добавляем 2 часа и 20 минут к текущему времени.
	twoHoursAnd20minLater := now.Add(time.Hour*2 + time.Minute*20)
	// Отнимаем 10 минут от текущего времени.
	tenMinutesBefore := now.Add(-time.Minute * 10)

	// Выводим результаты.
	fmt.Println("now + 2h 20m:", twoHoursAnd20minLater)
	fmt.Println("         now:", now)
	fmt.Println("   now - 10m:", tenMinutesBefore)

	// Добавляем 2 месяца и 15 дней к текущему времени.
	after2months15days := now.AddDate(0, 2, 15)
	// Отнимаем 1 год от текущего времени.
	oneYearBefore := now.AddDate(-1, 0, 0)

	fmt.Println()

	// Выводим результаты.
	fmt.Println("         1 year before:", oneYearBefore.Format("2006-01-02"))
	fmt.Println("                   now:", now.Format("2006-01-02"))
	fmt.Println("after 2 months 15 days:", after2months15days.Format("2006-01-02"))

	fmt.Println()

	// Выводим разницу во времени и проверяем отношения времени.
	fmt.Println("             now - tenMinutesBefore:", now.Sub(tenMinutesBefore))
	fmt.Println("      now is after tenMinutesBefore:", now.After(tenMinutesBefore))
	fmt.Println("now is before twoHoursAnd20minLater:", now.Before(twoHoursAnd20minLater))

	fmt.Println()

	// Сравниваем длительности времени.
	fmt.Println("      time.Minute < time.Hour:", time.Minute < time.Hour)
	fmt.Println("time.Second*120 > time.Minute:", time.Second*120 > time.Minute)
}

// test демонстрирует добавление года к заданному времени.
func test() {
	// Задаем часовой пояс для Москвы.
	loc := time.FixedZone("Moscow", 3*60*60)
	// Парсим строку времени с указанием часового пояса.
	moscowTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-01-31 12:49:01", loc)
	fmt.Println(moscowTime, err)
	// Добавляем 1 год к времени.
	moscowTime = moscowTime.AddDate(1, 0, 0)
	fmt.Println(moscowTime)
}
