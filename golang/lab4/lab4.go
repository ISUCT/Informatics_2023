package lab4

import (
	"fmt"
	"math"
)

func result(a, b, x float64) float64 {
	return ((math.Cbrt(a*x - b)) / (math.Pow(math.Log10(x), 2)))
}

func Task_A(a, b, x_start, x_end, x_delta float64) []float64 {
	y := make([]float64, 0, int((x_end-x_start)/x_delta))
	for x := x_start; x <= x_end; x += x_delta {
		y = append(y, result(a, b, x))
	}
	return y
}

func Task_B(a, b float64, values []float64) []float64 {
	y := make([]float64, 0, len(values))
	for _, x := range values {
		y = append(y, result(a, b, x))
	}
	return y
}

func output(results []float64) {
	for _, i := range results {
		fmt.Println(i)
	}
	fmt.Println("               ")
}

func Solution(a, b, x_start, x_end, x_delta float64) {
	results := Task_A(a, b, x_start, x_end, x_delta)
	fmt.Println("Solutions A")
	output(results)

	results = Task_B(a, b, []float64{0.35, 1.28, 3.51, 5.21, 4.16})
	fmt.Println("Solutions B")
	output(results)
}
