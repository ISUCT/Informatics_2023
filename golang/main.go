package main

import (
	"fmt"
	"math"
)

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

func result(a float64, x float64, b float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / Log(5, (a*(math.Pow(x, 2)))-b))
}

func main() {
	fmt.Println("Коротков Данила Максимович")

	a := 1.1
	b := 0.09
	x_begin := 1.2
	x_end := 2.2
	x_delta := 0.2

	fmt.Println("Решения задачи под А:")
	for x := x_begin; x <= x_end; x += x_delta {
		fmt.Println(result(a, x, b))
	}

	fmt.Println("Решения задачи под В:")
	var values = [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	for i := 0; i < len(values); i++ {
		fmt.Println(result(a, values[i], b))
	}
}
