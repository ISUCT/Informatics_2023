package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab3"
)

func main() {
	fmt.Println("Кочетов Максим Романович")
	fmt.Println("Lab 3")

	a := 2.0
	b := 3.0
	x_beg := 0.11
	x_end := 0.36
	x_stp := 0.05

	xList := []float64{0.08, 0.26, 0.35, 0.41, 0.53}

	answer_task_1 := lab3.Task_A(x_beg, x_end, x_stp, a, b)
	fmt.Println(answer_task_1)

	answer_task_2 := lab3.Task_B(xList, a, b)
	fmt.Println(answer_task_2)
}
