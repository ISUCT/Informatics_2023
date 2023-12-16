package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

// вариант 12
func print_answer(xArr, yArr []float64) {
	for i := range xArr {
		fmt.Println("При x=", xArr[i], "y=", yArr[i])
	}
}
func main() {
	var a float64 = 1.6

	x, y := lab4.TaskA(a, 1.2, 3.7, 0.5)
	print_answer(x, y)

	var listik []float64 = []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	x, y = lab4.TaskB(a, listik)
	print_answer(x, y)
}
