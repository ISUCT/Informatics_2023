package main

import (
	"fmt"
	"math"
)

func y(x float64) float64 {
	a := 1.35
	b := 0.98
	return (math.Cbrt(a * x + b)/math.Pow(math.Log10(x), 2))

}

func main() {
	fmt.Println("Андриянов Роман Ильич")
	fmt.Println("Задача A")
	for i := 1.14; i <= 4.24; i += 0.62 {
		fmt.Println("х = ", i, "y = ", y(i))
	}
	fmt.Println("Задача B")
	var note [5]float64 = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	for i := 0; i <= 4; i ++ {
		fmt.Println("x = ", note[i], "y = ", y(note[i]))
	}

}
