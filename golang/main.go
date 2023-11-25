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
	fmt.Println("Горюнов Александр Алексеевич")

	var a = -2.5
	var b = 3.4
	var begX = 3.5
	var endX = 6.5
	var detX = 0.6

	prnt(lab4.TaskA(a, b, begX, endX, detX))

	var xList = []float64{2.89, 3.54, 5.21, 6.28, 3.48}

	prnt(lab4.TaskB(xList, a, b), xList)
}
