package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Парамонов Илья Сергеевич 1/280")
	fmt.Println("16 вариант")
	fmt.Println("Задача А")
	fmt.Println(internal.TaskA(0.08, 1.08, 0.2))
	fmt.Println("Задача B")
	fmt.Println(internal.TaskB([]float64{0.1, 0.3, 0.4, 0.45, 0.65}))
}
