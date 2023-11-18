package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	a := 2.25
	fmt.Println("Задача A")
	x, y := lab4.TaskA(1.2, 2.7, 0.3, a)
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
