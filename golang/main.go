package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/laba4"
)

func main() {
	fmt.Println("Чмирков Максим Игоревич")
	a1 := laba4.TaskA(1.5, 3.5, 4.1, 2.7, 0.4)
	a2 := laba4.TaskB(4.1, 2.7, []float64{1.9, 2.15, 2.34, 2.74, 3.16})
	fmt.Println(a1)
	fmt.Println(a2)
}
