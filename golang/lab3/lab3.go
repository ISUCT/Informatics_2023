package lab3

import (
	"fmt"
	"math"
)

func result(a, x float64) float64 {
  	return (math.Pow(a,math.Pow(x,2) - 1) - math.Log10(math.Pow(x,2) - 1) + math.Cbrt(math.Pow(x,2) - 1))
}

func taskA(a, x_begin, x_end, x_delta float64) []float64 {
	y := make([]float64, 0, int((x_end - x_begin) / x_delta))
	for x := x_begin; x <= x_end; x += x_delta {
    		y = append(y, result(a, x))
  	}
  	return y
}

func taskB(a float64, values []float64) []float64 {
	y := make([]float64, 0, len(values))
	for _, x := range values {
    		y = append(y, result(a, x))
  	}
  	return y
}

func output(results []float64) {
	for _, i := range results {
		fmt.Println(i)
	}
	fmt.Println("===================================================")
}

func Solve(a, x_begin, x_end, x_delta float64) {
	results := taskA(a, x_begin, x_end, x_delta)
  	fmt.Println("Solutions to task A:")
  	output(results)
  
  	results = taskB(a, []float64{1.28, 1.36, 2.47, 3.68, 4.56})
  	fmt.Println("Solutions to task В:")
  	output(results)
}
