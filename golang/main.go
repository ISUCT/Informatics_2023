package main

import (
	"fmt"
	"math"
)

func Primer() {
	a := 2.0
	b := 1.1
	xn := 0.08
	xk := 1.08
	deltaX := 0.2
	e := 2.71828
	x1 := 0.1
	x2 := 0.3
	x3 := 0.4
	x4 := 0.45
	x5 := 0.65
	x := xn
	for x <= xk {
		result := math.Log(math.Abs(b*b-x*x)) / math.Log(e)
		result = result / math.Pow(math.Abs(x*x-a*a), 1/5.0)
		fmt.Printf("x=%.2f, y=%.2f\n", x, result)
		x += deltaX
	}
	values := []float64{x1, x2, x3, x4, x5}
	for _, x := range values {
		result := math.Log(math.Abs(b*b-x*x)) / math.Log(e)
		result = result / math.Pow(math.Abs(x*x-a*a), 1/5.0)
		formattedResult := fmt.Sprintf("%.3f", result)
		fmt.Println("При x=", x, " y=", formattedResult)
	}
}
func main() {
	fmt.Println("Мазуров Даниил Алексеевич")
	Primer()
}
