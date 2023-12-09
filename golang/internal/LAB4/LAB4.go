package lab4

import (
	"math"
)

func formulamain(x, a, b float64) float64 {
	if x <= 5 {
		return math.Pow(a+b*x, 3.5)/1.8 + math.Pow(math.Cos(a*x), 3)
	} else {
		return (math.Pow(math.Log10(math.Pow(a, 2)+x), 2) / math.Pow(a+x, 2))
	}
}
func TaskA(xNach, xKon, shag, a, b float64) ([]float64, []float64) {
	y := make([]float64, 0, int((xKon-xNach)/shag+1))
	x := make([]float64, 0, int((xKon-xNach)/shag+1))
	for i := xNach; i <= xKon; i += shag {
		x = append(x, i)
		y = append(y, formulamain(i, a, b))
	}
	return x, y
}
func TaskB(xArr []float64, a, b float64) ([]float64, []float64) {
	y := make([]float64, 0, len(xArr))
	x := make([]float64, 0, len(xArr))
	for i := 0; i < len(xArr); i++ {
		x = append(x, xArr[i])
		y = append(y, formulamain(xArr[i], a, b))
	}
	return x, y
}
