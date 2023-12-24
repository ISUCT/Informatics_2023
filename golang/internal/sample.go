package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func funkt16(a, b, c float64) float64 {
	return (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(c, 2)))) / (math.Pow(math.Abs(math.Pow(c, 2)-math.Pow(a, 2)), 0.2))
}

func TaskA(x_begin, x_end, x_delta, a, b float64) []float64 {
	var num_of_elements int = int((x_end) - (x_begin)/x_delta + 1)
	var answer_A = make([]float64, 0, num_of_elements)
	for i := x_begin; i <= x_end; i += x_delta {
		answer_A = append(answer_A, funkt16(a, b, i))
	}
	return answer_A
}

func TaskB(slice []float64, a, b float64) []float64 {
	var answer_B = make([]float64, 0, len(slice))
	for _, i := range slice {
		answer_B = append(answer_B, funkt16(a, b, i))
	}
	return answer_B
}
