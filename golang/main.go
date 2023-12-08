package main

import (
	"fmt"

	"isuct.ru/informatics2022/lab4"
)

func main() {
	fmt.Println("Голосова Анастасия Сергеевна")

	a := 1.35
	b := 0.98
	x_start := 1.14
	x_end := 4.24
	x_delta := 0.62

	lab4.Solution(a, b, x_start, x_end, x_delta)
}
