package lab4

import "math"

func toRadians(deg float64) float64 {
	return (deg) * (math.Pi / 180.0)
}

func calculate(a float64, b float64, x float64) float64 {
	var exp float64
	if x < 5 {
		exp = math.Pow(math.Log10(math.Pow(a, 2)+x), 2) / math.Pow(a+x, 2)
	} else {
		exp = math.Pow(a+b*x, 2.5) / (1.8 + math.Pow(math.Cos(toRadians(a*x)), 3))
	}
	return exp
}

func TaskA(a, b, begX, endX, detX float64) ([]float64, []float64) {
	var predX = int(math.Ceil((endX-begX)/detX + 1))
	var ylist = make([]float64, 0, predX)
	var xlist = make([]float64, 0, predX)
	for x := begX; x <= endX; x += detX {
		ylist = append(ylist, calculate(a, b, x))
		xlist = append(xlist, x)
	}
	return ylist, xlist
}

func TaskB(xlist []float64, a, b float64) []float64 {
	var ylist = make([]float64, 0, len(xlist))
	for _, v := range xlist {
		ylist = append(ylist, calculate(a, b, v))
	}
	return ylist
}
