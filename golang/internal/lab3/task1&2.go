package lab3

import ("math")

func Calc(a, b, x float64) float64 {
	return (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + math.Exp(-x/2))
}

func Task1(a float64, b float64, x1 float64, x2 float64, deltax float64) []float64{
	var answersA []float64 = make([]float64, 0, int((x2 - x1)/deltax) + 1)
	for x:=x1; x <= x2; x += deltax{
		y := Calc(a,b,x)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(a float64, b float64, znachx[] float64) []float64{
	var answersB []float64 = make([]float64, 0, len(znachx))
	for _, value := range znachx{
		x := value
		y := Calc(a,b,x)
		answersB = append(answersB, y)
	}
	return answersB
}