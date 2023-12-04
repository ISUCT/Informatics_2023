package main

import (
	"fmt"
	"math"
)

func main() {
	zadachkamathForRange()
	zadachkamathForValues()
}

func zadachkamathForRange() {
	a := 4.1
	b := 2.7
	xn := 1.2
	xk := 5.2
	deltaX := 0.8

	for x := xn; x <= xk; x += deltaX {
		calculateAndPrintResult(a, b, x)
	}
}

func zadachkamathForValues() {
	a := 4.1
	b := 2.7
	values := []float64{1.9, 2.15, 2.34, 2.73, 3.16}

	for i, x := range values {
		calculateAndPrintResult(a, b, x, i+1)
	}
}

func calculateAndPrintResult(a, b, x float64, index ...int) {
	chisl := a*math.Sqrt(x) - b*math.Log(x)/math.Log(5)
	znam := math.Log(math.Abs(x-1))/math.Log(10)

	if znam == 0 {
		fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
	} else {
		var message string
		if len(index) > 0 {
			message = fmt.Sprintf("Значение функции для x%d=%f: y%d=%f\n", index[0], x, index[0], chisl/znam)
		} else {
			message = fmt.Sprintf("При x=%f: y=%f\n", x, chisl/znam)
		}
		fmt.Print(message)
	}
}
