package lab4

import "math"

func Form(x, a float64) float64 {
	return (math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt((math.Pow(x, 2) - 1)))
}

func TaskA(Xs, Xe, step, a float64) ([]float64, []float64) {
	y := make([]float64, 0, (int((Xe-Xs)/step - 1)))
	x := make([]float64, 0, (int((Xe-Xs)/step - 1)))
	for i := Xs; i <= Xe; i += step {
		y = append(y, Form(i, a))
		x = append(x, i)
	}
	return x, y
}

func TaskB(a float64, xArr []float64) ([]float64, []float64) {
	x := make([]float64, 0, len(xArr))
	y := make([]float64, 0, len(xArr))
	for _, i := range xArr {
		y = append(y, Form(i, a))
		x = append(x, i)
	}
	return x, y
}
