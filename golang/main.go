package main

import (
	"fmt"
	"math"
)

func task1(x1 float64, x2 float64, deltax float64) {
	var a1 float64 = 7.2
	var b1 float64 = 4.2
	var answersA = []float64{}
	for x := x1; x <= x2; x += deltax {
		var y float64 = math.Sqrt((math.Abs(a1 - b1*x)) / (math.Pow(math.Log10(x), 3)))
		answersA = append(answersA, y)
	}
	fmt.Println(answersA)
}

func task2(x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) {
	var a2 float64 = 7.2
	var b2 float64 = 4.2
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = []float64{}
	for i := 0; i < len(znachX); i++ {
		x := znachX[i]
		var y float64 = math.Sqrt((math.Abs(a2 - b2*x)) / (math.Pow(math.Log10(x), 3)))
		answersB = append(answersB, y)
	}
	fmt.Println(answersB)
}

func main() {
	fmt.Println("Kochetkov Kirill")
	task1(1.81, 5.31, 0.7)
	task2(2.4, 2.8, 3.9, 4.7, 3.16)
}
