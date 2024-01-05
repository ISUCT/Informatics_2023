package lab4

import (
	"fmt"
	"math"
)

func result1(a float64, x float64) float64 {
	result := ((math.Pow(math.Log10(math.Pow(a, 2)+x), 2)) / math.Pow(a+x, 2))
	if x > 5 {
		return result
	}
	return -1
}

func result2(a float64, b float64, x float64) float64 {
	result := ((math.Pow(a+b*x, 2.5)) / (1.8 + math.Pow(math.Cos(a*x), 3)))
	if result >= 5 {
		return result
	}
	return -1
}

func TaskA(a, b, x_begin, x_end, x_delta float64) []float64 {
	answersA := make([]float64, 0, int((x_end-x_begin)/x_delta))
	for x := x_begin; x <= x_end; x += x_delta {
		answersA = append(answersA, result1(a, x))
		answersA = append(answersA, result2(a, b, x))
	}
	return answersA
}

func TaskB(a, b float64, values []float64) []float64 {
	answersB := make([]float64, 0, len(values))
	for _, x := range values {
		answersB = append(answersB, result1(a, x))
		answersB = append(answersB, result2(a, b, x))
	}
	return answersB
}

func Output(values []float64) {
	for _, i := range values {
		fmt.Println(i)
	}
}
