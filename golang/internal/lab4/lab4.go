package lab4

import (
	"math"
)

func CalculateY(x float64) float64 {

	return math.Asin(math.Pow(x, 2)) + math.Acos(math.Pow(x, 3))
}

// Задача под А

func A(n, k, dx float64) []float64 {
	y := make([]float64, 0, int((k-n)/dx+1))
	for i := n; i <= k; i += dx {
		y = append(y, CalculateY(i))
	}
	return y
}

// Задача под Б

func B(x []float64) []float64 {
	y := make([]float64, 0, len(x))
	for _, i := range x {
		y = append(y, CalculateY(i))
	}
	return y
}
