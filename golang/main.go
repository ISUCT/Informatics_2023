package main

import (
	"fmt"

	Lab4 "isuct.ru/informatics2022/internal"
)

func main() {
	r_a := Lab4.Part_1(2.0, 4.1, 0.77, 1.77, 0.2)
	var sliceb = []float64{1.24, 1.38, 2.38, 3.21, 0.68}
	r_b := Lab4.Part_2(sliceb, 2.0, 4.1)
	fmt.Println(r_a)
	fmt.Println(r_b)
	fmt.Println("Чудов Даниил")
	fmt.Println("17")
	fmt.Println("Не женат")
}
