package main

import (
	"fmt"
	"math"
)

func main() {
	//Лабораторная 2
	fmt.Println("Naumova Alena Alekseevna")

	//Лабораторная 3 - вариант 14

	//Решение под А
	a := 7.2
	b := 4.2
	x__begin := 1.81
	x__end := 5.31
	x__delta := 0.7
	fmt.Println("Решения под A:")
	for x := x__begin; x <= x__end; x = x + x__delta {
		var y1 = math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3)))
		fmt.Println(y1)
	}

	//Решение под Б
	fmt.Println("Решения под B:")
	var arguments = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for n := 0; n < 5; n++ {
		var y2 = math.Sqrt((math.Abs(a - b*arguments[n])) / (math.Pow(math.Log10(arguments[n]), 3)))
		fmt.Println(y2)
	}

}
