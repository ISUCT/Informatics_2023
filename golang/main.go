package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab3"
)

func main() {
	fmt.Println("Кремнев Юрий Юрьевич")

	a := 1.6
	x_begin := 1.2
	x_end := 3.7
	x_delta := 0.5

	lab3.Solve(a,x_begin,x_end,x_delta)
}