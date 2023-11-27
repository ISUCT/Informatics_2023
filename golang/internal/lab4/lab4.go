package lab4

import (
	"math"
)

func formula(x float64) float64 {
	var a float64 = 1.2
	var b float64 = 0.48
	return ((math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos((x * b * x)) + math.Exp(-x/2)))
}

func TaskA(xN, xK, xD float64) ([]float64, []float64) {
	x := []float64{}
	y := []float64{}
	for i := xN; i <= xK; i += xD {
		y = append(y, formula(i))
		x = append(x, i)
	}
	return x, y
}

func TaskB(xArray []float64) ([]float64, []float64) {
	x := []float64{}
	y := []float64{}
	for i := 0; i < len(xArray); i++ {
		y = append(y, formula(xArray[i]))
		x = append(x, xArray[i])
	}
	return x, y
}
