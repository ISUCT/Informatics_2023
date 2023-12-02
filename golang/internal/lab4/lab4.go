package lab4

import (
	"math"
)

func Assignment_1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64 {
	var res = []float64{}
	for i := x1; i < x2; i += deltax {
		x := i
		y := (math.Pow((x-a), (2.00/3.00)) + math.Pow(math.Abs(x+b), (1.00/5.00))) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), (1.00/9.00))
		res = append(res, y)
	}
	return res
}

func Assignment_2(a float64, b float64, x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) []float64 {
	var xs = []float64{x1, x2, x3, x4, x5}
	var res = []float64{}
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		y := (math.Pow((x-a), (2.00/3.00)) + math.Pow(math.Abs(x+b), (1.00/5.00))) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), (1.00/9.00))
		res = append(res, y)
	}
	return res
}
