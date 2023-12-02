package lab4

import (
	"fmt"
	"math"
)

func Log(base, j float64) float64 {
	return math.Log(j) / math.Log(base)
}

func result(a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / Log(5, (a*(math.Pow(x, 2)))-b))
}

func taskA(a, b, x_begin, x_end, x_delta float64) []float64 {
	otvetA := make([]float64, 0, int((x_end-x_begin)/x_delta))
	for x := x_begin; x <= x_end; x += x_delta {
		otvetA = append(otvetA, result(a, b, x))
	}
	return otvetA
}

func taskB(a float64, b float64, values []float64) []float64 {
	otvetB := make([]float64, 0, len(values))
	for _, x := range values {
		otvetB = append(otvetB, result(a, b, x))
	}
	return otvetB
}

func output(results []float64) {
	for _, i := range results {
		fmt.Println(i)
	}
	fmt.Println("===================================================")
}

func Solve(a, b, x_begin, x_end, x_delta float64) {
	results := taskA(a, b, x_begin, x_end, x_delta)
	fmt.Println("Решения задачи под А:")
	output(results)

	results = taskB(a, b, []float64{1.21, 1.76, 2.53, 3.48, 4.52})
	fmt.Println("Решения задачи под В:")
	output(results)
}
