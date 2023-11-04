package main

import (
	"fmt"
	"math"
)

func ToRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func Calculate(x float64) float64 {
	var exp float64
	exp = (math.Pow(math.Sin(ToRadians(x)), 3) + math.Pow(math.Cos(ToRadians(x)), 3)) * math.Log10(x)
	return exp
}

func main() {
	fmt.Println("Соколова Арина Михайловна")
	var begX = 0.11
	var endX = 0.36
	var detX = 0.05

	for x := begX; x <= endX; x += detX {
		fmt.Println(Calculate(x))
	}

	var xList = []float64{0.2, 0.3, 0.38, 0.43, 0.57}

	for i := 0; i < len(xList); i++ {
		fmt.Println(Calculate(xList[i]))
	}
}
