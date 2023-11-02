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
	fmt.Println("При x = ", x, " y = ", y)
	return 1
}

func main() {
	fmt.Println("Задача A")
	for i := 0.2; i <= 2.2; i += 0.4 {
		znachenie_funkcii(i)
	}
	fmt.Println("Задача B")
	var list_of_x [5]float64 = [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	var x int
	for x < 5 {
		znachenie_funkcii(list_of_x[x])
		x++
	}

}
