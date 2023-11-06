package main

import (
	"fmt"
	"math"
)

func result(a float64, b float64, x float64) float64 {
	return ((math.Cbrt(a*x-b)) / (math.Pow(math.Tan(x), 2)))
}

func main() {
	fmt.Println("Голосова Анастасия Сергеевна")

	a:= 1.35
	b:= 0.98
	x_start:= 1.14
	x_end:= 4.24
	x_delta:= 0.62

	fmt.Println("Решение задачи под буквой А")
	for x := x_start; x < x_end; x += x_delta {
		fmt.Println(result(a, b, x))
	}

	fmt.Println("Решение задачи под буквой B")
	var values = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	for i:=0 ; i< len(values); i++ {
		fmt.Println(result(a, b, values[i]))
	}
}



