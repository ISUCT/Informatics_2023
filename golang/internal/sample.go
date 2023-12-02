package internal

import "math"

func CalculatePoint(x, a, b float64) float64 {
	return  (math.Pow((a + b*x), 2.5)) / (1 + (math.Log10(a + b*x)))
}

func CalculateFuncWithStep(a, b, x_n, x_k, x_z float64) []float64 {
	
	results := make([]float64, 0, int((x_k - x_n) / x_z) + 1)
	
	for x := x_n; x <= x_k; x += x_z {
		intermediateResult := CalculatePoint(x, a, b)
		results = append(results, intermediateResult)
	}

	return results
}

func CalculateFuncWithArguments(a, b float64, xes []float64) []float64 {
	results := make([]float64, 0, len(xes))

	for _, x := range xes {
		results = append(results, CalculatePoint(a, b, x))
	}

	return results
}

