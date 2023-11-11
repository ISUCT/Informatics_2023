package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Кочетов Максим Романович")

	a := 2.0
	b := 3.0
	x_start := 0.11
	x_end := 0.36
	x_step := 0.05

	fmt.Println("Решение задачи А:")

	for x := x_start; x <= x_end; x += x_step {
		fmt.Println(zadacha1(a, b, x))
	}

	fmt.Println("Решение задачи Б:")

	massiv := [5]float64{0.08, 0.26, 0.35, 0.41, 0.53}
	answer := []float64{}

	for i := 0; i <= len(massiv)-1; i++ {
		answer = append(answer, zadacha1(a, b, massiv[i]))
	}

	fmt.Println(answer)

}

func zadacha1(a, b, x float64) float64 {
	result := math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b))
	return result
}
