package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func prnt(ylist, xlist []float64) {
	for i, v := range ylist {
		fmt.Printf("x = %f  y = %f\n", xlist[i], v)
	}
}

func main() {
	//лабораторная работа №2
	fmt.Println("Горюнов Александр Алексеевич")

	//лабароторная работа №3 + 4 - 5 вариант
	//задача 1
	var a = -2.5
	var b = 3.4

	prnt(lab4.TaskA(a, b, 3.5, 6.5, 0.6))

	//задача 2
	var xList = []float64{2.89, 3.54, 5.21, 6.28, 3.48}

	prnt(lab4.TaskB(xList, a, b), xList)
}
