package main

import (
	"fmt"

	lab4 "isuct.ru/informatics2022/laba4"
	lab5 "isuct.ru/informatics2022/labb5"
)

func main() {
	mouse, err := lab5.NewMouse("Mickey", "Male", 92)
	if err != " " {
		fmt.Println(err)
	}
	mouse.MickeyM()

	mouse.MinnieM()

	fmt.Println("Солдатов Иван Сергеевич")
	fmt.Println("------------------------First_Exercise-------------------------------")
	a := 2.0
	x_end := 4.2
	x_step := 0.6
	x_start := 1.2

	lab4.Solve(a, x_start, x_end, x_step)
}
