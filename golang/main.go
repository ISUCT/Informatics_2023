// gurev ivan 1-279 formula 6
package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Task A")
	xA, yA := lab4.TaskA(0.2, 2.2, 0.4)
	for j := range xA {
		fmt.Println("При x = ", xA[j], ", y = ", yA[j])
	}

	fmt.Println("Task B")
	var xB = []float64{0.1, 0.9, 1.2, 1.5, 2.3}
	yB := lab4.TaskB(xB)
	for j := range xB {
		fmt.Println("При x = ", xB[j], ", y = ", yB[j])
	}
}
