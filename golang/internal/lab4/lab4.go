package lab4

import (
	"math"
)

func Zadacha(x float64, a float64, b float64) float64 {
	return (math.Sqrt((math.Abs(a - b*x)) / (math.Pow(math.Log10(x), 3))))
}

func TaskA(xBegin, xEnd, xDelta, a, b float64) []float64 {
	var massive int = int((xEnd - xBegin) / xDelta)
	list := make([]float64, 0, massive)
	for n1 := xBegin; n1 <= xEnd; n1 += xDelta {
		list = append(list, Zadacha(n1, a, b))
	}
	return list
}

func TaskB(list []float64, a, b float64) []float64 {
	list2 := make([]float64, 0, len(list))
	for n2 := range list {
		list2 = append(list2, Zadacha(list[n2], a, b))
	}
	return list2
}



