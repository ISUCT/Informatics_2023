package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
)

func main() {
	fmt.Println("Кульбакин Аким Александрович")
	fmt.Println("Задача А")
	xn := 1.2
	xk := 2.2
	dx := 0.2
	a := 1.1
	b := 0.09
	x, y := lab4.TaksA(xn, xk, dx, a, b)
	for i := range y {
		fmt.Println("при x =", x[i], "y =", y[i])
	}
	fmt.Println("Задача B")
	xarray := []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	x, y = lab4.TaskB(xarray, a, b)
	for i := range y {
		fmt.Println("при x =", x[i], "y =", y[i])
	}
}
