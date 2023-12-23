package lab4

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	if x >= 1 {
		return math.Pow(1.2, x) - math.Pow(x, 1.2)
	} else {
		return math.Acos(x)
	}
}

func TaskA(start, end, x float64) []float64 {
	var answers []float64 = make([]float64, 0, int((end-start)/x))
	for i := start; i < end; i += x {
		answers = append(answers, f(i))
	}
	return answers
}

func TaskB(numbers []float64) []float64 {
	var answers []float64 = make([]float64, 0, len(numbers))
	for i := 0; i < len(numbers); i++ {
		answers = append(answers, f(numbers[i]))
	}
	return answers
}

func Output(answers []float64) {
	for _, i := range answers {
		fmt.Println(i)
	}
}
