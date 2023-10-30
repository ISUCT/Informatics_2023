package main

import ("fmt"
"math")

func main() {
	//Лабораторная работа 2
	fmt.Println("Курушина Елизавета Андреевна")

	//Лабораторная работа 3
	a := 0.05
	b := 0.06
	x_start := 0.2
	x_end := 0.95
	x_delta := 0.15
	fmt.Println("Значение функции №10 при данных задачи A")
	for x := x_start; x <= x_end; x = x+ x_delta (
		var y1 = math.Acos(math.Pow(x,2)- math.Pow(b,2))/math.Asin(math.Pow(x,2)-math.Pow(a,2))
		fmt.Println(y1)
	)
	var h = [5]float64{0.15, 0.26, 0.37, 0.48, 0,56}
	fmt.Println("Значение функции №10 при данных задачи B")
	for i := 0; i<5; i++ (
		var y2 = math.Acos(math.Pow(x,2)- math.Pow(b,2))/math.Asin(math.Pow(x,2)-math.Pow(a,2))
		fmt.Println(y2)
	)
}
