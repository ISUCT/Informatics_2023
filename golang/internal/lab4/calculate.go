package lab4

import (
	"math"
)

func Calc(a, b, x float64) float64 {
	return math.Pow(((a + b*x) / (math.Pow((math.Log10(x)), 3))), 1.0/5.0)
}

func Task_A(a, b, xn, xk, dx float64) []float64 {
	y1 := []float64{}
	for x := xn; x <= xk; x += dx {
		y1 = append(y1, Calc(a, b, x))
	}
	return y1
}

func Task_B(a, b float64, vhodnie_x []float64) []float64 {
	y2 := []float64{}
	for _, x := range vhodnie_x {
		y2 = append(y2, Calc(a, b, x))
	}
	return y2
}
