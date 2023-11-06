package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ахмадеева Арина Игоревна")
	a := 2.0
	b := 0.95
	xn := 1.25
	xk := 2.75
	deltax := 0.3
	e := 2.71828
	x1 := 2.2
	x2 := 3.78
	x3 := 4.51
	x4 := 6.58
	x5 := 1.2
	x := xn
	for x <= xk {
		result := (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(e, x/a))
		fmt.Printf("x=%.2f, y=%.2f\n", x, result)
		x += deltax
	}
	values := []float64{x1, x2, x3, x4, x5}
	for _, x := range values {
		result := (1 + math.Pow(math.Log10(x/a), 2)) / (b - math.Pow(e, x/a))
		formattedResult := fmt.Sprintf("%.3f", result)
		fmt.Println("При x=", x, " y=", formattedResult)
	}
}
