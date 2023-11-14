package main

import (
	"fmt"
	"math"
)

func up(x float64) float64 {
	return math.Log(x*x*x - 1)
}

func down(x, a, b float64) float64 {
	arg := a*x*x - b
	base := 5.0
	return math.Log(arg) / math.Log(base)
}

func main() {
	fmt.Println("Задача A:")
	a := 1.1
	b := 0.09
	for x := 1.2; x <= 2.2; x += 0.2 {
		upf := up(x)
		downf := down(x, a, b)
		result := upf / downf
		fmt.Println("X =", result)
	}
	fmt.Println("\n Задача B:")
	var numbers [5]float64 = [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	for x := 0; x <= 4; x++ {
		upf := up(numbers[x])
		downf := down(numbers[x], a, b)
		result := upf / downf
		fmt.Println("X =", result)
	}
}
