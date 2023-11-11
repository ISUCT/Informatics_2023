package main

import (
	"fmt"
	"math"
)

// Лабораторная работа №2
func main() {
	fmt.Println("Иванова Анастасия Евгеньевна")
	//Лабораторная работа № 3
	a := 0.4
	b := 0.8
	x_begin := 3.2
	x_end := 6.2
	x_delta := 0.6
	fmt.Println("Значение функции №7 при данных задачи А =")
	for x := x_begin; x <= x_end; x = x + x_delta {
		var y1 = (math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a/b) * math.Pow(a*b, 1/3.0))
		fmt.Println(y1)
	}
	var argument = [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	fmt.Println("Значение функции №7 при данных задачи Б =")
	for i := 0; i < 5; i++ {
		var y2 = (math.Pow(a, argument[i]) - math.Pow(b, argument[i])) / (math.Log10(a/b) * math.Pow(a*b, 1/3.0))
		fmt.Println(y2)
	}
}
