package lab4

import (
	"math"
)
func Task1(a, b, xn, xk, deltax float64) []float64 {
	var answers1 []float64 = make([]float64, 0)
	for x := xn; x <= xk; x+= deltax {
		y := (math.Pow(math.Arccos(math.Pow(x, 2) - math.Pow(b, 2))) / math.Pow(math.Arcsin(math.Pow(x, 2) - math.Pow(a, 2))))
	}
}

func Task2(a float64, b float64, znX []float64) []float64 {
	var answers2 []float64 = make([]float64, 0)
	for _, value := range znX {
		x := value
		y := (math.Pow(math.Arccos(math.Pow(x, 2) - math.Pow(b, 2))) / math.Pow(math.Arcsin(math.Pow(x, 2) - math.Pow(a, 2))))
	}
return answers2
}