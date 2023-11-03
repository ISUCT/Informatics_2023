package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	return (math.Pow(1.6, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt((math.Pow(x, 2) - 1)))

}

func main() {
	fmt.Println("Задача A")
	for i := 1.2; i <= 3.8; i += 0.5 {
		fmt.Println("При x=", i, "y=", formula(i))
	}
	fmt.Println("Задача B")
	var listik [5]float64 = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for i := 0; i < 5; i++ {
		fmt.Println("При x=", listik[i], "y=", formula(listik[i]))

	}
}
