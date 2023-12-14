package internal

import (
	"fmt"
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Zadacha_A(x float64) []float64 {
	fmt.Println("Задача A:")
	var resultA []float64
	for x <= 2.2 {
		if math.Abs(x) >= 1 {
			y := math.Pow(1.2, x) - math.Pow(x, 1.2)
			resultA = append(resultA, y)
		}
		if math.Abs(x) < 1 {
			y := math.Acos(x)
			resultA = append(resultA, y)
		}
		x += 0.4
	}
	return resultA
}

func Zadacha_B(numbers []float64) []float64 {
	fmt.Println("Задача B:")
	var resultB []float64
	for _, z := range numbers {
		if math.Abs(z) >= 1 {
			y := math.Pow(1.2, z) - math.Pow(z, 1.2)
			resultB = append(resultB, y)
		}
		if math.Abs(z) < 1 {
			y := math.Acos(z)
			resultB = append(resultB, y)
		}
	}
	return resultB
}
