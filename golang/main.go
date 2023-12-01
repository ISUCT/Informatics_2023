package main

import (
	"fmt"
	"math"
)

func part_1() {
	var a float64 = 2.0
	var b float64 = 4.1
	var r_a = []float64{}
	for i := 0.77; i < 1.77; i += 0.2 {
		x := i
		var y float64 = (math.Log2((math.Pow(b, 2))+(math.Pow(x, 2))) / (math.Pow(math.Abs((math.Pow(x, 2) - (math.Pow(a, 2)))), (1.00 / 3.00))))
		r_a = append(r_a, y)
	}
	fmt.Println(r_a)
}

func part_2() {
	var a float64 = 2.0
	var b float64 = 4.1
	var xs = [5]float64{1.24, 1.38, 2.38, 3.21, 0.68}
	var r_b = []float64{}
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		var y float64 = (math.Log2((math.Pow(b, 2))+(math.Pow(x, 2))) / (math.Pow(math.Abs((math.Pow(x, 2) - (math.Pow(a, 2)))), (1.00 / 3.00))))
		r_b = append(r_b, y)
	}
	fmt.Println(r_b)
}

func main() {
	part_1()
	part_2()
	fmt.Println("Чудов Даниил")
	fmt.Println("17")
	fmt.Println("Не женат")
}
