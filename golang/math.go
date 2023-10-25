package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.5
	x1 := 1.1
	x2 := 2.4
	x3 := 3.6
	x4 := 1.7
	x5 := 3.9

	for x := 1.28; x <= 3.28; x += 0.4 {
		verh := (1 + (math.Sin(b*b*b+x*x*x) * math.Sin(b*b*b+x*x*x)))
		niz := math.Cbrt(b*b*b + x*x*x)
		result := verh / niz

		fmt.Println("(Задача А)Результат первой формулы из таблицы - ", result)
	}
	verh1 := (1 + (math.Sin(b*b*b+x1*x1*x1) * math.Sin(b*b*b+x1*x1*x1)))
	niz1 := math.Cbrt(b*b*b + x1*x1*x1)
	result1 := verh1 / niz1

	fmt.Println("(Задача В)Результат формулы при [x1 = 1.1]  - ", result1)

	verh2 := (1 + (math.Sin(b*b*b+x2*x2*x2) * math.Sin(b*b*b+x2*x2*x2)))
	niz2 := math.Cbrt(b*b*b + x2*x2*x2)
	result2 := verh2 / niz2

	fmt.Println("(Задача В)Результат формулы при [x2 = 2.4]  - ", result2)

	verh3 := (1 + (math.Sin(b*b*b+x3*x3*x3) * math.Sin(b*b*b+x3*x3*x3)))
	niz3 := math.Cbrt(b*b*b + x3*x3*x3)
	result3 := verh3 / niz3

	fmt.Println("(Задача В)Результат формулы при [x3 = 3.6]  - ", result3)

	verh4 := (1 + (math.Sin(b*b*b+x4*x4*x4) * math.Sin(b*b*b+x4*x4*x4)))
	niz4 := math.Cbrt(b*b*b + x4*x4*x4)
	result4 := verh4 / niz4

	fmt.Println("(Задача В)Результат формулы при [x4 = 1.7]  - ", result4)

	verh5 := (1 + (math.Sin(b*b*b+x5*x5*x5) * math.Sin(b*b*b+x5*x5*x5)))
	niz5 := math.Cbrt(b*b*b + x5*x5*x5)
	result5 := verh5 / niz5

	fmt.Println("(Задача В)Результат формулы при [x5 = 3.9]  - ", result5)
}

