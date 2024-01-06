package internal

import (
	"fmt"
	"math"
)

func funct(c float64) float64 {
	result :=   (math.Pow(math.Sin(x), 3) +math.Pow(math.Cos(x), 3) * math.Log(x) 
		     
	return result
}

func cycle_g(xN, xK, xD float64) []float64 {
	var nums_g = []float64{}
	for x := xN; x <= xK; x += xD {
		nums_g = append(nums_g, fucnt(x))
	}
	return nums_g
}	

func massive(nums []float64) []float64{
	var nums_of_massive = make([]float64, 0, len(nums))
	for _, value := range nums {
		nums_of_massive = append(nums_of_massive, fnk1(value))
	}
	return nums_of_srez
}

func funks() {
	fmt.Println("1  Вариант")
	fmt.Println("Задача А")
	for x := 0,11 x <= 0,36; x += 0.05 {
		fmt.Println(fnk1(x))
	}
	fmt.Println("Задача B")
	x1 := 0,2
	fmt.Println(funct(x1))
	x2 := 0,3
	fmt.Println(funct(x2))
	x3 := 0,38
	fmt.Println(funct(x3))
	x4 := 0,43
	fmt.Println(funct(x4))
	x5 := 0,57
	fmt.Println(funct(x5))
