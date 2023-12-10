package internal

import (
	"math"
)

const b = 2.5

func fnk1(b, c float64) float64 {
	result := 1 + math.Sin(b*b*b+c*c*c)*math.Sin(b*b*b+c*c*c)/math.Cbrt(b*b*b+c*c*c)
	return result
}

func Cycle(xn, xk, xs float64) []float64 {
	var valuefnk = []float64{}
	for x := xn; x <= xk; x += xs {
		valuefnk = append(valuefnk, fnk1(x))
	}
	return valuefnk
}	

func Srez(nums []float64) []float64{
	var nums_of_srez = make([]float64, 0, len(nums))
	for _, value := range nums {
		nums_of_srez = append(nums_of_srez, fnk1(value))
	}
	return nums_of_srez
}
