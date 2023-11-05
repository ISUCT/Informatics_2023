package main

import (
	"fmt"
	"math"
)

var deltaX float64 = 0.08
var x1 float64 = 0.1
var x2 float64 = 0.35
var x3 float64 = 0.4
var x4 float64 = 0.55
var x5 float64 = 0.6

func primer1(xn float64, xk float64) {
	for {
		var x float64 = xn
		var y float64 = math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
		fmt.Println("Пример 1-ый:" + fmt.Sprint(y))
		xn = xn + deltaX
		if xn >= xk {
			break
		}
	}
}
func primer2() {
	dannie := [5]float64{x1, x2, x3, x4, x5}
	for _, xp := range dannie {
		x := xp
		var y float64 = math.Pow((math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)), 3)
		fmt.Println("Пример 2-ой:" + fmt.Sprint(y))
	}
}
func main() {
	fmt.Println("Кузьмин Владислав Иванович")
	primer1(0.26, 0.66)
	primer2()
}
