package main

import (
	"fmt"
	"math"
)

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
	for x := x_begin; x <= x_end; x += x_delta {
		var y1 = (a*math.Cbrt(x) - b*(math.Log10(x))/math.Log10(5)) / (math.Pow(math.Log10(x-1), 3))
		fmt.Println(y1)
	}

	//Решение под Б
	fmt.Println("Решения под Б:")
	var list = [5]float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for n := 0; n < 5; n++ {
		var y2 = (a*math.Cbrt(list[n]) - b*(math.Log10(list[n]))/math.Log10(5)) / (math.Pow(math.Log10(list[n]-1), 3))
		fmt.Println(y2)
	}
}
