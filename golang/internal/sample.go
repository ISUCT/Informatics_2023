package internal

import (
	"math"
)

func radians(deg float64) float64 {
	return (deg) * (math.Pi / 180)
}

func calculate(x float64) float64 {
	return (math.Pow(0.48, 3) + math.Pow(math.Sin(radians(1.2 * x)), 2)) / (math.Acos(x * 0.48 * x) + math.Pow(math.E, -x/2))
}

func TaskA(x_start, x_end, x_delta float64) []float64 {
	var cap float64 = (x_end - x_start) / x_delta
	var cap_int int = int(cap)
	answerA := make([]float64, 0, cap_int)
	for i := x_start; i <= x_end; i += x_delta {
		answerA = append(answerA, calculate(i))
	}
	return answerA
} 

func TaskB(xes []float64) []float64 {
	answerB := make([]float64, 0, len(xes))
	for _, i:= range xes {
		answerB = append(answerB, calculate(i))
	}
	return answerB
}