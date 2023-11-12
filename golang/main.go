package main

import (
	"fmt"
	"math"
)

func main() {
	// Лабораторная 2
	fmt.Println("Байгулова Ева Валерьевна")

	// Лабораторная 3 - Вариант 3
	a := 2.0
	b := 0.95
	x_begin := 1.25
	x_end := 2.75
	x_delta := 0.3
	fmt.Println("Значение функции №3 при данных задачи А =")
	for x := x_begin; x <= x_end; x = x + x_delta {
		var y1 = (1 + math.Pow(math.Log10(x/a), 2)) / (b - (math.Exp(x / a)))
		fmt.Println(y1)
	}
	var argument = [5]float64{2.2, 3.78, 4.51, 6.58, 1.2} // Массив значений x1-5
	fmt.Println("Значение функции №3 при данных задачи B =")
	for i := 0; i < 5; i++ {
		var y2 = (1 + math.Pow(math.Log10(argument[i]/a), 2)) / (b - (math.Exp(argument[i] / a)))
		fmt.Println(y2)
	}
}
