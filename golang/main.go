package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вингорадов Антон Алексеевич 1/278")
	for i := 1.25; i <= 3.25; i = i + float64(0.4) {
		res := func_resolution(i)
		fmt.Println(res)
	}
	var numbers [5]float64 = [5]float64{1.84, 2.71, 3.81, 4.56, 5.62}
	for i := 0; i <= len(numbers)-1; i++ {
		res := func_resolution(numbers[i])
		fmt.Println(res)
	}
}

func func_resolution(x float64) float64 {
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
