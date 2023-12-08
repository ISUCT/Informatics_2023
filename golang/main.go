package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Кочетов Максим Романович")
	fmt.Println("Lab 4")

	a := 2.0
	b := 3.0
	x_beg := 0.11
	x_end := 0.36
	x_stp := 0.05

	xList := []float64{0.08, 0.26, 0.35, 0.41, 0.53}

	fmt.Println(lab4.Task_A(x_beg, x_end, x_stp, a, b))
	fmt.Println(lab4.Task_B(xList, a, b))
}
