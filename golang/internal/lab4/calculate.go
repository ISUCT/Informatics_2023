package lab4

import (
  
  "math"
)

func formula(x float64) float64 {
  a :=2.0 
  return ((math.Pow((math.Log10(a+x)),2)) / (math.Pow((a+x),2)))
}

fanc TaskA(xn,xk,dx float64) ([]float64, []float64){
	x:=float64
	y:=float64
	for i:=xn; i<=xk; i+=dx{
		x = appennd(x,i)
		y = appennd(y,f(i))
	}
	return x,y
}
func TaskB(array []float64) ([]float64, []float64){
	y:=[]float64{}
	x:=[]float64{}
	for _, i:=range array{
		y=appennd(y,formula(i))
		x=appennd(x,i)
	}
	return x,y
}