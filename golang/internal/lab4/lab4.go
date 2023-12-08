package lab4

import (
	"math"
)

func calcY(a float64, b float64, x float64) float64 {
	return (a*math.Sqrt(x) - b*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
}
func Task1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64 {
	var answersA = make([]float64, 0, int(math.Round(x2-x1)/deltax))
	for x := x1; x <= x2; x += deltax {
		var y float64 = calcY(a, b, x)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(a float64, b float64, xValues []float64) []float64 {
	var answersB = make([]float64, 0, len(xValues))
	for _, x := range xValues {
		var y float64 = calcY(a, b, x)
		answersB = append(answersB, y)
	}
	return answersB
}
