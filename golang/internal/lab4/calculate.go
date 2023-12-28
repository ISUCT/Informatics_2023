package lab4

import (
  
  "math"
)

func formula(x float64) float64 {
  a :=2.0 
  return ((math.Pow((math.Log10(a+x)),2)) / (math.Pow((a+x),2)))
}

func TaskA(xn,xk,dx float64) ([]float64, []float64){
	x := make([]float64, 0, (int((xk-xn)/dx + 1)))
	y := make([]float64, 0, len(x))
	for i:=xn; i<=xk; i+=dx{
		x = append(x, i)
		y = append(y, formula(i))
	}
	return x,y
}
func TaskB(array []float64) []float64 {
	y := make([]float64, 0, len(array))
	for _, i := range array {
	 y = append(y, f(i))
	}
	return y
}