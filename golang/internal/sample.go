package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

func Task_A(x_begin, x_end, x_step float64) []float64 {
	var otvet_a = make([]float64, 0, 6)
	for i := x_begin; i <= x_end; i += x_step {
		otvet_a = append(otvet_a, function(i))
	}
	return otvet_a
}

func Task_B(xs []float64) []float64 {
	var otvet_b = make([]float64, 0, 5)
	for _, i := range xs {
		otvet_b = append(otvet_b, function(i))
	}
	return otvet_b
}

func function(x float64) float64 {
	return (math.Pow(math.Log10(2+x), 2) / math.Pow(2+x, 2))
}