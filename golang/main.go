package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Ахмадеева Арина Игоревна")

	taskAResults := lab4.TaskA(2.0, 0.95, 1.25, 2.75, 0.3, 2.71828)
	fmt.Println("Результаты TaskA:", taskAResults)

	taskBResults := lab4.TaskB(2.0, 0.95, 2.71828, 2.2, 3.78, 4.51, 6.58, 1.2)
	fmt.Println("Результаты TaskB:", taskBResults)
}
