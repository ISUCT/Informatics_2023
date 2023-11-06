package main

import (
	"fmt"
	"math"
)

func result1(a float64, x float64) float64 {
	result := ((math.Pow(math.Log10(math.Pow(a, 2)+x), 2)) / math.Pow(a+x, 2))
	if x > 5 {
		return result
	}
	return -1
}

func result2(a float64, b float64, x float64) float64 {
	result := ((math.Pow(a+b*x, 2.5)) / (1.8 + math.Pow(math.Cos(a*x), 3)))
	if result >= 5 {
		return result
	}
	return -1
}

func main() {
	fmt.Println("Журавлев Сергей Викторович")

	a := -2.4
	b := 3.4
	x_begin := 3.5
	x_end := 6.5
	x_delta := 0.6

	fmt.Println("Решения задачи под А:")
	for x := x_begin; x <= x_end; x += x_delta {
		fmt.Println(result1(a, x), result2(a, b, x))
	}

	fmt.Println("Решения задачи под В:")
	var values = [5]float64{2.89, 3.54, 5.21, 6.28, 3.48}
	for i := 0; i < len(values); i++ {
		fmt.Println(result1(a, values[i]), result2(a, b, values[i]))
	}
}
