package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Сысуйкин Денис Дмитриевич")
	SliceTaskA := lab4.TaskA(2.25, 1.2, 2.7, 0.3)
	fmt.Println("Результаты TaskA:", SliceTaskA)
	SliceForTaskB := []float64{2.25, 1.31, 1.39, 1.44, 1.56, 1.92}
	resultS := lab4.TaskB(SliceForTaskB, 2.25)
	fmt.Println("Результаты TaskB:", resultS)
}
