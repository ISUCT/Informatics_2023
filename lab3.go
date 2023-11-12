package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Долев Максим Алексеевич")

	a := 2.0
	b := 0.95
	x_begin := 1.25
	x_end := 2.75
	x_step := 0.3

	fmt.Println("Решения задачи А:")
	for x := x_begin; x <= x_end; x += x_step {
		fmt.Println(problem_1(a, b, x))
	}

	fmt.Println("Решения задачи B:")

	digits := [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}
	answer := []float64{}

	for i := 0; i <= len(digits)-1; i++ {
		answer = append(answer, problem_1(a, b, digits[i]))
	}

	fmt.Println(answer)
}

func problem_1(a, b, x float64) float64 {
	result := 1 + (math.Pow(math.Log10(x/a), 2))/b - math.Pow(math.E, x/a)
	return result
}
