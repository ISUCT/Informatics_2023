package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	//Лабораторная 2
	fmt.Println("Naumova Alena Alekseevna")

	//Лабораторная 3 - вариант 14

	//Решение под А
	a := 7.2
	b := 4.2
	xBegin := 1.81
	xEnd := 5.31
	xDelta := 0.7
	fmt.Println("Решения под A:")
	fmt.Println(lab4.TaskA(xBegin, xEnd, xDelta, a, b))

	//Решение под Б
	fmt.Println("Решения под B:")
	fmt.Println(lab4.TaskB([]float64{2.4, 2.8, 3.9, 4.7, 3.16}, a, b))
}
