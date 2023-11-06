package main

import (
	"fmt"
	"math"
)

func task1(x1 float64, x2 float64) {
	var a1 float64 = 4.1
	var b1 float64 = 2.7
	var deltax float64 = 0.7
	var answersA = []float64{}
	for x := x1; x <= x2; x += deltax {
		var y float64 = (a1*math.Sqrt(x) - b1*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
		answersA = append(answersA, y)
	}
	fmt.Println(answersA)
}

func task2(x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) {
	var a2 float64 = 4.1
	var b2 float64 = 2.7
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = []float64{}
	for i := 0; i < len(znachX); i++ {
		x := znachX[i]
		var y float64 = (a2*math.Sqrt(x) - b2*(math.Log(x)/math.Log(5))) / (math.Log10(math.Abs(x - 1)))
		answersB = append(answersB, y)
	}
	fmt.Println(answersB)
}

func main() {
	fmt.Println("Полина Костина")
	task1(1.2, 5.2)
	task2(1.9, 2.15, 2.34, 2.73, 3.16)
}
