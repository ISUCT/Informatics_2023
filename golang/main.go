package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Якиманская Мария Алексеевна")

	a1 := lab4.Task1(7.2, 1.3, 1.56, 4.71, 0.63)
	fmt.Println(a1)

	a2 := lab4.Task2(7.2, 1.3, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(a2)
}
