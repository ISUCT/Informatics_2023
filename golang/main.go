package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Терентьева Анастасия Романовна")

	fmt.Println(internal.Calc(1.23, 7.23, 1.2))
	fmt.Println(internal.Calcmassiw([]float64{1.88, 2.26, 3.84, 4.55, -6.21}))
}
