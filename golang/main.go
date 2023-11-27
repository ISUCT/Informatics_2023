package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Корчагин Пётр Сергеевич")
	fmt.Println("Ответы на задачу A")
	xN := 0.7
	xK := 2.2
	xD := 0.2
	x, y := lab4.TaskA(xN, xK, xD)
	for i := range y {
		fmt.Println("x=", x[i], "y=", y[i])
	}
	fmt.Println("Ответы на задачу B")
	xArray := []float64{0.25, 0.36, 0.56, 0.94, 1.28}
	x, y = lab4.TaskB(xArray)
	for i := range y {
		fmt.Println("x=", x[i], "y=", y[i])
	}
}
