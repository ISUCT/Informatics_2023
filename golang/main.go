package main

import (
	"fmt"
	"math"
)

func CalculateFunction(x float64) float64 {
	arcsin := math.Asin(x)
	arccos := math.Acos(x)
	return math.Pow(
		math.Pow(arcsin, 2.0)+
			math.Pow(arccos, 4.0),
		3.0)
}

func CalculateFunctionInRanges(xStart, xEnd, xStep float64) []float64 {
	results := make([]float64, 0)
	for x := xStart; x <= xEnd; x += xStep {
		results = append(results, CalculateFunction(x))
	}
	return results
}
func CalculateFunctionWithXValues(xes []float64) []float64 {
	results := make([]float64, 0, len(xes))
	for _, x := range xes {
		results = append(results, CalculateFunction(x))
	}
	return results
}
func main() {
	rangedResults := CalculateFunctionInRanges(0.26, 0.66, 0.08)
	xValuesResults := CalculateFunctionWithXValues([]float64{0.1, 0.35, 0.4, 0.55, 0.6})
	fmt.Println("Кувшинов Владислав")
	fmt.Println(rangedResults)
	fmt.Println(xValuesResults)
}
