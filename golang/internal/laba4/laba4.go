package laba4

import (
	"math"
)

func TaskA(xn float64, xk float64, a1, b1, deltax float64) []float64 {
	var answers1 = make([]float64, 0, 6)
	for x := xn; x <= xk; x += deltax {
		var y float64 = (a1*math.Pow(x, 1/3) - b1*(math.Log(x)/math.Log(5))) / math.Pow(math.Log10(x-1), 3)
		answers1 = append(answers1, y)
	}
	return answers1
}

func TaskB(a2, b2 float64, znachX []float64) []float64 {
	var answers2 = make([]float64, 0, 5)
	for i := 0; i < len(znachX); i++ {
		var x float64 = znachX[i]
		var y float64 = (a2*math.Pow(x, 1/3) - b2*(math.Log(x)/math.Log(5))) / math.Pow(math.Log10(x-1), 3)
		answers2 = append(answers2, y)
	}
	return answers2
}
