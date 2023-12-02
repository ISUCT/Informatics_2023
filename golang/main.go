package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Task A")
	fmt.Println(lab4.TaskA(0.2, 2.2, 0.4))

	fmt.Println("\nTask B")
	fmt.Println(lab4.TaskB([5]float64{0.1, 0.9, 1.2, 1.5, 2.3}))
}
