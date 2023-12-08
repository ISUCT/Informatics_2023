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
	fmt.Println("Соколова Арина Михайловна")

	//лабараторная работа №3 - 4 вариант 21
	//TASK 1
	prnt(lab4.TaskA(0.11, 0.36, 0.05))

	//TASK 2
	var xList = []float64{0.2, 0.3, 0.38, 0.43, 0.57}

	prnt(lab4.TaskB(xList), xList)
}
