package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

const b = 2.5

func fnk1(b, x float64) float64 {
	result := 1 + math.Sin(b*b*b+x*x*x)*math.Sin(b*b*b+x*x*x)/math.Cbrt(b*b*b+x*x*x)
	return result
}

func Cycle(xn, xk, xs float64) []float64 {
	var valuefnk  = make([]float64, 0, int((xk-xn)/xs) +1)
	for x := xn; x <= xk; x += xs {
		valuefnk = append(valuefnk, fnk1(b, x))
	}
	return valuefnk
}	

func Srez(b float64, nums []float64) []float64{
	var nums_of_srez = make([]float64, 0, len(nums))
	for _, x := range nums {
		nums_of_srez = append(nums_of_srez, fnk1(x, b))
	}
	return nums_of_srez
}

