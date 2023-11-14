package main

import (
	"fmt"
	"math"
)

func zadacha(x float64, a float64, b float64) float64 {
	return (math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3))))
}

func main() {
	//Лабораторная 2
	fmt.Println("Naumova Alena Alekseevna")

	//Лабораторная 3 - вариант 14

	//Решение под А
	a := 7.2
	b := 4.2
	xBegin := 1.81
	xEnd := 5.31
	xDelta := 0.7
	fmt.Println("Решения под A:")
	for y := xBegin; y <= xEnd; y = y + xDelta {
		fmt.Println(zadacha(y, a, b))
	}

	//Решение под Б
	fmt.Println("Решения под B:")
	var arguments = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for n := 0; n < len(arguments); n++ {
		c := arguments[n]
		fmt.Println(zadacha(c, a, b))
	}

}
