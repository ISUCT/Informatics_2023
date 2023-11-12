package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func PrintResult(xL []float64, yL []float64) {
	for i := range yL {
		fmt.Printf("x = %.2f, y = %f\n", xL[i], yL[i])
	}
}

func main() {
	fmt.Println("Аксенов Иван Дмитриевич")
	var b = 2.5
	var begX = 1.28
	var finX = 3.28
	var detX = 0.4

	var yList, corrXList = lab4.TaskA(begX, finX, detX, b)
	PrintResult(corrXList, yList)

	var xList = []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	yList = lab4.TaskB(xList, b)
	PrintResult(xList, yList)

}
