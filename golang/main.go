package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Kochetkov Kirill")

	res1 := lab4.Task1(7.2, 4.2, 1.81, 5.31, 0.7)
	fmt.Println(res1)

	res2 := lab4.Task2(7.2, 4.2, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(res2)
}
