package main

import(
	"math"
	"fmt"
)



func task1(x1 float64, x2 float64) {
	var a1 float64 = 1.2
	var b1 float64 = 0.48
	var deltax float64 = 0.3
	var answersA = []float64{}
	for x:=x1; x < x2; x += deltax{
		var y float64 = (math.Pow(b1, 3) + math.Pow(math.Sin(a1*x), 2)) / (math.Acos(x*b1*x) + math.Exp(-x/2))
		answersA = append(answersA, y)
	}
	fmt.Println(answersA)
}

func task2(x1 float64,x2 float64,x3 float64,x4 float64,x5 float64) {
	var a2 float64 = 1.2
	var b2 float64 = 0.48
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = []float64{}
	for i:=0; i < len(znachX); i++{
		x := znachX[i]
		var y float64 = (math.Pow(b2, 3) + math.Pow(math.Sin(a2*x), 2)) / (math.Acos(x*b2*x) + math.Exp(-x/2))
		answersB= append(answersB, y)
	}
	fmt.Println(answersB)
}

func main() {
	task1(0.7, 2.2)
	task2(0.25, 0.36, 0.56, 0.94, 1.28)
}