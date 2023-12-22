package lab4

import (
	"math"
)

func formula(x float64) float64 {
	return (math.Pow(math.Pow(math.Asin(x), 2)+math.Pow(math.Acos(x), 4), 3))
}

func TaskA(xn, xc, dx float64) ([]float64, []float64) {
	y := make([]float64, 0, int((xc-xn)/dx-1))
	x := make([]float64, 0, cap(y))
	for i := xn; i <= xc; i += dx {
		y = append(y, formula(i))
		x = append(x, i)
	}
	return x, y
}
func TaskB(array []float64) ([]float64, []float64) {
	y := make([]float64, 0, len(array))
	x := make([]float64, 0, cap(y))
	for _, i := range array {
		y = append(y, formula(i))
		x = append(x, i)
	}
	return x, y
}
