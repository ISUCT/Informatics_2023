package internal

import (
	"math"
)

func CalculateFuncForOnePoint(x, a, b float64) float64 {
	return (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(math.E, x/a))
}

func CalculateFuncForSteps(a, b, xStart, xEnd, xStep float64) []float64 {
	results := make([]float64, 0, int((xEnd-xStart)/xStep)+1)

	for x := xStart; x <= xEnd; x += xStep {
		results = append(results, CalculateFuncForOnePoint(x, a, b))
	}

	return results
}

func CalculateFuncForXes(a, b float64, xes []float64) []float64 {
	results := make([]float64, 0, len(xes))

	for _, x := range xes {
		results = append(results, CalculateFuncForOnePoint(x, a, b))
	}

	return results
}
