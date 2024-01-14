package lab4

import (
	"fmt"
	"math"
)

func result(a, b, x float64) float64 {
	if x > 5 {
		result := ((math.Pow(math.Log10(math.Pow(a, 2)+x), 2)) / math.Pow(a+x, 2))
		return result
	}
	result := (math.Pow(a+b*x, 2.5)) / (1.8 + math.Pow(math.Cos(a*x), 3))
	if result >= 5 {
		return result
	}
	return -1
}

func TaskA(a, b, x_begin, x_end, x_delta float64) []float64 {
	answersA := make([]float64, 0, int((x_end-x_begin)/x_delta))
	for x := x_begin; x <= x_end; x += x_delta {
		answersA = append(answersA, result(a, b, x))
	}
	return answersA
}

func TaskB(a, b float64, values []float64) []float64 {
	answersB := make([]float64, 0, len(values))
	for _, x := range values {
		answersB = append(answersB, result(a, b, x))
	}
	return answersB
}

func Output(values []float64) {
	for _, i := range values {
		fmt.Println(i)
	}
}

func Solve(a, b, x_begin, x_end, x_delta float64) {
	results := TaskA(a, b, x_begin, x_end, x_delta)
	fmt.Println("Solutions to task A:")
	Output(results)

	results = TaskB(a, b, []float64{2.89, 3.54, 5.21, 6.28, 3.48})
	fmt.Println("Solutions to task B:")
	Output(results)
}
