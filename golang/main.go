package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	if x >= 1 {
		return math.Pow(1.2, x) - math.Pow(x, 1.2)
	} else {
		return math.Acos(x)
	}
}

func task_a(start, end, x float64) {
	for i := start; i < end; i += x {
		fmt.Println(f(i))
	}
}

func task_b(numbers [5]float64) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(f(numbers[i]))
	}
}

func main() {
	fmt.Println("Task A")
	task_a(0.2, 2.2, 0.4)

	fmt.Println("\nTask B")
	task_b([5]float64{0.1, 0.9, 1.2, 1.5, 2.3})
}
