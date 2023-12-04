package internal

import (
	"math"
)

const a, b = 2.0, 1.1

func Summ(a, b int) int {
	return a + b
}

var x_begin float64 = 0.08
var x_end float64 = 1.08
var x_delta float64 = 0.2

func funkt16(c float64) float64 { 
	return (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(c, 2)))) / (math.Pow(math.Abs(math.Pow(c, 2)-math.Pow(a, 2)), 0.2))
}

func TaskA(x_begin, x_end, x_delta float64) []float64 {
	var num_of_elements int = int((x_end) - (x_begin)/x_delta + 1)
	var answer_A = make([]float64, 0, num_of_elements)
	for i := x_begin; i <= x_end; i = i + x_delta {
		answer_A = append(answer_A, funkt16(i))
	}
	return answer_A
}

func TaskB(slice []float64) []float64 {
	var answer_B = make([]float64, 0, len(slice))
	for _, i := range slice {
		answer_B = append(answer_B, funkt16(i))
	}
	return answer_B
}
