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
	var numbers [5]float64 = [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	a := 1.1
	b := 0.09
	for x := 0; x <= 4; x++ {
		upf := up(numbers[x])
		downf := down(numbers[x], a, b)
		result := upf / downf
		fmt.Println("X =", result)
	}
}
