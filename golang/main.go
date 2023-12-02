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
	cat := internal.Cat{Age: 10, Sex: "male", Name: "Sobaka"}
	fmt.Println(cat.GetAge())
	cat.SetAge(5)
	fmt.Println(cat.GetAge())
	cat.PrintAllProperty()
}
