package main

import (
	"fmt"
	"math"
)

func main() {
	task_A()
	fmt.Printf("\n")
	task_B()
}

func task_A() {
	a := 1.6
	deltaX := 0.5
	xStart := 1.2
	xEnd := 3.7
	fmt.Println("task_A")
	for x := xStart; x <= xEnd; x += deltaX {
		y := math.Pow(a, math.Pow(x*x-1, 1)) - math.Log10(x*x-1) + math.Pow(x*x-1, 1.0/3.0)
		fmt.Printf("x = %.2f, y = %.2f\n", x, y)
	}
}

func task_B() {
	a := 1.6
	xValues := []float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println("task_B")
	for _, x := range xValues {
		y := math.Pow(a, math.Pow(x*x-1, 1)) - math.Log10(x*x-1) + math.Pow(x*x-1, 1.0/3.0)
		fmt.Printf("x = %.2f, y = %.2f\n", x, y)
	}
}
