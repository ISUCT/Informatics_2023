package main

import (
	"fmt"

	"github.com/stretchr/testify/internal/lab4"
)

func main() {
	fmt.Println("Хосровян Александр Арменович")

	xBeg := 0.11
	xEnd := 0.36
	detX := 0.05

	fmt.Println("Task A:")
	fmt.Println(lab4.TaskA(xBeg, xEnd, detX))

	fmt.Println("Task B:")
	var list = []float64{0.2, 0.3, 0.38, 0.43, 0.57}
	fmt.Println(lab4.TaskB(list))
}
