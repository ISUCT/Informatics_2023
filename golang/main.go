package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Солдатов Иван Сергеевич")
	fmt.Println("------------------------First_Exercise-------------------------------")
	a := 2.0
	x_end := 4.2
	x_step := 0.6
	for x_start := 1.2; x_start <= x_end; x_start += x_step {
		fmt.Println(formula(a, x_start, x_end, x_step))
	}
	fmt.Println("------------------------Second_Exercise-------------------------------")
	array := []float64{1.16, 1.32, 1.47, 1.65, 1.93}
	for _, x_start := range array {
		fmt.Println(formula(a, x_start, x_end, x_step))
	}
}

func formula(a, x_start, x_end, x_step float64) float64 {
	return ((math.Pow((math.Log(a + x_start)), 2)) / (math.Pow((a + x_start), 2)))
}
