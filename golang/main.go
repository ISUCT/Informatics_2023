package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Чурсина Софья 1/278")

	var a float64 = 4.1
	var b float64 = 2.7

	fmt.Println("Задача A")
	for x := 1.5; x <= 3.5; x += 0.4 {
		res := internal.Func_res(a, b, x)
		fmt.Println(res)
	}

	fmt.Println("Задача B")
	array := [5]float64{1.9, 2.15, 2.34, 2.74, 3.16}
	for _, v := range array {
		res := internal.Func_res(a, b, v)
		fmt.Println(res)
	}
}
