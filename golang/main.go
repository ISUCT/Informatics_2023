package main

import (
	"fmt"
)

func main() {
	a := 2.5
	b := 4.6
	x_n := 1.1
	x_k := 3.6
	x_z := 0.5

	fmt.Println("Значение функции №18 при данных задачи А =", internal.TaskA(x_n, x_k, x_z, a, b))
	fmt.Println("Значение функции №18 при данных задачи Б =", internal.TaskB([]float64{1.2, 1.28, 1.36, 1.46, 2.35}, a, b))
}

