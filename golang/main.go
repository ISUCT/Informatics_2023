package main

import(
	"math"
	"fmt"
)



func task1(a float64, b float64, x1 float64, x2 float64, deltax float64) {
	var answersA = []float64{}
	for x:=x1; x <= x2; x += deltax{
		var y float64 = (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + math.Exp(-x/2))
		answersA = append(answersA, y)
	}
	fmt.Println(answersA)
}

func task2(a float64, b float64, x1 float64, x2 float64,x3 float64,x4 float64,x5 float64) {
	var znachX = []float64{x1, x2, x3, x4, x5}
	var answersB = []float64{}
	for i:=0; i < len(znachX); i++{
		x := znachX[i]
		var y float64 = (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + math.Exp(-x/2))
		answersB= append(answersB, y)
	}
	fmt.Println(answersB)
}

func main() {
	task1(1.2, 0.48, 0.7, 2.2, 0.3)
	task2(1.2, 0.48, 0.25, 0.36, 0.56, 0.94, 1.28)
}