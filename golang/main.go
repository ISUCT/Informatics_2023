package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/laba4"
)

func main() {
	fmt.Println("Яковлева Алёна")
	y1 := laba4.Task1(0.22, 0.92, 0.14)
	y2 := laba4.Task2([]float64{0.1, 0.35, 0.4, 0.55, 0.6})
	fmt.Println(y1)
	fmt.Println(y2)
}
