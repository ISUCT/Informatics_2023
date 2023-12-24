package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Task A")
	answersA := lab4.TaskA(0.2, 2.2, 0.4)
	lab4.Output(answersA)

	fmt.Println("\nTask B")
	answersB := lab4.TaskB([]float64{0.1, 0.9, 1.2, 1.5, 2.3})
	lab4.Output(answersB)
}
