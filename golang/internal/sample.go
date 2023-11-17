package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}
func log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

func Func_res(a, b, x float64) float64 {
	return (a*math.Cbrt(x) - b*log(5, x)) / math.Pow(math.Log10(x-1), 3)
}
