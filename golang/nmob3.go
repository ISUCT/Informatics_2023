package main

import (
	"fmt"
	"math"
)

func Function(x float64) float64 {

	var1 := math.Log(2.0) * math.Abs(1.1*1.1-x*x)
	var2 := 5 * math.Sqrt(math.Abs(x*x-2.0*2.0))
	return var1 / var2
}

func Nmob(x_n, x_k, x_z float64) []float64 {
	y1 := []float64{}
	for x := x_n; x <= x_k; x += x_z {
		y1 = append(y1, Function(x))
	}
	return y1
}

func Nmob1(argument [5]float64) []float64 {
	y2 := []float64{}
	for i := 0; i < 5; i++ {
		y2 = append(y2, Function(argument[i]))
	}
	return y2
}

func main() {
	x_n := 0.08
	x_k := 1.08
	x_z := 0.2
	fmt.Println("Znachenie func №16 pri dannih A:", Nmob(x_n, x_k, x_z))
	var argument = [5]float64{0.1, 0.3, 0.4, 0.45, 0.65}
	fmt.Println("Znachenie func №16 pri dannih B:", Nmob1(argument))
	fmt.Println("Funcia reshena")
}
