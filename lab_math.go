package main

import (
	"fmt"
	"math"
)

func main() {

	a := 2.0
	b := 0.95
	x_n := 1.25
	x_k := 2.75
	x_z := 0.3
	const e float64 = 2.71
	fmt.Println("Значение функции №3 при данных задачи A")
	for x := x_n; x <= x_k; x = x_z {
		var y1 = (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(e, x/a))
		fmt.Println(y1)
		break
	}

	var argument = [5]float64{2.2, 3.78, 4.51, 6.58, 1.2}
	fmt.Println("")
	for i := 0; i < 5; i++ {
		var y2 = (1 + math.Pow(math.Log10(argument[i]/a), 2)) / (b - math.Pow(e, argument[i]/a))
		fmt.Println(y2)
		break
	}

}
