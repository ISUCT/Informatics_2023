package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Терентьева Анастасия Романовна")

	a := 0.8
	b := 0.4
	x_begin := 1.23
	x_end := 7.23
	x_step := 1.2

	fmt.Println("Решения задачи А:")
	for x := x_begin; x <= x_end; x += x_step {
		fmt.Println(problem_1(a, b, x))
	}

	fmt.Println("Решения задачи B:")

	digits := [5]float64{1.88, 2.26, 3.84, 4.55, -6.21}
	answer := []float64{}

	for i := 0; i <= len(digits)-1; i++ {
		answer = append(answer, problem_1(a, b, digits[i]))
	}

	fmt.Println(answer)
}

func problem_1(a, b, x float64) float64 {
	result := (math.Pow(math.Pow((x-a), 2), drob(3.0)) + math.Pow(math.Abs(x+b), drob(5.0))) / (math.Pow((math.Pow(x, 2) - math.Pow((a+b), 2)), drob(9.0)))
	return result
}

func drob(y float64) float64 {
	result_1 := 1 / y
	return result_1
}
