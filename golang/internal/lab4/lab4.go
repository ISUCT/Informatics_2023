package lab4

import "math"

func formula(x, a, b float64) float64 {

	return (math.Log(math.Abs(math.Pow(b, 2)-math.Pow(x, 2))) / (math.Pow(math.Abs(math.Pow(x, 2)-math.Pow(a, 2)), 0.2)))
}
func TaskA(Xn, Xa, step, a, b float64) ([]float64, []float64) {
	y := make([]float64, 0, (int((Xa-Xn)/step - 1)))
	x := make([]float64, 0, (int((Xa-Xn)/step - 1)))
	for i := Xn; i <= Xa; i += step {
		y = append(y, formula(i, a, b))
		x = append(x, i)
	}
	return x, y
}
func TaskB(array []float64, a, b float64) ([]float64, []float64) {
	y := make([]float64, 0, len(array))
	x := make([]float64, 0, cap(y))
	for _, i := range array {
		y = append(y, formula(i, a, b))
		x = append(x, i)
	}
	return x, y
}
