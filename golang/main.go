package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	return (math.Pow(math.Pow(math.Asin(x), 2)+math.Pow(math.Acos(x), 4), 3))
}
func main() {
	fmt.Println("Прибылов Алексей Александрович")
	fmt.Println("Задача A")
	for i := 0.26; i <= 0.66; i += 0.08 {
		fmt.Println("При x=", i, "y=", formula(i))
	}
	fmt.Println("Задача B")
	var list [5]float64 = [5]float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for i := 0; i <= 4; i++ {
		fmt.Println("При x=", list[i], "y=", formula(list[i]))
	}
}
