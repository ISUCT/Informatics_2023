package lab4

import (
	"math"
)

func Calc(x float64) float64 {
	return math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
}

func Primer1(xn, xk, deltaX float64) []float64 {
	var results []float64
	for xn <= xk {
		y := Calc(xn)
		results = append(results, y)
		xn += deltaX
	}
	return results
}

func Primer2(myslicefortask2 []float64) {
	for i, x := range myslicefortask2 {
		myslicefortask2[i] = Calc(x)
	}
}
