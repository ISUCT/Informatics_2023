package main

import (
	"fmt"
	"math"
)

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

func y(a, b, x float64) float64 {
	return (a*math.Cbrt(x) - b*Log(5, x)) / math.Pow(math.Log10(x-1), 3)
}

func main() {
	fmt.Println("Чурсина Софья 1/278")

	a := 4.1
	b := 2.7

	fmt.Println("Задача A")
	for x := 1.5; x <= 3.5; x += 0.4 {
		res := y(a, b, x)
		fmt.Println(res)
	}

	fmt.Println("Задача B")
	array := [5]float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for _, v := range array {
		res := y(a, b, v)
		fmt.Println(res)
	}
}
