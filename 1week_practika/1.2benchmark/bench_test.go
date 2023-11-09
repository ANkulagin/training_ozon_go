package main

import (
	"testing"
)

func Benchmark_v1(b *testing.B) {
	// Подготовка данных: создание строки с некорректным форматом (incorrectLine) и строкой, содержащей ошибку в поле Amount (incorrectAmount).
	incorrectLine := "sdfsdfsdf"
	incorrectAmount := "Name:John, Amount:sdf"

	// Запуск бенчмарка для функции parseLineV1.
	for i := 0; i < b.N; i++ {
		parseLineV1(incorrectLine)
		parseLineV1(incorrectAmount)
	}
}

func Benchmark_v2(b *testing.B) {
	// Подготовка данных: создание строки с некорректным форматом (incorrectLine) и строкой, содержащей ошибку в поле Amount (incorrectAmount).
	incorrectLine := "sdfsdfsdf"
	incorrectAmount := "Name:John, Amount:sdf"

	// Запуск бенчмарка для функции parseLineV2.
	for i := 0; i < b.N; i++ {
		parseLineV2(incorrectLine)
		parseLineV2(incorrectAmount)
	}
}
