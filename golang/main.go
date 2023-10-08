package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Iakovleva Daria Dmitrievna 1/278")
	fmt.Printf("\n")
	fmt.Println("24 вариант")
	task_A()
	fmt.Printf("\n")
	task_B()
}

func task_A() {
	a := 7.2
	b := 1.3
	xn := 1.56
	xk := 4.71
	dx := 0.63
	fmt.Println("task_A")
	for x := xn; x <= xk; x += dx {
		y := math.Pow(((a + b*x) / (math.Pow((math.Log10(x)), 3))), 1.0/5.0)
		fmt.Printf("%.2f\t", x)
		fmt.Printf("%.2f\n", y)
	}
}
func task_B() {
	a := 7.2
	b := 1.3
	vhodnie_x := []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("task_B")
	for _, x := range vhodnie_x {
		y := math.Pow(((a + b*x) / (math.Pow((math.Log10(x)), 3))), 1.0/5.0)
		fmt.Printf("%.2f\t", x)
		fmt.Printf("%.2f\n", y)
	}
}
