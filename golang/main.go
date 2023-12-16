package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Чурсина Софья 1/278")
	fmt.Println(internal.Task_A(1.5, 3.5, 0.4))
	fmt.Println(internal.Task_B([]float64{1.9, 2.15, 2.34, 2.74, 3.16}))

	//Задание на структуры
	cat, _ := internal.NewCat(
		"Sobaka",
		"male",
		0,
	)
	if err := cat.SetAge(-5); err != nil {
		fmt.Println("Wrong age value")
	}
	fmt.Println(cat.GetAge())

	if err := cat.SetAge(5); err != nil {
		fmt.Println("Wrong age value")
	}
	fmt.Println(cat.GetAge())

	cat2, err := internal.NewCat(
		"Bobik",
		"male",
		3,
	)
	if err != nil {
		panic("Bad age")
	}
	fmt.Println(cat2)
	cat.PrintAllProperty()
}
