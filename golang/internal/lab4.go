package lab4test

import "math"

func Task1(a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
}

func Task2(_ []float64, a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
}
