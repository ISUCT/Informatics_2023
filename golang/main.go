package main

import (
	"fmt"
	"math"

	lab4 "isuct.ru/informatics2022/laba4"
)

func formula(a, x float64) float64 {
	return ((math.Pow((math.Log(a + x)), 2)) / (math.Pow((a + x), 2)))
}

func main() {
	fmt.Println("Солдатов Иван Сергеевич")
	fmt.Println("------------------------First_Exercise-------------------------------")
	a := 2.0
	x_end := 4.2
	x_step := 0.6
	x_start := 1.2
	for x := x_start; x <= x_end; x += x_step {
		fmt.Println(formula(a, x))
	}
	fmt.Println("------------------------Second_Exercise-------------------------------")
	array := []float64{1.16, 1.32, 1.47, 1.65, 1.93}
	for _, x := range array {
		fmt.Println(formula(a, x))
	}
	fmt.Println("------------------------Lab_4-------------------------------")
	lab4.Solve(a, x_start, x_end, x_step)
}
