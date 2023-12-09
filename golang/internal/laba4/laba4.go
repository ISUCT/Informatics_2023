package laba4

import (
	"math"
)

func Calc(x float64) float64 {
	return math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), 1/7)
}

func Task1(xn, xk, deltax float64) []float64 {
	var answersA []float64 = make([]float64, int((xk-xk)/deltax)+1)
	for x := xn; x <= xk; xn += deltax {
		y := Calc(x)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(znX []float64) []float64 {
	var answersB []float64 = make([]float64, len(znX))
	for _, value := range znX {
		x := value
		y := Calc(x)
		answersB = append(answersB, y)
	}
	return answersB
}
