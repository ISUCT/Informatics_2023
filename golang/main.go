package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	a := 2.25
	Xs := 1.2
	Xe := 2.7
	step := 0.3
	fmt.Println("Задача A")
	x, y := lab4.TaskA(Xs, Xe, step, a)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("Задача B")
	var list_of_x []float64 = []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	x, y = lab4.TaskB(a, list_of_x)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
}
