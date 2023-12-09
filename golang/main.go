package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	var a = 0.4
	var b = 0.8
	fmt.Println(internal.TaskA(a, b, 3.2, 6.2, 0.6))
	var xList = []float64{4.48, 3.56, 2.78, 5.28, 3.21}
	fmt.Println(internal.TaskB(xList, a, b), xList)
}
