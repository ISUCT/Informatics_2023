package internal

import (
	"fmt"
	"math"
)

func Summ(a, b int) int {
	return a + b
}

const a = 1.1
const b = 0.09

func down(x, a, b float64) float64 {
	arg := a*x*x - b
	base := 5.0
	return math.Log(arg) / math.Log(base)
}

func up(x float64) float64 {
	return math.Log(x*x*x - 1)
}

func SolveB(numbers []float64) []float64 {
	fmt.Println("\n Задача B:")
	var resultB []float64 = make([]float64, 0, 5)
	for x := 0; x <= 4; x++ {
		upf := up(numbers[x])
		downf := down(numbers[x], a, b)
		resultB = append(resultB, upf/downf)
	}
	return resultB
}

func SolveA(a, b float64) []float64 {
	fmt.Println("Задача A:")
	var resultA []float64 = make([]float64, 0, 6)
	for x := 1.2; x <= 2.2; x += 0.2 {
		upf := up(x)
		downf := down(x, a, b)
		resultA = append(resultA, upf/downf)
	}
	return resultA
}
