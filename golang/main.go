package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Красильникова Анастасия 1/278")
	for i := 1.81; i <= 5.31; i = i + float64(0.7) {
		res := func_resolution(i)
		fmt.Println(res)
	}
	var numbers [5]float64 = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for i := 0; i <= len(numbers)-1; i++ {
		res := func_resolution(numbers[i])
		fmt.Println(res)
	}
}

func func_resolution(x float64) float64 {
	var a float64 = 7.2
	var b float64 = 4.2
	r := b * x
	diff := a - r
	mod := math.Abs(diff)
	pow := math.Pow(math.Log10(x), 3)
	fraction := mod / pow
	square := math.Cbrt(fraction)
	y := square

	return y
}
