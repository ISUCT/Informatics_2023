// gurev ivan 1-279 formula 6
package main

import (
	"fmt"
	"math"
)

func znachenie_funkcii(x float64) float64 {
	var y float64
	x = math.Round(math.Abs(x)*100) / 100
	if x < 1 {
		y = math.Acos(x)
	} else if x >= 1 {
		y = math.Pow(1.2, x) - math.Pow(x, 1.2)
	}
	return y
}
func main() {
	fmt.Println("Задача A")
	for i := 0.2; i <= 2.2; i += 0.4 {
		i = math.Round(math.Abs(i)*100) / 100
		y := znachenie_funkcii(i)
		fmt.Println("При x = ", i, ", y = ", y)
	}
	fmt.Println("Задача B")
	var list_of_x [5]float64 = [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	var x int
	for x < 5 {
		list_of_x[x] = math.Round(math.Abs(list_of_x[x])*100) / 100
		y := znachenie_funkcii(list_of_x[x])
		fmt.Println("При x = ", list_of_x[x], ", y = ", y)
		x++
	}
}
