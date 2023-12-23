package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

func formula(x float64) float64 {
	return ((1 + math.Pow(math.Sin(math.Pow(2.5, 3.0)+math.Pow(x, 3.0)), 2)) / math.Cbrt(math.Pow(2.5, 3)+math.Pow(x, 3)))
}

func Zadacha_A(x_start, x_end, x_delta float64) []float64 {
	result_a := make([]float64, 0, int((x_end-x_start)/(x_delta))+1.0)
	for x := x_start; x <= x_end; x += x_delta {
		result_a = append(result_a, formula(x))
	}
	return result_a
}

func Zadacha_B(m []float64) []float64 {
	result_b := make([]float64, 0, len(m))
	for _, x := range m {
		result_b = append(result_b, formula(x))
	}
	return result_b
}
