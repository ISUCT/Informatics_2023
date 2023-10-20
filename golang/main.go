package main

import (
	"fmt"
	"math"
)

func radians(z float64) float64 {
	return z * (math.Pi / 180.0)
}

func formula(x float64) float64 {
	return ((math.Pow(math.Sin(radians(x)), 3) + math.Pow(math.Cos(radians(x)), 3)) * math.Log(x))
}
func main() {
	fmt.Println("Хосровян Александр Арменович")

	x_begin := 0.11
	x_end := 0.36
	x_delta := 0.05

	fmt.Println("Решения под А:")
	for i := x_begin; i <= x_end; i += x_delta {
		fmt.Println(formula(i))
	}

	fmt.Println("Решения под B:")
	var list = [5]float64{0.2, 0.3, 0.38, 0.43, 0.57}

	for i := 0; i < len(list); i++ {
		x := list[i]
		fmt.Println(formula(x))
	}
}
