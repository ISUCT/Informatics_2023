package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Задача A ", lab4.TaskA(0.11, 0.36, 0.05))
	fmt.Println("Задача B ", lab4.TaskB([]float64{0.08, 0.26, 0.35, 0.41, 0.53}))
func formula(x float64) float64 {
	return math.Asin(math.Pow(x, 2)) + math.Asin(math.Pow(x, 3))
}

// 11 вариант
