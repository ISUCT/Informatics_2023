package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.05
	b := 0.48
	x_begin := 0.7
	x_end := 2.2
	x_step := 0.3

	for x := x_begin; x <= x_end; x += x_step {
		fmt.Println(problem(a, b, x))
	}

	digits := [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	answer := []float64{}

	for i := 0; i <= len(digits)-1; i++ {
		answer = append(answer, problem(a, b, digits[i]))
	}

	fmt.Println(answer)

}
func problem(a, b, x float64) float64 {
	result := math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)/math.Acos(x*b) + math.Exp(-x/2)
	return result
}
