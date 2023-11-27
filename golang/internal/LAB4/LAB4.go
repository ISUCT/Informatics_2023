package lab4

import (
	"math"
)

func formulausl(x float64) float64 {
	var a float64 = -2.5
	var b float64 = 3.4
	return math.Pow(a+b*x, 3.5)/1.8 + math.Pow(math.Cos(a*x), 3)
}
func formulamain(x float64) float64 {
	var a float64 = -2.5
	if formulausl(x) >= 5 {
		return (math.Pow(math.Log10(math.Pow(a, 2)+x), 2) / math.Pow(a+x, 2))
	} else {
		return (0)
	}
}
func TaskA(xNach, xKon, shag float64) ([]float64, []float64) {
	y := []float64{}
	x := []float64{}
	for i := xNach; i <= xKon; i += shag {
		x = append(x, i)
		y = append(y, formulamain(i))
	}
	return x, y
}
func TaskB(xArr []float64) ([]float64, []float64) {
	y := []float64{}
	x := []float64{}
	for i := 0; i < len(xArr); i++ {
		x = append(x, xArr[i])
		y = append(y, formulamain(xArr[i]))
	}
	return x, y
}
