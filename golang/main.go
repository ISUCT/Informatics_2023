package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	// Лабораторная 2
	fmt.Println("Байгулова Ева Валерьевна")

	// Лабораторная 3,4 - Вариант 3
	fmt.Println("Значение функции №3 при данных задачи А =", internal.TaskA(1.25, 2.75, 0.3))
	fmt.Println("Значение функции №3 при данных задачи B =", internal.TaskB([]float64{2.2, 3.78, 4.51, 6.58, 1.2}))
}
