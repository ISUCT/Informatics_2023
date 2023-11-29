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

func TaskA(xBeg, xEnd, detX float64) ([]float64, []float64) {
	var cap int = int(math.Ceil((xEnd-xBeg)/detX + 1))
	var yList = make([]float64, 0, cap)
	var xlist2 = make([]float64, 0, cap)
	for i := xBeg; i <= xEnd; i += detX {
		yList = append(yList, formula(i))
		xlist2 = append(xlist2, i)
	}
	return yList, xlist2
}
func TaskB(xList []float64) []float64 {
	var yList = make([]float64, 0)
	for i := range xList {
		yList = append(yList, formula(xList[i]))
	}
	return yList
}
