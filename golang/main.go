package main

import (
	"fmt"
	"math"
)

func Log(x float64) float64 {
	base := 5.0
	return (math.Log(x) / math.Log(base))
}

func form(x float64) float64 {
	a := 1.1
	b := 0.09
	return ((math.Log10(x*x - 1)) / (Log((a*math.Pow(x, 2) - b))))
}
func main() {
	fmt.Println("Кульбакин Аким Александрович")
	fmt.Println("Задача А")
	for i := 1.2; i <= 2.2; i += 0.2 {
		fmt.Println("При x=", i, "y=", form(i))
	}
	fmt.Println("Задача B")
	var list_of_x [5]float64 = [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	for i := 0; i < 5; i++ {
		fmt.Println("При x=", list_of_x[i], "y=", form(list_of_x[i]))
	}
}
