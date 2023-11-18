package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

// вариант 12
func main() {
	var a float64 = 1.6
	fmt.Println("Задача A")
	x, y := lab4.TaskA(a, 1.2, 3.7, 0.5)
	for i, _ := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
	fmt.Println("Задача B")
	var listik []float64 = []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	x, y = lab4.TaskB(a, listik)
	for i, _ := range y {
		fmt.Println("При x=", x[i], "y=", y[i])
	}
}
