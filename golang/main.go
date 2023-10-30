package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Целиков Алексей Александрович")
}

func form(x float64) float64 {
	var a float64
	a = 2.25
	return (math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1/3)))
}

func math_lab() float64 {
	fmt.Println("Задача A")
	for i := 1.2; i <= 2.7; i += 0.3 {
		return form(i)
	}
	fmt.Println("Задача B")
	var list_of_x [5]float64 = [5]float64{1.31, 1.39, 1.44, 1.56, 1.92}
	for i := 0; i < 5; i++ {
		return (form(list_of_x[i]))
	}
}
