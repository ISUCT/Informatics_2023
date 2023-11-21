package lab4

import "math"

func Formula(consta, x float64) float64 {
	return (math.Pow(consta, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt((math.Pow(x, 2) - 1)))
}

func TaskA(consta, start, end, step float64) ([]float64, []float64) {
	y := make([]float64, 0)
	x := make([]float64, 0)
	for i := start; i <= end; i += step {
		y = append(y, Formula(consta, i))
		x = append(x, i)
	}
	return x, y
}

func TaskB(consta float64, xArr []float64) ([]float64, []float64) {
	y := make([]float64, 0)
	x := make([]float64, 0)
	for _, i := range xArr {
		y = append(y, Formula(consta, i))
		x = append(x, i)
	}
	return x, y
}
