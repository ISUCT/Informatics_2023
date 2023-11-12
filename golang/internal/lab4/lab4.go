package lab4

import "math"

func ToRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func Calculate(x float64, b float64) float64 {
	return (1 + math.Pow(math.Sin(ToRadians(math.Pow(b, 3)+math.Pow(x, 3))), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
}

func TaskA(begX, finX, detX, b float64) ([]float64, []float64) {
	var yList = make([]float64, 0)
	var corrXList = make([]float64, 0)
	for i := begX; i <= finX; i += detX {
		yList = append(yList, Calculate(i, b))
		corrXList = append(corrXList, i)
	}
	return yList, corrXList
}

func TaskB(xList []float64, b float64) []float64 {
	var yList = make([]float64, 0)
	for i := range xList {
		yList = append(yList, Calculate(xList[i], b))
	}
	return yList
}
