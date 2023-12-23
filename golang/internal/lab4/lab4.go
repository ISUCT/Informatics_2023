package lab4

import (
	"math"
)

func radians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func formula(x float64) float64 {
	return ((math.Pow(math.Sin(radians(x)), 3) + math.Pow(math.Cos(radians(x)), 3)) * math.Log(x))
}

func TaskA(xBeg, xEnd, detX float64) []float64 {
	var lenx int = int((xEnd-xBeg)/detX + 1)
	z := make([]float64, 0, lenx)
	for i := xBeg; i <= xEnd; i += detX {
		z = append(z, formula(i))
	}
	return z
}
func TaskB(list []float64) []float64 {
	var list2 = make([]float64, 0, len(list))
	for i := range list {
		list2 = append(list2, formula(list[i]))
	}
	return list2
}
