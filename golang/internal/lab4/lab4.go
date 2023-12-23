package lab4

import "math"

func calculateResult(a, x float64) float64 {
	return (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
}

func TaskA(a, xn, xk, deltaX float64) []float64 {
	x := xn
	sliceTaskA := make([]float64, 0, int((xk-xn)/deltaX)+1)
	for x <= xk {
		sliceTaskA = append(sliceTaskA, calculateResult(a, x))
		x += deltaX
	}
	return sliceTaskA
}

func TaskB(slice []float64, a float64) []float64 {
	results := make([]float64, 0, len(slice))
	for _, x := range slice {
		results = append(results, calculateResult(a, x))
	}
	return results
}
