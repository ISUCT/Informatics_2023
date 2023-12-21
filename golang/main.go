package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Королев Никита")
	answ1 := lab4.Task1(1.6, 1.2, 3.7, 0.05)
	answ2 := lab4.Task2(1.6, []float64{1.28, 1.36, 2.47, 3.68, 4.56})
	fmt.Println(answ1)
	fmt.Println(answ2)
}
