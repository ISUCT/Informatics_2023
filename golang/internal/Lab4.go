package Lab4

import (
	"math"
)

func Calc(a, b, x float64) float64 {
	y := (math.Log2((math.Pow(b, 2))+(math.Pow(x, 2))) / (math.Pow(math.Abs((math.Pow(x, 2) - (math.Pow(a, 2)))), (1.00 / 3.00))))
	return y
}

func Part_1(a float64, b float64, xn float64, xk float64, dx float64) []float64 {
	size := int((xk-xn)/dx) + 1
	var r_a = make([]float64, 0, size)
	for x := xn; x <= xk; x += dx {
		y := Calc(a, b, x)
		r_a = append(r_a, y)
	}
	return r_a
}

func Part_2(slice []float64, a float64, b float64) []float64 {
	var r_b = make([]float64, 0)
	for i := 0; i < len(slice); i++ {
		x := slice[i]
		y := Calc(a, b, x)
		r_b = append(r_b, y)
	}
	return r_b
}
