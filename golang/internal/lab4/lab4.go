package lab4

import (
	"math"
)

func calcY(a, b, x float64) float64 {
	return math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3)))
}

func Task1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64 {
	answersA := make([]float64, 0, int(math.Round(x2-x1)/deltax))
	for x := x1; x <= x2; x += deltax {
		var y float64 = calcY(a, b, x)
		answersA = append(answersA, y)
	}
	return answersA
}
func Task2(a float64, b float64, xValues []float64) []float64 {
	answersB := make([]float64, 0, len(xValues))
	for _, x := range xValues {
		var y float64 = calcY(a, b, x)
		answersB = append(answersB, y)
	}
	return answersB
}
