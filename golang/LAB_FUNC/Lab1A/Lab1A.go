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
	for x := 1.2; x <= 2.2; x += 0.2 {
		a := 1.1
		b := 0.09
		upf := up(x)
		downf := down(x, a, b)
		result := upf / downf
		fmt.Println("X =", result)
	}
}
