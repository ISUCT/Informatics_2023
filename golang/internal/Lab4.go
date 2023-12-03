package Lab4

import (
	"math"
)

func Calc(a, b, x float64) float64 {
	y := (math.Log2((math.Pow(b, 2))+(math.Pow(x, 2))) / (math.Pow(math.Abs((math.Pow(x, 2) - (math.Pow(a, 2)))), (1.00 / 3.00))))
	return y
}

func Part_1(a float64, b float64) []float64 {
	var r_a = make([]float64, 0)
	for x := 0.77; x < 1.77; x += 0.2 {
		y := Calc(a, b, x)
		r_a = append(r_a, y)
	}
	return r_a
}

func Part_2(a float64, b float64) []float64 {
	var xs = [5]float64{1.24, 1.38, 2.38, 3.21, 0.68}
	var r_b = make([]float64, 0)
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		y := Calc(a, b, x)
		r_b = append(r_b, y)
	}
	return r_b
}
