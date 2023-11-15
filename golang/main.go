package main

import (
	"fmt"
	"math"
)

const a, b = 2.0, 1.1

func funkt16(c float64) float64 { 
	return (math.Log(math.Abs(math.Pow(b, 2) - math.Pow(c, 2)))) / (math.Pow(math.Abs(math.Pow(c, 2)-math.Pow(a, 2)), 0.2))
}

func taskB(d float64) float64 {
	return funkt16(d)
}

func main() {
	fmt.Println("Парамонов Илья Сергеевич")
	fmt.Println("16 вариант")
	fmt.Println("Задача А")
	for x := 0.08; x <= 1.08; x += 0.2 {
		fmt.Println(funkt16(x))
	}
	fmt.Println("Задача В")
	fmt.Println(taskB(0.1))
	fmt.Println(taskB(0.3))
	fmt.Println(taskB(0.4))
	fmt.Println(taskB(0.45))
	fmt.Println(taskB(0.65))
}
