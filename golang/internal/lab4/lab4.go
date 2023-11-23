package lab4

import (
	"math"
)

func znachenie_funkcii(x float64) float64 {
	x = math.Round(math.Abs(x)*100) / 100
	if x < 1 {
		return math.Acos(x)
	}
	return math.Pow(1.2, x) - math.Pow(x, 1.2)
}
func TaskA(xn, xk, dx float64) ([]float64, []float64) {
	var list_of_x = []float64{}
	var list_of_y = []float64{}
	for i := xn; i <= xk; i += dx {
		i := math.Round(math.Abs(i)*100) / 100
		y := znachenie_funkcii(i)
		list_of_x = append(list_of_x, i)
		list_of_y = append(list_of_y, y)
	}
	return list_of_x, list_of_y
}
func TaskB(a []float64) []float64 {
	var list_of_by = []float64{}
	var i int
	for i < len(a) {
		y := znachenie_funkcii(a[i])
		list_of_by = append(list_of_by, y)
		i++
	}
	return list_of_by
}
