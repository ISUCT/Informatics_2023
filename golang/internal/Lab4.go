package Lab4

import "math"

func WorkCalc(a float64, b float64, x float64) float64 {
	return (math.Pow(math.Log10(a+x), 2)) / (math.Pow((a + x), 2))
}

func Work1(a float64, b float64, xn float64, xk float64, deltaX float64) []float64 {
	sliceForWork1 := make([]float64, 0, int((xk-xn)/deltaX)+1)
	for x := xn; x <= xk; x += deltaX {
		sliceForWork1 = append(sliceForWork1, WorkCalc(a, b, x))
	}
	return sliceForWork1
}

func Work2(slice []float64, a float64, b float64) []float64 {
	result := make([]float64, len(slice))
	for i, x := range slice {
		result[i] = WorkCalc(a, b, x)
	}
	return result
}
