package main

import (
	"fmt"
	"math"
)

func main() {

	funcs := []func(){
		zadacha_A,
		zadacha_B,
	}

	for _, f := range funcs {
		f()
	}

}

func zadacha_A() {
	fmt.Println("Задача A:")
	x := 0.2
	for x == x {
		if math.Abs(x) >= 1 {
			y := math.Pow(1.2, x) - math.Pow(x, 1.2)
			fmt.Println(y)
		}
		if math.Abs(x) < 1 {
			y := math.Acos(x)
			fmt.Println(y)
		}
		if x > 2.2 {
			break
		}
		x = x + 0.4
	}
}

func zadacha_B() {
	numbers := [5]float64{0.1, 0.9, 1.2, 1.5, 2.3}
	fmt.Println("Задача B:")
	for _, z := range numbers {
		if math.Abs(z) >= 1 {
			y := math.Pow(1.2, z) - math.Pow(z, 1.2)
			fmt.Println(y)
		}
		if math.Abs(z) < 1 {
			y := math.Acos(z)
			fmt.Println(y)
		}
	}
}
