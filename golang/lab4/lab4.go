package lab4

import (
	"fmt"
	"math"
)

func result(a float64, b float64, x float64) float64 {
	return (math.Pow((a+b*x)/(math.Pow(math.Log10(x), 3)), 1/5))
}

func TaskA(a, b, x_start, x_end, x_delta float64) []float64 {
	answersA := make([]float64, 0, int((x_end-x_start)/x_delta))
	for x := x_start; x < x_end; x += x_delta {
		answersA = append(answersA, result(a, b, x))
	}
	return answersA
}

func TaskB(a, b float64, values []float64) []float64 {
	answersB := make([]float64, 0, len(values))
	for i := 0; i < len(values); i++ {
		answersB = append(answersB, result(a, b, values[i]))
	}
	return answersB
}

func Output(answers []float64) {
	for _, i := range answers {
		fmt.Println(i)
	}
}
