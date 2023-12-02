package myFunctions

import (
	"math"
)

func Function(x float64) float64 {
	return math.Log(x) + math.Sqrt(x)
}

func Nmob(x_n, x_k float64) []float64 {
	var nmob []float64
	for x := x_n; x <= x_k; x++ {
		nmob = append(nmob, Function(x))
	}
	return nmob
}

func Nmob1(nums []float64) []float64 {
	var nmob1 []float64
	for _, x := range nums {
		nmob1 = append(nmob1, Function(x))
	}
	return nmob1
}
