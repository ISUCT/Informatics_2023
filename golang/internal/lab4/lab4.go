package lab4

import "math"

func calculateResult(a, x float64) float64 {
	return (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
}

func TaskA(a, xn, xk, deltaX float64) []float64 {
	x := xn
	results := []float64{}
	for x <= xk {
		result := calculateResult(a, x)
		results = append(results, result)
		x += deltaX
	}
	return results
}

func TaskB(a, x1, x2, x3, x4, x5 float64) []float64 {
	values := []float64{x1, x2, x3, x4, x5}
	results := []float64{}
	for _, x := range values {
		result := calculateResult(a, x)
		results = append(results, result)
	}
	return results
}
