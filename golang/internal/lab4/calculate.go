package lab4

import (
  "math"
)

func f(x float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(math.Pow(x, 2)-2.5))) + math.Cbrt(math.Log10(math.Pow(x, 2))))
}

func TaskA(xn, xk, dx float64) ([]float64, []float64){
	x:=[]float64
	y:=[]float64
	for i:=xn; i <= xk; i += dx{
		x = appennd(x, i)
		y = appennd(y, f(i))
	}
	return x,y
}
func TaskB(array []float64) ([]float64, []float64){
	x:=[]float64
	y:=[]float64
	for _, i := range array{
		x = appennd(x, i)
		y = appennd(y, f(i))
	}
	return x,y
}