package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Парамонов Илья Сергеевич 1/280")
	fmt.Println("16 вариант")
	a := 2.0
	b := 1.1
	fmt.Println("Задача А")
	x_begin := 0.08
	x_end := 1.08
	x_delta := 0.2
	fmt.Println(internal.TaskA(x_begin, x_end, x_delta, a, b))
	fmt.Println("Задача B")
	slice := []float64{0.1, 0.3, 0.4, 0.45, 0.65}
	fmt.Println(internal.TaskB(slice, a, b))
}
