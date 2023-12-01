package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal/lab3"
)

func main() {
	a1 := lab3.Task1(1.2, 0.48, 0.7, 2.2, 0.3)
	a2 := lab3.Task2(1.2, 0.48, [5]float64{0.25, 0.36, 0.56, 0.94, 1.28})
	fmt.Println(a1)
	fmt.Println(a2)
}