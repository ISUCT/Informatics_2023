package lab4

import "math"

func ToRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func Calculate(x float64) float64 {
	var exp float64
	exp = (math.Pow(math.Sin(ToRadians(x)), 3) + math.Pow(math.Cos(ToRadians(x)), 3)) * math.Log10(x)
	return exp
}

//TASK 1
func TaskA(begX, endX, detX float64) ([]float64, []float64) {
	predX := int(math.Ceil((endX-begX)/detX + 1))
	yList := make([]float64, 0, predX)
	xList := make([]float64, 0, predX)
	for x := begX; x <= endX; x += detX {
		yList = append(yList, Calculate(x))
		xList = append(xList, x)
	}
	return yList, xList
}

//TASK 2
func TaskB(xList []float64) []float64 {
	var yList = make([]float64, 0, len(xList))
	for _, v := range xList {
		yList = append(yList, Calculate(v))
	}
	return yList
}
