package lab4

import (
	"math"
)

func f(x float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(math.Pow(x, 2)-2.5))) + math.Cbrt(math.Log10(math.Pow(x, 2))))
}

func TaskA(xn, xk, dx float64) ([]float64, []float64) {
	x := make([]float64, 0, (int((xk-xn)/dx + 1)))
	y := make([]float64, 0, len(x))
	for i := xn; i <= xk; i += dx {
		x = append(x, i)
		y = append(y, f(i))
	}
	return x, y
}
func TaskB(array []float64) []float64 {
	y := make([]float64, 0, len(array))
	for _, i := range array {
		y = append(y, f(i))
	}
	return y
}