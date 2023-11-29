package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

func TaskA(x_begin, x_end, x_delta float64) []float64 {
	ansA := []float64{}
	for i := x_begin; i <= x_end; i = i + x_delta {
		ansA = append(ansA, Ansvers(i))
	}
	return ansA
}

func TaskB(numbers []float64) []float64 {
	ansB := make([]float64, 0, len(numbers))
	for _, value := range numbers {
		ansB = append(ansB, Ansvers(value))
	}
	return ansB
}

func Ansvers(x float64) float64 {
	chisl := math.Pow(2.5+4.6*x, 2.5)
	znam := 1 + math.Log10(2.5+4.6*x)
	y := chisl / znam
	return y
}
