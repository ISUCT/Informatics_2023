package laba4

import (
	"math"
)

func calc(a, b, x float64) float64 {
	return (a*math.Pow(x, 1/3) - b*(math.Log(x)/math.Log(5))) / math.Pow(math.Log10(x-1), 3)
}

func TaskA(xn float64, xk float64, a1, b1, deltax float64) []float64 {
	var answers1 = make([]float64, 0, int((xk-xn)/deltax)+1)
	for x := xn; x <= xk; x += deltax {
		y := calc(a1, b1, x)
		answers1 = append(answers1, y)
	}
	return answers1
}

func TaskB(a2, b2 float64, znachX []float64) []float64 {
	var answers2 = make([]float64, 0, len(znachX))
	for _, x := range znachX {
		y := calc(a2, b2, x)
		answers2 = append(answers2, y)
	}
	return answers2
}
