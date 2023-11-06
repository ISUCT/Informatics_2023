package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Смирнов Глеб Денисович 1/280")

	a := 2.5
	b := 4.6
	x_begin := 1.1
	x_end := 3.6
	x_delta := 0.5
	x1 := 1.2
	x2 := 1.28
	x3 := 1.36
	x4 := 1.46
	x5 := 2.35

	fmt.Println("Решения задачи под A:")
	for x := x_begin; x <= x_end; x = x + x_delta {
		chisl := math.Pow(a+b*x, 2.5)
		znam := 1 + math.Log10(a+b*x)
		if znam == 0 {
			fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
		} else {
			y := chisl / znam
			fmt.Printf("При x=%f: y=%f\n", x, y)
		}
	}
	fmt.Println("Решения задачи под B:")
	values := []float64{x1, x2, x3, x4, x5}
	for i, x := range values {
		chisl := math.Pow(a+b*x, 2.5)
		znam := 1 + math.Log10(a+b*x)
		if chisl == 0 {
			fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
		} else {
			y := chisl / znam
			fmt.Printf("Значение функции для x%d=%f: y%d=%f\n", i+1, x, i+1, y)
		}
	}
}
