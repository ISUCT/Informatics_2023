package lab4

import (
	"math"
)

func Calc(a, b, x float64) float64 {
	return math.Pow((a+b*x)/math.Pow(math.Log10(x), 3), 1/5)
}

func Task1(a, b, xn, xk, deltax float64) []float64 {
	numElements := int((xk-xn)/deltax + 1)
	answersA := make([]float64, 0, numElements)
	for x := xn; x <= xk; x += deltax {
		y := Calc(a, b, x)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(a float64, b float64, znX []float64) []float64 {
	answersB := make([]float64, 0, len(znX))
	for _, value := range znX {
		x := value
		y := Calc(a, b, x)
		answersB = append(answersB, y)
	}
	return answersB
}
