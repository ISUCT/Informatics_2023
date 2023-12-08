package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Каретина Светлана Олеговна")
	fmt.Println("Lab 4")

	a := 0.05
	b := 0.48
	x_beg := 0.7
	x_end := 2.2
	x_stp := 0.3

	xList := []float64{0.25, 0.36, 0.56, 0.94, 1.28}

	fmt.Println(lab4.Task_A(x_beg, x_end, x_stp, a, b))
	fmt.Println(lab4.Task_B(xList, a, b))
}
