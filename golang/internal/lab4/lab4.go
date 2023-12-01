package lab4

import (
	"math"
)

func Log5(x float64) float64 {
	base := 5.0
	return (math.Log(x) / math.Log(base))
}

func form(x float64) float64 {
	a := 1.1
	b := 0.09
	return math.Log10(x*x-1) / Log5(a*math.Pow(x, 2)-b)
}
func TaksA(xn, xk, dx float64) ([]float64, []float64) {
	y := []float64{}
	x := []float64{}
	for i := xn; i <= xk; i += dx {
		x = append(x, i)
		y = append(y, form(i))
	}
	return x, y
}
func TaskB(xarray []float64) ([]float64, []float64) {
	y := []float64{}
	x := []float64{}
	for i := 0; i < len(xarray); i++ {
		x = append(x, xarray[i])
		y = append(y, form(xarray[i]))
	}
	return x, y
}
