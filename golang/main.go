package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	return math.Asin(math.Pow(x, 2)) + math.Asin(math.Pow(x, 3))
}

func main() {
	fmt.Println("Задача А")
	for i := 0.11; i <= 0.36; i += 0.05 {
		fmt.Println("Если х = ", i, "то y = ", formula(i))
	}
	fmt.Println("Задача B")
	var array [5]float64 = [5]float64{0.08, 0.26, 0.35, 0.41, 0.53}
	for i := 0; i < 5; i++ {
		fmt.Println("Если x = ", array[i], "то y = ", formula(array[i]))
	}
}
