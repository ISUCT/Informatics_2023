package main

import (
	"fmt"
	"math"
)

func task1(a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
}

func task2(_ []float64, a float64, b float64, x float64) float64 {
	return (math.Log10(math.Pow(x, 3)-1) / (math.Log2(a*math.Pow(x, 2)-b) / math.Log2(5)))
}

func main() {
	fmt.Println("Жуков Михаил Андреевич")
	fmt.Println("17 лет")
	//dannie
	var a float64 = 1.1
	var b float64 = 0.09
	//task1
	var deltaX float64 = 0.2
	var xn float64 = 1.2
	var xk float64 = 2.2
	fmt.Println("AnswerOnTask1:")
	for x := xn; x <= xk; x += deltaX {
		fmt.Println(task1(a, b, x))
	}
	//task2
	mySliceForTask2 := []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	fmt.Println("AnswerOnTask2:")
	fmt.Println(mySliceForTask2)
	for _, xp := range mySliceForTask2 {
		x := (xp)
		fmt.Println(task2(mySliceForTask2, a, b, x))
	}
}
