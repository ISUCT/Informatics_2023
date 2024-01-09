package internal

import (
	"math"
)

func Summ(a, b int) int {
	return a + b
}

const a = 0.8
const b = 0.4

func Calc(start, finish, delta float64) []float64 {
	var slice = make([]float64, 0, int64((finish-start)/delta+1))
	for x := start; x <= finish; x += delta {
		slice = append(slice, problem1(a, b, x))
	}
	return slice
}

func Calcmassiw(long []float64) []float64 {
	var answer = make([]float64, 0, len(long))
	for _, val := range long {
		answer = append(answer, problem1(a, b, val))
	}
	return answer
}

func problem1(a, b, x float64) float64 {
	result := (math.Pow(math.Pow((x-a), 2), drob(3.0)) + math.Pow(math.Abs(x+b), drob(5.0))) / (math.Pow((math.Pow(x, 2) - math.Pow((a+b), 2)), drob(9.0)))
	return result
}

func drob(y float64) float64 {
	result_1 := 1 / y
	return result_1
}
