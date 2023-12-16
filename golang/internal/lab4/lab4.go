package lab4

import (
	"math"
)

func Calc(x float64) float64 {
	return math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
}

func Primer1(xn, xk, deltaX float64) []float64 {
	results := make([]float64, 0, int((xk-xn)/deltaX)+1)
	for xn <= xk {
		y := Calc(xn)
		results = append(results, y)
		xn += deltaX
	}
	return results
}

func Primer2(slice []float64) []float64 {
	myslicefortask2 := make([]float64, 0, len(slice))
	for i, x := range slice {
		myslicefortask2[i] = Calc(x)
	}
	return myslicefortask2
}
