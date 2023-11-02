package main

import (
	"fmt"
	"math"
)

func form(x float64) float64 {
	a := 2.25
	return (math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1/3)))
}
func main() {
	fmt.Println("Целиков Алексей Александрович")
	fmt.Println("Задача A")
	for i := 1.2; i <= 2.7; i += 0.3 {
		fmt.Println("При x=", i, "y=", form(i))
	}
	fmt.Println("Задача B")
	var list_of_x [5]float64 = [5]float64{1.31, 1.39, 1.44, 1.56, 1.92}
	for i := 0; i < 5; i++ {
		fmt.Println("При x=", list_of_x[i], "y=", form(list_of_x[i]))
	}
}
