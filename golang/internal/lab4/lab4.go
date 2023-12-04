package lab4

import (
	"math"
)

func Log5(x float64) float64 {
	base := 5.0
	return (math.Log(x) / math.Log(base))
}
func form(x, a, b float64) float64 {
	return math.Log10(math.Pow(x, 2)-1) / Log5(a*math.Pow(x, 2)-b)
}
func TaksA(xn, xk, dx, a, b float64) ([]float64, []float64) {
	y := make([]float64, 0, (int((xk-xn)/dx - 1)))
	x := make([]float64, 0, (int((xk-xn)/dx - 1)))
	for i := xn; i <= xk; i += dx {
		x = append(x, i)
		y = append(y, form(i, a, b))
	}
	return x, y
}
func TaskB(xarray []float64, a, b float64) ([]float64, []float64) {
	x := make([]float64, 0, len(xarray))
	y := make([]float64, 0, len(xarray))
	for i := 0; i < len(xarray); i++ {
		x = append(x, xarray[i])
		y = append(y, form(xarray[i], a, b))
	}
	return x, y
}
