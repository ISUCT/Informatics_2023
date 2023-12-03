package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	n1 := lab4.Task1(4.1, 2.7, 1.2, 5.2, 0.7)
	fmt.Println(n1)
	n2 := lab4.Task2(4.1, 2.7, []float64{1.9, 2.15, 2.34, 2.73, 3.16})
	fmt.Println(n2)
	fmt.Println("Полина Костина")
}
