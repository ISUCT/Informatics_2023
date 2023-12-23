package lab4

import (
	"math"
)

func f(x float64) float64 {
	return math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1.0/4.0) + math.Pow(math.Log10(math.Pow(x, 2)), 1.0/3.0)
}
func TaskA(start, end, x float64) []float64 {
	var answers []float64 = make([]float64, 0, int((end-start)/x))
	for i := start; i < end; i += x {
		answers = append(answers, f(i))
	}
	return answers
}
func TaskB(numbers [5]float64) []float64 {
	var answers []float64 = make([]float64, 0, len(numbers))
	for i := 0; i < len(numbers); i++ {
		answers = append(answers, f(numbers[i]))
	}
	return answers
}
