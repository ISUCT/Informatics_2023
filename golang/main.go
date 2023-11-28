package main

import (
	"fmt"
	"math"
)

func calcY(a float64, b float64, x float64) float64 {
	return (a*math.Sqrt(x) - b*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
}

func task1(a float64, b float64, x1 float64, x2 float64, deltax float64) {
	var answersA = make([]float64, 0, int(math.Round(x2-x1)/deltax))
	for x := x1; x <= x2; x += deltax {
		y := calcY(a, b, x)
		answersA = append(answersA, y)
	}
	fmt.Println(answersA)
}

func task2(a float64, b float64, x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) {
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = make([]float64, 0, len(znachX))
	for i := 0; i < len(znachX); i++ {
		x := znachX[i]
		y := calcY(a, b, x)
		answersB = append(answersB, y)
	}
	fmt.Println(answersB)
}

func main() {
	task1(4.1, 2.7, 1.2, 5.2, 0.7)
	task2(4.1, 2.7, 1.9, 2.15, 2.34, 2.73, 3.16)
	fmt.Println("Hello world")
}
