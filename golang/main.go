package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.0
	b := 1.1
	fmt.Println("16 вариант")
	fmt.Println("Задача А")
	for x := 0.08; x <= 1.08; x += 0.2 {
		res1 := math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x, 2)))
		res2 := math.Pow(math.Abs(math.Pow(x, 2)-math.Pow(a, 2)), 0.2)
		res := res1 / res2
		fmt.Println(res)
	}
	fmt.Println("Задача В")
	x1 := 0.1
	x2 := 0.3
	x3 := 0.4
	x4 := 0.45
	x5 := 0.65
	y1 := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x1, 2)))) / (math.Pow(math.Abs(math.Pow(x1, 2)-math.Pow(a, 2)), 0.2))
	fmt.Println(y1)
	y2 := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x2, 2)))) / (math.Pow(math.Abs(math.Pow(x2, 2)-math.Pow(a, 2)), 0.2))
	fmt.Println(y2)
	y3 := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x3, 2)))) / (math.Pow(math.Abs(math.Pow(x3, 2)-math.Pow(a, 2)), 0.2))
	fmt.Println(y3)
	y4 := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x4, 2)))) / (math.Pow(math.Abs(math.Pow(x4, 2)-math.Pow(a, 2)), 0.2))
	fmt.Println(y4)
	y5 := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(x5, 2)))) / (math.Pow(math.Abs(math.Pow(x5, 2)-math.Pow(a, 2)), 0.2))
	fmt.Println(y5)
}
