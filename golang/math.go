package internal

import (
	"fmt"
	"math"
)

func fnk1(c float64) float64 {
	result :=   (math.Pow(math.Sin(x), 3) +math.Pow(math.Cos(x), 3) * math.Log(x) 
		     
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
	for x := 0,11 x <= 0,36; x += 0.05 {
		fmt.Println(fnk1(x))
	}
	fmt.Println("Задача B")
	x1 := 0,2
	fmt.Println(fnk1(x1))
	x2 := 0,3
	fmt.Println(fnk1(x2))
	x3 := 0,38
	fmt.Println(fnk1(x3))
	x4 := 0,43
	fmt.Println(fnk1(x4))
	x5 := 0,57
	fmt.Println(fnk1(x5))



