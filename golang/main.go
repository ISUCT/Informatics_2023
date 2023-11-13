package main

import (
	"fmt"
	"math"
)

const a, b = 2.0, 1.1

func funkt16(c float64) float64 {
	result := (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(c, 2)))) / (math.Pow(math.Abs(math.Pow(c, 2)-math.Pow(a, 2)), 0.2))
	return result
}
func main() {
	fmt.Println("Парамонов Илья Сергеевич")
	fmt.Println("16 вариант")
	fmt.Println("Задача А")
	for x := 0.08; x <= 1.08; x += 0.2 {
		fmt.Println(funkt16(x))
	}
	fmt.Println("Задача В")
	x1 := 0.1
	fmt.Println(funkt16(x1))
	x2 := 0.3
	fmt.Println(funkt16(x2))
	x3 := 0.4
	fmt.Println(funkt16(x3))
	x4 := 0.45
	fmt.Println(funkt16(x4))
	x5 := 0.65
	fmt.Println(funkt16(x5))
}
