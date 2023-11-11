package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Cycle(xn, xk, xd float64) []float64 {
	var numbers_cycle = []float64{}
	for i := xn; i <= xk; i = i + xd {
		numbers_cycle = append(numbers_cycle, Func_resolution(i))
	}
	return numbers_cycle
}

func Massive(numbers [5]float64) []float64 {
	var numbers_massive = []float64{}
	for _, val := range numbers {
		numbers_massive = append(numbers_massive, Func_resolution(val))
	}
	return numbers_massive
}

func Func_resolution(x float64) float64 {
	square := math.Pow(x, 2)
	diff := square - 2.5
	mod := math.Abs(diff)
	root4 := math.Pow(mod, float64(1)/float64(4))
	nat_log := math.Log(square)
	root3 := math.Pow(nat_log, float64(1)/float64(3))
	y := root4 + root3
	return y
}
