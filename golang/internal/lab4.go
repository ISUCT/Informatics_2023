package lab4

import (
	"fmt"
	"math"
)

func Zad1() {
	a := 2.0
	b := 1.1
	xn := 0.08
	xk := 1.08
	deltaX := 0.2
	e := 2.71828
	x := xn
	for x <= xk {
		result := math.Log(math.Abs(b*b-x*x)) / math.Log(e)
		result = result / math.Pow(math.Abs(x*x-a*a), 1/5.0)
		fmt.Printf("x=%.2f, y=%.2f\n", x, result)
		x += deltaX
	}
}
func Zad2(values []float64) []float64 {
	result := make([]float64, len(values))
	a := 2.0
	b := 1.1
	e := 2.71828
	for i, x := range values {
		result[i] = math.Log(math.Abs(b-x)) / math.Log(e)
		result[i] = result[i] / math.Pow(math.Abs(x-a), 1/5.0)
	}
	return result
}
