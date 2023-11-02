package main

import (
	"fmt"
	"math"
)

func primer(x float64, a float64, b float64) float64 {
	return (a*math.Cbrt(x) - b*(math.Log10(x))/math.Log10(5)) / (math.Pow(math.Log10(x-1), 3))
}

func main() {
	
	//Лабораторная работа 2
	fmt.Println("Chalunin Artem Romanovich")

	//Лабораторная работа номер 3 - варинант 23

	//Решение под А
	fmt.Println("Решения под A:")
	a := 4.1
	b := 2.7
	x_begin := 1.5
	x_end := 3.5
	x_delta := 0.4
	for n1 := x_begin; n1 <= x_end; n1 += x_delta {
		fmt.Println(primer(n1, a, b))
	}

	//Решение под Б
	fmt.Println("Решения под Б:")
	var list = [5]float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for n2 := 0; n2 < len(list); n2++ {
		c := list[n2]
		fmt.Println(primer(c, a, b))
	}
}
