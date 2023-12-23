package lab4

import "math"

func TaskA(a, b, xn, xk, deltax, e float64) []float64 {
	results := make([]float64, 0)
	x := xn
	for x <= xk {
		result := (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(e, x/a))
		results = append(results, result)
		x += deltax
	}
	return results
}

func TaskB(a, b, e float64, values ...float64) []float64 {
	results := make([]float64, 0)
	for _, x := range values {
		result := (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(e, x/a))
		results = append(results, result)
	}
	return results
}
