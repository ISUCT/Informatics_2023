package internal

import (
	"fmt"
	"math"
)

func Summ(a, b int) int {
	return a + b
}

func Print_answer() {
	var numbers_cycle = []float64{}
	var numbers_massive = []float64{}
	for i := 1.25; i <= 3.25; i = i + float64(0.4) {
		numbers_cycle = append(numbers_cycle, Func_resolution(i))
	}
	var numbers [5]float64 = [5]float64{1.84, 2.71, 3.81, 4.56, 5.62}
	for i := 0; i <= len(numbers)-1; i++ {
		numbers_massive = append(numbers_massive, Func_resolution(numbers[i]))
	}
	fmt.Println(numbers_massive)
	fmt.Println(numbers_cycle)
}

func Func_resolution(x float64) float64 {
	square := square(x)

	diff := square - 2.5

	mod := math.Abs(diff)

	root4 := math.Pow(mod, root_nth_degree(4))

	nat_log := math.Log(square)

	root3 := math.Pow(nat_log, root_nth_degree(3))

	y := root4 + root3

	return y
}

func square(x float64) float64 {
	return x * x
}
func root_nth_degree(degree float64) float64 {
	return 1 / degree
}
