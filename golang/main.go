package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	a := 2.0
	b := 1.1
	return (math.Log(math.Abs(math.Pow(b, 2)-math.Pow(x, 2))) / (math.Pow(math.Abs(math.Pow(x, 2)-math.Pow(a, 2)), 0.2)))
}
func main() {
	fmt.Println("Родин Владислав Евгеньевич")
	fmt.Println("Задача А")
	for i := 0.08; i <= 1.08; i += 0.2 {
		fmt.Println("При x=", i, "y=", formula(i))
	}
	fmt.Println("Задача B")
	var list [5]float64 = [5]float64{0.1, 0.3, 0.4, 0.45, 0.65}
	for i := 0; i <= 4; i++ {
		fmt.Println("При x=", list[i], "y=", formula(list[i]))
	}
}
