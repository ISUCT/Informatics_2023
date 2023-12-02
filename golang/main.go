package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Сысуйкин Денис Дмитриевич")
	fmt.Println("Результаты TaskA:", lab4.TaskA(2.25, 1.2, 2.7, 0.3))
	fmt.Println("Результаты TaskB:", lab4.TaskB(2.25, 1.31, 1.39, 1.44, 1.56, 1.92))
}
