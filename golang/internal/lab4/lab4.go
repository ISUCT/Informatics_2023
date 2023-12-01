package lab4

import (
	"math"
)

func Task1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64 {
	var answersA = make([]float64, 0, int(math.Round(x2-x1)/deltax))
	for x := x1; x <= x2; x += deltax {
		var y float64 = (a*math.Sqrt(x) - b*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(a float64, b float64, x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) []float64 {
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = make([]float64, 0, len(znachX))
	for i := 0; i < len(znachX); i++ {
		x := znachX[i]
		var y float64 = (a*math.Sqrt(x) - b*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
		answersB = append(answersB, y)
	}
	return answersB
}
