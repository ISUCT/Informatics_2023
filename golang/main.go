package main

import (
	"fmt"
	"math"
)

func Soch(a, b int) int {
	return a + b
}

func Equation(x float64) float64 {
	a := 0.4
	b := 0.8
	return (math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a/b) * math.Pow(a*b, 1/3.0))
}

func CalcA(x_begin, x_end, x_delta float64) []float64 {
	var num_of_elements int = int((x_end) - (x_begin)/x_delta + 1)
	var ResultA = make([]float64, 0, num_of_elements)
	for i := x_begin; i <= x_end; i = i + x_delta {
		ResultA = append(ResultA, Equation(i))
	}
	return ResultA
}

func CalcB(slice []float64) []float64 {
	var ResultB = make([]float64, 0, len(slice))
	for _, i := range slice {
		ResultB = append(ResultB, Equation(i))
	}
	return ResultB
}

func main() {
	fmt.Println(CalcA(3.2, 6.2, 0.6))
	fmt.Println(CalcB([]float64{4.48, 3.56, 2.78, 5.28, 3.21}))
}
