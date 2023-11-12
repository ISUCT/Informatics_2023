package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Кочетов Данил Романович")

	a := 0.05
	b := 0.06
	x_begin := 0.2
	x_end := 0.95
	x_step := 0.15

	fmt.Println("Решения задачи А:")
	for x := x_begin; x <= x_end; x += x_step {
		fmt.Println(problem_1(a, b, x))
	}

	fmt.Println("Решения задачи B:")

	digits := [5]float64{0.15, 0.26, 0.37, 0.48, 0.56}
	answer := []float64{}

	for i := 0; i <= len(digits)-1; i++ {
		answer = append(answer, problem_1(a, b, digits[i]))
	}

	fmt.Println(answer)
}

func problem_1(a, b, x float64) float64 {
	result := math.Acos(math.Pow(x, 2)-math.Pow(b, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(a, 2))
	return result
}
