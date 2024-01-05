package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Журавлев Сергей Викторович")

	a := -2.5
	b := 3.4
	x_begin := 3.5
	x_end := 6.5
	x_delta := 0.6

	fmt.Println("Решения задачи под А:")
	values := lab4.TaskA(a, b, x_begin, x_end, x_delta)
	lab4.Output(values)

	fmt.Println("Решения задачи под В:")
	values = lab4.TaskB(a, b, []float64{2.89, 3.54, 5.21, 6.28, 3.48})
	lab4.Output(values)
}
