package main

import (
	"fmt"
	"math"
)

func ToRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func task(x float64, b float64) float64 {
	return (1 + math.Pow(math.Sin(ToRadians(math.Pow(b, 3)+math.Pow(x, 3))), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
}

func main() {
	fmt.Println("Антонов Кирилл Юрьевич")
	var b = 2.5
	var begX = 1.28
	var finX = 3.28
	var detX = 0.4

	for i := begX; i < finX; i += detX {
		fmt.Println(task(i, b))
	}

	var xList = [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}

	for i := 0; i < len(xList); i++ {
		fmt.Println(task(xList[i], b))
	}
}
