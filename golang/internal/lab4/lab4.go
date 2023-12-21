package lab4

import "math"

func calc(a, x float64) float64 {
	y := math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Pow(math.Pow(x, 2)-1, 1/3)
	return y
}

func Task1(a, xn float64, xk, deltax float64) []float64 {
	var answersA []float64 = make([]float64, 0, int((xk-xn)/deltax)+1)
	for x := xn; x <= xk; x += deltax {
		y := calc(a, x)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(a float64, znachx []float64) []float64 {
	var answersB []float64 = make([]float64, 0, len(znachx))
	for _, value := range znachx {
		x := value
		y := calc(a, x)
		answersB = append(answersB, y)
	}
	return answersB
}
