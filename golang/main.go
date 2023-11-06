package main

import (
	"fmt"
	"math"
)

var a float64 = 2.0
var deltaX float64 = 0.6

var x1 float64 = 1.16
var x2 float64 = 1.32
var x3 float64 = 1.47
var x4 float64 = 1.65
var x5 float64 = 1.93

func work1(xn float64, xk float64) {
	for {
		var x float64 = xn
		var y float64 = (math.Pow(math.Log10(a+x), 2)) / (math.Pow((a + x), 2))
		fmt.Println("ResultsOfWork1:" + fmt.Sprint(y))
		xn = xn + deltaX
		if xn >= xk {
			break
		}
	}
}

func work2() {
	Met := [5]float64{x1, x2, x3, x4, x5}
	fmt.Println(Met)
	for _, xp := range Met {
		x := xp
		var y float64 = (math.Pow(math.Log10(a+x), 2)) / (math.Pow((a + x), 2))
		fmt.Println("ResultsOfWork2:" + fmt.Sprint(y))
	}
}

func main() {
	fmt.Println("Рыбников Юрий Анатольевич")
	work1(1.2, 4.2)
	work2()
}
