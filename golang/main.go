package main

import (
	"fmt"
	"math"
)

// fortask1
var a float64 = 1.1
var b float64 = 0.09
var deltaX float64 = 0.2

// fortask2
var x1 float64 = 1.21
var x2 float64 = 1.76
var x3 float64 = 2.53
var x4 float64 = 3.48
var x5 float64 = 4.52

func task1(xn float64, xk float64) {
	for {
		var x float64 = xn
		var y float64 = (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
		fmt.Println("ResultsOfTask1:" + fmt.Sprint(y))
		xn = xn + deltaX
		if xn >= xk {
			break
		}
	}
}

func task2() {

	peremennie := [5]float64{x1, x2, x3, x4, x5}
	fmt.Println(peremennie)

	for _, xp := range peremennie {
		x := (xp)
		var y float64 = (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
		fmt.Println("ResultsOfTask2:" + fmt.Sprint(y))
	}
}

func main() {
	fmt.Println("Жуков Михаил Андреевич")
	fmt.Println("17 лет")
	task1(1.2, 2.2)
	task2()
}
