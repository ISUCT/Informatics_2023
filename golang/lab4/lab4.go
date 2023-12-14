package Lab4

import (
	"fmt"
	"math"
)

func ZadachkamathForRange() []float64 {
	a := 4.1
	b := 2.7
	xn := 1.2
	xk := 5.2
	deltaX := 0.8

	results := make([]float64, 0)
	for x := xn; x <= xk; x += deltaX {
		result := CalculateResult(a, b, x)
		results = append(results, result)
	}
	return results
}

func ZadachkamathForValues() []float64 {
	a := 4.1
	b := 2.7
	values := []float64{1.9, 2.15, 2.34, 2.73, 3.16}

	results := make([]float64, 0)
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
