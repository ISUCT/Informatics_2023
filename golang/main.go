package main

import (
	"fmt"
	"math"
)

func task1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64 {
	answersA := make([]float64, 0, 10)
	for x := x1; x <= x2; x += deltax {
		var y float64 = math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3)))
		answersA = append(answersA, y)
	}
	return answersA
}
func task2(a float64, b float64, xValues []float64) []float64 {
	answersB := make([]float64, 0, 10)
	for _, x := range xValues {
		var y float64 = math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3)))
		answersB = append(answersB, y)
	}
	return answersB
}

func main() {
	fmt.Println("Kochetkov Kirill")
	results1 := task1(7.2, 4.2, 1.81, 5.31, 0.7)
	fmt.Println(results1)
	results2 := task2(7.2, 4.2, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	fmt.Println(results2)
}
