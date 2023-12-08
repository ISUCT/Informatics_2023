package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Долев Максим Алексеевич")
	fmt.Println("Lab 4")

	a := 2.0
	b := 0.95
	x_beg := 1.25
	x_end := 2.75
	x_stp := 0.3

	xList := []float64{2.2, 3.78, 4.51, 6.58, 1.2}

	fmt.Println(lab4.Task_A(x_beg, x_end, x_stp, a, b))
	fmt.Println(lab4.Task_B(xList, a, b))
}
