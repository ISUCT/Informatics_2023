package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Cycle(xn, xk, xd float64) []float64 {
	var numbers_cycle = make([]float64, 0, int(((xk-xn)/xd)+1))
	for i := xn; i <= xk; i = i + xd {
		numbers_cycle = append(numbers_cycle, FuncResolution(i))
	}
	return numbers_cycle
}

func Massive(numbers []float64) []float64 {
	var numbers_massive = make([]float64, 0, len(numbers))
	for _, val := range numbers {
		numbers_massive = append(numbers_massive, FuncResolution(val))
	}
	return numbers_massive
}

func FuncResolution(x float64) float64 {
	square := math.Pow(x, 2)
	diff := square - 2.5
	mod := math.Abs(diff)
	root4 := math.Pow(mod, 1.0/4.0)
	nat_log := math.Log(square)
	root3 := math.Pow(nat_log, float64(1)/float64(3))
	y := root4 + root3
	return y
}
