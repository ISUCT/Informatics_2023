package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Полозков Максим Михайлович")
	a := 2.0
	x_n := 1.2
	x_k := 4.2
	x_z := 0.6
	for x := x_n; x <= x_k; x = x + x_z {
		var ya = (math.Pow(math.Log10(a+x), 2) / math.Pow(a+x, 2))
		fmt.Println("Результат задачи A =", ya)

	}
	var argument = [5]float64{1.16, 1.32, 1.47, 1.65, 1.93}
	for z := 0; z < 5; z++ {
		var yb = (math.Pow(math.Log10(a+argument[z]), 2) / math.Pow(a+argument[z], 2))
		fmt.Println("Результат задачи B =", yb)

	}
}
