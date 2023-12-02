package main

import (
	"fmt"
	"myFunctions"
)

func main() {
	var x_n, x_k float64 = 1, 5
	nmob := myFunctions.Nmob(x_n, x_k)
	nmob1 := myFunctions.Nmob1([]float64{6, 7, 8})

	sum := 0.0
	for _, val := range nmob {
		sum += val
	}

	fmt.Println("Sum of Function values for x in range (", x_n, ",", x_k, "]:", sum)
	fmt.Println("Znachenie func №16 pri dannih A:", nmob)
	fmt.Println("Znachenie func №16 pri dannih B:", nmob1)
}
