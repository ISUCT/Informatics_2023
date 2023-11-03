package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	a := 1.2
	b := 0.48
	e := 2.71
	return ((math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos((x * b * x)) + math.Pow(e, (-x/2))))
}
func main() {
	fmt.Println("Корчагин Пётр Сергеевич")
	fmt.Println("Ответы на задачу A")
	for i := 0.7; i <= 2.2; i += 0.3 {
		fmt.Println("x=", i, "y=", formula(i))
	}
	fmt.Println("Ответы на задачу B")
	var listx [5]float64 = [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	for i := 0; i < 5; i++ {
		fmt.Println("x=", listx[i], "y=", formula(listx[i]))
	}
}
