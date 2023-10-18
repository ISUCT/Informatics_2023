package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Кремнев Юрий Юрьевич")

	a := 1.6
	x_begin := 1.2
	x_end := 3.7
	x_delta := 0.5
	
	fmt.Println("Решения задачи под А:")
	for x := x_begin; x <= x_end; x = x + x_delta {
		var result = (math.Pow(a,math.Pow(x,2) - 1) - math.Log10(math.Pow(x,2) - 1) + math.Cbrt(math.Pow(x,2) - 1))
		fmt.Println(result)
	}

	fmt.Println("Решения задачи под В:")
	var values = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for i := 0; i < 5; i++ {
		var result = (math.Pow(a,math.Pow(values[i],2) - 1) - math.Log10(math.Pow(values[i],2) - 1) + math.Cbrt(math.Pow(values[i],2) - 1))
		fmt.Println(result)
	}


}
