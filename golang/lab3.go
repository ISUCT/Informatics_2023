package main

import (
	"fmt"
	"math"
)

func Calculate(a float64, b float64, x float64) float64 {
	return (math.Cbrt(a*x+b) / math.Pow(math.Log10(x), 2))
}

func main() {
	var a = 1.35
	var b = 0.98
	var begx = 1.14
	var finx = 4.24
	var detx = 0.62
	for i := begx; i <= finx; i += detx {
		fmt.Println(Calculate(a, b, i))
	}
	var xList = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	for i := 0; i < len(xList); i++ {
		fmt.Println(Calculate(a, b, xList[i]))
	}
}
