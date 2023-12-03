package internal

import (
	"fmt"
	"math"
)

const b = 2.5

func fnk1(c float64) float64 {
	result := 1 + math.Sin(b*b*b+c*c*c)*math.Sin(b*b*b+c*c*c)/math.Cbrt(b*b*b+c*c*c)
	return result
}

func cycle(xn, xk, xs float64) []float64 {
	var valuefnk = []float64{}
	for x := xn; x <= xk; x += xs {
		valuefnk = append(valuefnk, fnk1(x))
	}
	return valuefnk
}	

func Srez(nums []float64) []float64{
	var nums_of_srez = make([]float64, 0, len(nums))
	for _, value := range nums {
		nums_of_srez = append(nums_of_srez, fnk1(value))
	}
	return nums_of_srez
}

func fnk2() {
	fmt.Println("1  Вариант")
	fmt.Println("Задача А")
	for x := 1.28; x <= 3.28; x += 0.4 {
		fmt.Println(fnk1(x))
	}
	fmt.Println("Задача B")
	x1 := 1.1
	fmt.Println(fnk1(x1))
	x2 := 2.4
	fmt.Println(fnk1(x2))
	x3 := 3.6
	fmt.Println(fnk1(x3))
	x4 := 1.7
	fmt.Println(fnk1(x4))
	x5 := 3.9
	fmt.Println(fnk1(x5))

}

