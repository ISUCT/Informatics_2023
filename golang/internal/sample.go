package internal

import (
	"fmt"
	"math"
)

const b = 2.5

func funktion1(c float64) float64 {
	result := 1 + math.Sin(b*b*b+c*c*c)*math.Sin(b*b*b+c*c*c)/math.Cbrt(b*b*b+c*c*c)
	return result
}

func main() {
	fmt.Println("1  Вариант")
	fmt.Println("Задача А")
	for x := 1.28; x <= 3.28; x += 0.4 {
		fmt.Println(funktion1(x))
	}
	fmt.Println("Задача B")
	x1 := 1.1
	fmt.Println(funktion1(x1))
	x2 := 2.4
	fmt.Println(funktion1(x2))
	x3 := 3.6
	fmt.Println(funktion1(x3))
	x4 := 1.7
	fmt.Println(funktion1(x4))
	x5 := 3.9
	fmt.Println(funktion1(x5))

}
