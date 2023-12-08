package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Formula(x float64) float64 {
	return (math.Acos(math.Pow(x, 2)-math.Pow(0.06, 2)) / math.Asin(math.Pow(x, 2)-math.Pow(0.05, 2)))
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
