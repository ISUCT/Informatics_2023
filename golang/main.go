package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Журавлев Сергей Викторович")

	a := -2.5
	b := 3.4
	x_begin := 3.5
	x_end := 6.5
	x_delta := 0.6

	lab4.Solve(a, b, x_begin, x_end, x_delta)
}
