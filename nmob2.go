package main

import (
	"fmt"
	"math"
)

func nmob2() {

	a := 2.0
	b := 1.1
	x_n := 0.08
	x_k := 1.08
	x_z := 0.2
	fmt.Println("Znachenie func №16 pri dannix A")
	for x := x_n; x <= x_k; x = x_z {
		// ln(a) | b2 - x2 |
		var1 := math.Log(a) * math.Abs(b*b-x*x)

		// 5√ | x2 - a2 |
		var2 := 5 * math.Sqrt(math.Abs(x*x-a*a))
		y1 := var1 / var2
		fmt.Println(y1)
		break
	}

	fmt.Println("Znachenie func №16 pri dannix B")
	var argument = [5]float64{0.1, 0.3, 0.4, 0.45, 0.65}
	fmt.Println("")
	for i := 0; i < 5; i++ {
		// ln(a) | b2 - x2 |
		var1 := math.Log(a) * math.Abs(b*b-argument[i]*argument[i])

		// 5√ | x2 - a2 |
		var2 := 5 * math.Sqrt(math.Abs(argument[i]*argument[i]-a*a))
		y2 := var1 / var2
		fmt.Println(y2)
		break
	}

}
