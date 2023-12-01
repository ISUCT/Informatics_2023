package lab4

import (
	"math"
)

func Task1(a, b, xn, xk, deltax float64) []float64 {
    var answers1 []float64 = make([]float64, 0)
    for x := xn; x <= xk; x += deltax {
        y := (a + math.Pow(math.Tan(bx), 2)) / (b + math.Pow(math.Atan(ax), 2))
        answers1 = append(answers1, y)
    }
    return answers1
}

func Task2(a float64, b float64, znX []float64) []float64 {
    var answers2 []float64 = make([]float64, 0)
    for _, value := range  znX {
        x := value
        y := (a + math.Pow(math.Tan(bx), 2)) / (b + math.Pow(math.Atan(ax), 2))
        answers2 = append(answers2, y)
    }
    return answers2
}