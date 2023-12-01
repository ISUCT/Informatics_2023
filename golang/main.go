package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Киселев Ярослав Денисович")
	fmt.Println("---")

	a1 := Task1(0.05, 0.06, 0.2 0.95, 0.15)
	fmt.Print(a1)
	fmt.Println("---")

	a2 := Task2(0.05, 0.06, []float64{0.15, 0.26, 0.37, 0,48, 0.56})
	fmt.Println(a2)
}
