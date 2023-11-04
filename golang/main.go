package main

import "fmt"
import (
	"fmt"
	"math"
)

func ToRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func Calculate(a float64, b float64, x float64) float64 {
	var exp float64
	if x < 5 {
		exp = math.Pow(math.Log10(math.Pow(a, 2)+x), 2) / math.Pow(a+x, 2)
	} else {
		exp = math.Pow(a+b*x, 2.5) / (1.8 + math.Pow(math.Cos(ToRadians(a*x)), 3))
	}

	return exp
}

func main() {
	fmt.Println("Горюнов Александр Алексеевич")

	var a = -2.5
	var b = 3.4
	var begX = 3.5
	var endX = 6.5
	var detX = 0.6

	for x := begX; x <= endX; x += detX {
		fmt.Println(Calculate(a, b, x))
	}

	var xList = []float64{2.89, 3, 54, 5.21, 6.28, 3.48}

	for i := 0; i < len(xList); i++ {
		fmt.Println(Calculate(a, b, xList[i]))
	}
}
