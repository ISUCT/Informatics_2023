package main

import (
	"fmt"
	"math"
)

func formulausl(x float64) float64 {
	var a float64
	a = -2.5
	var b float64
	b = 3.4
	return math.Pow(a+b*x, 3.5)/1.8 + math.Pow(math.Cos(a*x), 3)
}
func formulamain(x float64) float64 {
	var a float64
	a = -2.5
	if formulausl(x) >= 5 {
		return (math.Pow(math.Log10(math.Pow(a, 2)+x), 2) / math.Pow(a+x, 2))
	} else {
		return (0)
	}
}
func main() {
	fmt.Println("Гвоздарёв Ярослав Владимирович")
	fmt.Println("задача A")
	for i := 3.5; i <= 6.5; i += 0.6 {
		fmt.Println("При x=", i, "y=", formulamain(i))
	}
	fmt.Println("задача B")
	var list_x [5]float64 = [5]float64{2.89, 3.54, 5.21, 6.28, 3.48}
	for i := 0; i < 5; i++ {
		fmt.Println("При x=", list_x[i], "y=", formulamain(list_x[i]))
	}
}
