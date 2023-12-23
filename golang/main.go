package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func output(answersX, answersY []float64) {
	for i := range answersX {
		fmt.Printf("x = %.2f, y = %.2f\n", answersX[i], answersY[i])
	}
}

func main() {
	fmt.Println("Бусыгин Андрей Михайлович")
	xL, yL := lab4.TaskA(1.25, 3.25, 0.4)
	output(xL, yL)
	fmt.Println("__________________")
	xL, yL = lab4.TaskB([]float64{1.84, 2.71, 3.81, 4.56, 5.62})
	output(xL, yL)
}
