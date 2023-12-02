package lab4test

import "math"

func TaskA(a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
}

func Task1Slice(a float64, b float64, xn float64, xk float64, deltaX float64) []float64 {
	sliceForTask1 := make([]float64, 0)
	for x := xn; x <= xk; x += deltaX {
		sliceForTask1 = append(sliceForTask1, TaskA(a, b, x))
	}
	return sliceForTask1
}

func TaskB(slice []float64, a float64, b float64) []float64 {
	result := make([]float64, len(slice))
	for i, x := range slice {
		result[i] = (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
	}
	return result
}
