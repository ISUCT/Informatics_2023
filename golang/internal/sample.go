package internal

import (
	"math"
)

func formul(a, b, x float64) float64 {
	return ((math.Pow(a, x)) + (math.Pow(b, x))) / (math.Log(a/b) / 10.0) * math.Sqrt(a*b)
}

func TaskA(a, b, xn, xk, deltaX float64) ([]float64, []float64) {
	var predx = int(math.Ceil((xk-xn)/deltaX + 1))
	var xlist = make([]float64, 0, predx)
	var ylist = make([]float64, 0, predx)
	for x := xn; x <= xk; x += deltaX {
		ylist = append(ylist, formul(a, b, x))
		xlist = append(xlist, x)
	}
	return ylist, xlist
}
func TaskB(xlist []float64, a, b float64) []float64 {
	var ylist = make([]float64, 0, len(xlist))
	for _, v := range xlist {
		ylist = append(ylist, formul(a, b, v))
	}
	return ylist
}
