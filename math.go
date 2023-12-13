package main

import (
	"fmt"
	"math"
)

func calculateY(a, b, x, xk float64) float64 {
	numerator := math.Pow(a, 2) - math.Pow(b, 2)
	denominator := math.Log10(a / b)
	power := math.Sqrt(a * b)

	y := numerator / denominator * math.Pow(3, power)
	determinant := xk - x
	x2 := x + determinant/4
	x3 := x + determinant/2
	x4 := x + (3*determinant)/4
	x5 := xk

	fmt.Println("y =", y)
	fmt.Println("x =", x)
	fmt.Println("x2 =", x2)
	fmt.Println("x3 =", x3)
	fmt.Println("x4 =", x4)
	fmt.Println("x5 =", x5)

	return y
}

func main() {
	a := 0.4
	b := 0.8
	x := 3.2
	xk := 6.2

	result := calculateY(a, b, x, xk)
	fmt.Println(result)
	fmt.Println("Дружков Артём Михайлович")
}
