package lab4

import (
	"math"
)

func calculate(a float64, b float64, x float64) float64 {
	return (math.Cbrt(a*x+b) / math.Pow(math.Log10(x), 2))
}

// task 1
func TaskA(begX, finX, detX, a, b float64) ([]float64, []float64) {
	var predeterminedCap int = int(math.Ceil((finX-begX)/detX + 1))
	var yList = make([]float64, 0, predeterminedCap)
	var corrXList = make([]float64, 0, predeterminedCap)
	for i := begX; i <= finX; i += detX {
		yList = append(yList, calculate(i, a, b))
		corrXList = append(corrXList, i)
	}
	return yList, corrXList
}

// task 2
func TaskB(xList []float64, a float64, b float64) []float64 {
	var yList = make([]float64, 0)
	for i := range xList {
		yList = append(yList, calculate(xList[i], a, b))
	}
	return yList
}
