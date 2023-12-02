package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func prnt(xL []float64, yL []float64) {
	for i := range yL {
		fmt.Printf("x = %.2f, y = %f\n", xL[i], yL[i])
	}
}

func main() {
	fmt.Println("Aminov_Nikita_Andreevich")

	//lab3 + 4 - вариант 1
	//task 1
	var a = 1.35
	var b = 0.98
	prnt(lab4.TaskA(1.14, 4.24, 0.62, a, b))

	//task 2
	var xList = []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	prnt(xList, lab4.TaskB(xList, a, b))

}
