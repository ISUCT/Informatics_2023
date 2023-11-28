package lab4

import "math"

func formula(x float64) float64 {
	return math.Asin(math.Pow(x, 2.0)) + math.Asin(math.Pow(x, 3.0))
}

func TaskA(xn, xk, dx float64) []float64 {
	y := make([]float64, 0, 5)
	for i := xn; i <= xk; i += dx {
		y = append(y, formula(i))
	}
	return y
}

func TaskB(x []float64) []float64 {
	y := make([]float64, 0, 5)
	for _, i := range x {
		y = append(y, formula(i))
	}
	return y
}
