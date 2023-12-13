package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func rez(x, y []float64) {
	for i, v := range y {
		fmt.Printf("x = %f | y = %f \n", x[i], v)
	}
}

func main() {
	fmt.Println("Антонов Кирилл Юрьевич")
	var b = 2.5
	var begX = 1.28
	var finX = 3.28
	var detX = 0.4

	var ifeal, fink = lab4.TaskA(begX, finX, detX, b)
	rez(fink, ifeal)

	var xList = []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	ifeal = lab4.TaskB(xList, b)
	rez(xList, ifeal)
}
