package lab4

import (
	"fmt"
	"math"
)

func formula(a, x float64) float64 {
	return ((math.Pow((math.Log(a + x)), 2)) / (math.Pow((a + x), 2)))
}

func TaskA(a, x_start, x_end, x_step float64) []float64 {
	y := make([]float64, 0, int((x_end-x_start)/x_step))
	for x := x_start; x <= x_end; x += x_step {
		y = append(y, formula(a, x))
	}
	return y
}

func TaskB(a float64, array []float64) []float64 {
	y := make([]float64, 0, len(array))
	for _, x := range array {
		y = append(y, formula(a, x))
	}
	return y
}

func output(results []float64) {
	for _, i := range results {
		fmt.Println(i)
	}
}

func Solve(a, x_start, x_end, x_step float64) {
	results := TaskA(a, x_start, x_end, x_step)
	fmt.Println("Exercise_A:")
	output(results)

	results = TaskB(a, []float64{1.16, 1.32, 1.47, 1.65, 1.93})
	fmt.Println("Exercise_В:")
	output(results)
}
