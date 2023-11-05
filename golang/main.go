package main

import (
	"fmt"
	"math"
)

func main() {
	zadachkamath()
}
func zadachkamath() {
	a := 4.1
	b := 2.7
	xn := 1.2
	xk := 5.2
	deltaX := 0.8
	x1 := 1.9
	x2 := 2.15
	x3 := 2.34
	x4 := 2.73
	x5 := 3.16
	for x := xn; x <= xk; x += deltaX {
		chisl := a * math.Sqrt(x) - b*math.Log(x)/math.Log(5)
		znam := math.Log(math.Abs(x-1))/math.Log(10)
		if znam == 0 {
			fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
		} else {
			y := chisl / znam
			fmt.Printf("При x=%f: y=%f\n", x, y)
		}
	}
	values := []float64{x1, x2, x3, x4, x5}
	for i, x := range values {
		chisl := a * math.Sqrt(x) - b*math.Log(x)/math.Log(5)
		znam := math.Log(math.Abs(x-1))/math.Log(10)
		if chisl == 0 {
			fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
		} else {
			y := chisl / znam
			fmt.Printf("Значение функции для x%d=%f: y%d=%f\n", i+1, x, i+1, y)
		}
	}
}
