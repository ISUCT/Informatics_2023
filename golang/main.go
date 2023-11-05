package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
	a := 2.25
	xn := 1.2
	xk := 2.7
	deltaX := 0.3
	x1 := 1.31
	x2 := 1.39
	x3 := 1.44
	x4 := 1.56
	x5 := 1.92
	x := xn
	for x <= xk {
		result := (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
		fmt.Printf("\nx = %.2f | y = %.2f\n", x, result)
		x += deltaX
	}
	values := []float64{x1, x2, x3, x4, x5}
	for _, x := range values {
		result := (math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1))
		formattedResult := fmt.Sprintf("%.2f", result)
		fmt.Println("\nЕсли x = " + fmt.Sprint(x))
		fmt.Println("Получается y =", formattedResult)
	}
}
