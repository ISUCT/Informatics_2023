package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}
func log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

func func_res(x float64) float64 {
	return (4.1*math.Cbrt(x) - 2.7*log(5, x)) / math.Pow(math.Log10(x-1), 3)
}

func Task_A(x_start, x_end, x_step float64) []float64 {
	var answer_a = []float64{}
	for i := x_start; i <= x_end; i += x_step {
		answer_a = append(answer_a, func_res(i))
	}
	return answer_a
}

func Task_B(y []float64) []float64 {
	var answer_b = []float64{}
	for _, i := range y {
		answer_b = append(answer_b, func_res(i))
	}
	return answer_b
}
