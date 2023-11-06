package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Петров Леонид Алексеевич")

	a := 2.5
	b := 4.6
	x_n := 1.1
	x_k := 3.6
	x_z := 0.5
	for x := x_n; x <= x_k; x = x_z {
		var dz_a = (math.Pow((a + b*x), 2.5)) / (1 + (math.Log10(a + b*x)))
		fmt.Println(dz_a)
		break
	}

	var argument = [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	for i := 0; i < 5; i++ {
		var dz_b = (math.Pow(a+b*argument[i], 2.5)) / (1 + (math.Log10(a + b*argument[i])))
		fmt.Println(dz_b)
		break
	}
}
