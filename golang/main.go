package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)


func main() {
	//Лабораторная работа 2
	fmt.Println("Курушина Елизавета Андреевна")

	//Лабораторне работы 3 и 4
	fmt.Println("Значение функции №3 при данных задачи А =", internal.TaskA(0.2, 0.95, 0.15))
	fmt.Println("Значение функции №3 при данных задачи B =", internal.TaskB([]float64{0.15, 0.26, 0.37, 0.48, 0.56}))
}
