package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Родин Владислав Евгеньевич")
	a := 2.0
	b := 1.1
	Xn := 0.08
	Xa := 1.08
	step := 0.2
	list := []float64{0.1, 0.3, 0.4, 0.45, 0.65}
	fmt.Println("Задача А")
	x, y := lab4.TaskA(Xn, Xa, step, a, b)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("Задача B")
	x, y = lab4.TaskB(list, a, b)
	for i := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
}
