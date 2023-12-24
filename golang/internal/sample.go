package internal

import (
	"fmt"
	"math"
)

func ZadachkamathForRange(a, b, xn, xk, deltaX float64) []float64 {
	results := make([]float64, 0, int((xk-xn)/deltaX)+1)
	for x := xn; x <= xk; x += deltaX {
		result := CalculateResult(a, b, x)
		results = append(results, result)
	}
	return results
}

func ZadachkamathForValues(a, b float64, values []float64) []float64 {
	results := make([]float64, 0, len(values))
	for i, x := range values {
		result := CalculateResult(a, b, x)
		fmt.Printf("Значение функции для x%d=%f: y%d=%f\n", i+1, x, i+1, result)
		results = append(results, result)
	}
	return results
}

func CalculateResult(a, b, x float64) float64 {
	chisl := a*math.Sqrt(x) - b*math.Log(x)/math.Log(5)
	znam := math.Log(math.Abs(x-1)) / math.Log(10)

	if znam == 0 {
		return 0
	} else {
		return chisl / znam
	}
}
