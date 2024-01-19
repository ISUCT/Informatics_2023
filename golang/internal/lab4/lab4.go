package lab4

import (
	"math"
)

func Task1(a, b, xn, xk, deltax float64) []float64 {
	var answers1 []float64 = make([]float64, 0, int((xk-xn)/deltax)+1)
	for x := xn; x <= xk; x += deltax {
		y := (math.Acos(math.Pow(x, 2) - math.Pow(b, 2))) / (math.Asin(math.Pow(x, 2) - math.Pow(a, 2)))
		answers1 = append(answers1, y)
	}
	return answers1
}

func Task2(a float64, b float64, znX []float64) []float64 {
	var answers2 []float64 = make([]float64, 0, len(znX))
	for _, x := range znX {
		y := (math.Acos(math.Pow(x, 2) - math.Pow(b, 2))) / (math.Asin(math.Pow(x, 2) - math.Pow(a, 2)))
		answers2 = append(answers2, y)
	}
	return answers2
}
