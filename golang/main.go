package main

import (
	"fmt"
	"math"
)

const a float64 = 4.1
const b float64 = 2.7
const x_start float64 = 1.2
const x_end float64 = 5.2
const x_delta float64 = 0.8

var xs = [5]float64{1.9, 2.15, 2.34, 2.73, 3.16}

func function(x float64) float64 {
	return ((a * math.Sqrt(x)) - (b * Log5(x))) / (math.Log10(math.Abs(x - 1)))
}

func Log5(x float64) float64 {
	return math.Log(x) / math.Log(5)
}

func main() {
	fmt.Println("Коротков Михаил Александрович")
	fmt.Println("Вариант №13")
	fmt.Println("Задача A")
	for i := x_start; i <= x_end; i += x_delta {
		fmt.Printf("%.5f\n", function(i))
	}
	fmt.Println("\nЗадача B")
	for i := 0; i < len(xs); i++ {
		fmt.Printf("%.5f\n", function(xs[i]))
	}
}
