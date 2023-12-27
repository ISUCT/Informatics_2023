package internal 
import "math"

func Summ(a, b int) int {
	return a + b
}
func Formula(a, b, x float64) float64 {
	return (math.Pow((a + b*x), 2.5)) / (1 + (math.Log10(a + b*x)))
}
func TaskA(x_n, x_k, x_z, a, b float64) []float64 {
	var num_of_elements int = int((x_k) - (x_n)/x_z + 1)
	var otvetA = make([]float64, 0, num_of_elements)
	for i := x_n; i <= x_k; i += x_z {
		otvetA = append(otvetA, Formula(a, b, i))
	}
	return otvetA
}
func TaskB(slice []float64, a, b float64) []float64 {
	var otvetB = make([]float64, 0, len(slice))
	for i := range slice {
		otvetB = append(otvetB, Formula(a, b, slice[i]))
	}
	return otvetB
}