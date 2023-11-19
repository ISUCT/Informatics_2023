package lab4

import (
	"math"
)

func Zadacha(x, a, b float64) float64 {
	return (a*math.Cbrt(x) - b*(math.Log10(x))/math.Log10(5)) / (math.Pow(math.Log10(x-1), 3))
}

func Task_A(x_begin, x_end, x_delta, a, b float64) []float64 {
	var lenmassive int = int((x_end - x_begin) / x_delta)
	output_a := make([]float64, 0, lenmassive)
	for n1 := x_begin; n1 <= x_end; n1 += x_delta {
		output_a = append(output_a, Zadacha(n1, a, b))
	}
	return output_a
}

func Task_B(list []float64, a, b float64) []float64 {
	output_b := make([]float64, 0, len(list))
	for n2 := range list {
		output_b = append(output_b, Zadacha(list[n2], a, b))
	}
	return output_b
}
