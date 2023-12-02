package main

import (
	"fmt"
	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Петров Леонид Алексеевич")
	fmt.Println(internal.CalculatePoint(1.1, 3.6, 0.5))
	fmt.Println(internal.CalculateFuncWithArguments(2.5, 4.6, []float64 {1.2, 1.28, 1.36, 1.46, 2.35}))
}	
