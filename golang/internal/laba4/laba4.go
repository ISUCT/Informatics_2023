package laba4

import (
	"math"
)

func Calc(x float64) float64 {
	return math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1/4) + math.Pow(math.Log10(math.Pow(x, 2)), 1/3)
}
func TaskA(xn float64, xk float64, deltax float64) []float64 {
	var answersy []float64 = make([]float64, 0, int(((xk-xn)/deltax))+1)
	for x := xn; x <= xk; x += deltax {
		y := Calc(x)
		answersy = append(answersy, y)
	}
	return answersy
}

func TaskB(znX []float64) []float64 {
	var answersy []float64 = make([]float64, 0, len(znX))
	for i := 0; i < len(znX); i++ {
		x := znX[i]
		y := Calc(x)
		answersy = append(answersy, y)
	}
	return answersy
}
