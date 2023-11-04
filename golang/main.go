package main

import (
	"fmt"
	"math"
)

func name() {
	fmt.Println("Жуков Даниил Евгеньевич")
}

func Radians(deg float64) float64 {
	return (deg) * (math.Pi / 180)
}

func Calculate(x, a, b float64) float64 {
	return (math.Pow(b, 3)+math.Pow(math.Sin(Radians(a*x)), 2))/math.Acos(x*b*x) + math.Pow(math.E, -x/2)
}

func main() {
	a := 1.2
	b := 0.48
	x_start := 0.7
	x_end := 2.2
	x_delta := 0.3
	fmt.Println("Задача A")
	for i := x_start; i <= x_end; i += x_delta {
		fmt.Println(Calculate(i, a, b))
	}
	var Xes = [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	for i := 0; i < len(Xes); i++ {
		fmt.Println(Calculate(Xes[i], a, b))
	}
}
