package internal

import "math"

func Summ(a, b int) int {
	return a + b
}

func Formula(x float64) float64 {
	return (math.Pow((2.5 + 4.6*x), 2.5)) / (1 + (math.Log10(2.5 + 4.6*x)))
}

func TaskA(x_n, x_k, x_z float64) []float64 {
	var num_of_elements int = int((x_k) - (x_n)/x_z + 1) 
	var otvetA = make([]float64, 0, num_of_elements)
	for i := x_n; i <= x_end; i= i + x_z {
		otvetA = append(otvetA, Formula(i))
	}
	return otvetA
}

func TaskB(slice []float64) []float64 {
	var otvetB = make([]float64, 0, len(slice))
	for _, i := range slice {
		otvetB = append(otvetB, Formula(i))
	}
	return otvetB
}