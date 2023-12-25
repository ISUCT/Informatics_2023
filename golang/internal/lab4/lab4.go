package lab4

import (
	"math"
)

func toRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func calculate(x float64, b float64) float64 {
	return (1 + math.Pow(math.Sin(toRadians(math.Pow(b, 3)+math.Pow(x, 3))), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
}

func TaskA(begX, finX, detX, b float64) ([]float64, []float64) {
	var predeterminedCap int = int(math.Ceil((finX-begX)/detX + 1))
	var yList = make([]float64, 0, predeterminedCap)
	var corrXList = make([]float64, 0, predeterminedCap)
	for i := begX; i <= finX; i += detX {
		yList = append(yList, calculate(i, b))
		corrXList = append(corrXList, i)
	}
	return yList, corrXList
}
func TaskB(xList []float64, b float64) []float64 {
	var yList = make([]float64, 0, len(xList))
	for i := range xList {
		yList = append(yList, calculate(xList[i], b))
	}
	return yList
}
