package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/structures"
)

func main() {
	fmt.Println("Прибылов Алексей Александрович")
	fmt.Println("Задача A")
	x, y := lab4.TaskA(0.26, 0.66, 0.08)
	for i := range y {
		fmt.Println("При x =", x[i], "y =", y[i])
	}
	fmt.Println("Задача B")
	var list []float64 = []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	x, y = lab4.TaskB(list)
	for i := range y {
		fmt.Println("При x =", x[i], "y =", y[i])
	}
	//lab5
	char := structures.Char("Alduin", "Berserk", 20)
	fmt.Println(char)
	char.Lvl(30)
	char.Class("Hunt")
	fmt.Println(char)
}
