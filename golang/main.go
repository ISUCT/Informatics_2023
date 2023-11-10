package main

import (
	"fmt"
	"math"
)

func result(a float64, b float64, x float64) float64 {
	return (math.Pow((a+b*x)/(math.Pow(math.Log10(x), 3)), 1/5))
}

func main() {
	fmt.Println("Спиридонова Полина Алексеевна")

	a := 7.2
	b := 1.3
	x_start := 1.56
	x_end := 4.71
	x_delta := 0.63

	fmt.Println("Решение задачи под буквой А")
	for x := x_start; x < x_end; x += x_delta {
		fmt.Println(result(a, b, x))
	}

	fmt.Println("Решение задачи под буквой B")
	var values = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for i := 0; i < len(values); i++ {
		fmt.Println(result(a, b, values[i]))
	}
}
