package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Hello world")
	fmt.Printf("Jakovleva Darja Dmitrievna 1/278\n\n")

	fmt.Println("24 вариант")
	a := 7.2
	b := 1.3
	xn := 1.56
	xk := 4.71
	dx := 0.63
	vhodnie_x := []float64{2.4, 2.8, 3.9, 4.7, 3.16}

	yA := lab4.Task_A(a, b, xn, xk, dx)
	yB := lab4.Task_B(a, b, vhodnie_x)

	fmt.Println("Task_A:", yA)
	fmt.Println("Task_B:", yB)
}
