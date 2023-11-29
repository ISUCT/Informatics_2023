package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Formula(x float64) float64 {
	a := 2.0
	b := 0.95
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - (math.Exp(x / a)))
}

func TaskA(x_begin, x_end, x_delta float64) []float64 {
	var num_of_elements int = int((x_end) - (x_begin)/x_delta + 1)
	var otvetA = make([]float64, 0, num_of_elements)
	for i := x_begin; i <= x_end; i = i + x_delta {
		otvetA = append(otvetA, Formula(i))
	}
	return otvetA
}

func TaskB(slice []float64) []float64 {
	var otvetB = make([]float64, 0, len(slice))
	for _, i := range slice {
		otvetB = append(otvetB, Formula(i))
	}
	return otvetB
}
