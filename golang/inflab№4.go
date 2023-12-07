package main

import (
	"fmt"
	"math"
)

func main() {
	zadachamathForRange()
	zadachamathForValues()
}

func zadachamathForRange() {
	a := 1.35
	b := 0.98
	xn := 1.14
	xk := 4.24
	deltaX := 0.62

	for x := xn; x <= xk; x += deltaX {
		calculateAndPrintResult(a, b, x)
	}
}

func zadachamathForValues() {
	a := 1.35
	b := 0.98
	values := []float64{0.35, 1.28, 3.51, 5.21, 4.16}

	for i, x := range values {
		calculateAndPrintResult(a, b, x, i+1)
	}
}

func calculateAndPrintResult(a, b, x float64, index ...int) {
	chislitel := (math.Cbrt(a * x + b)
	znamen := math.Pow(math.Log10(x), 2))

	if znamen == 0 {
		fmt.Printf("При x=%f знаменатель равен нулю, деление на ноль невозможно.\n", x)
	} else {
		var message string
		if len(index) > 0 {
			message = fmt.Sprintf("Значение функции для x%d=%f: y%d=%f\n", index[0], x, index[0], chislitel/znamen)
		} else {
			message = fmt.Sprintf("При x=%f: y=%f\n", x, chislitel/znamen)
		}
		fmt.Print(message)
	}
}
