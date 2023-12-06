package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Task A")
	answersA := lab4.TaskA(1.25, 3.25, 0.4)
	fmt.Println(answersA)
	fmt.Println("\nTask B")
	answersB := lab4.TaskB([5]float64{1.84, 2.71, 3.81, 4.56, 5.62})
	fmt.Println(answersB)
}
