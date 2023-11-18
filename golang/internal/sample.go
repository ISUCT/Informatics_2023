package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

func log5(x float64) float64 {
	return math.Log(x) / math.Log(5)
}

func Task_A(x_start, x_end, x_delta float64) []float64 {
	var otvet_a = []float64{}
	for i := x_start; i <= x_end; i += x_delta {
		otvet_a = append(otvet_a, function(i))
	}
	return otvet_a
}

func Task_B(xs []float64) []float64 {
	var otvet_b = []float64{}
	for _, i := range xs {
		otvet_b = append(otvet_b, function(i))
	}
	return otvet_b
}

func function(x float64) float64 {
	const a float64 = 4.1
	const b float64 = 2.7
	return ((a * math.Sqrt(x)) - (b * log5(x))) / (math.Log10(math.Abs(x - 1)))
}