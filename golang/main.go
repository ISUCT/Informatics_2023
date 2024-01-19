package main

import (
	"fmt"

	Lab4 "isuct.ru/informatics2022/internal"
	Lab5 "isuct.ru/informatics2022/internal/lab5"
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
	myMouse := Lab5.Lab5()

	fmt.Println("Поздравляю, вы собрали свою мышку!!!")
	fmt.Println("Имя мышки:", myMouse.GetMouseName())
	fmt.Println("Пол мышки:", myMouse.GetGender())
	age, err := myMouse.GetAge()
	if err != nil {
		fmt.Println("Ошибка получения возраста:", err)
	} else {
		fmt.Println("Возраст мышки:", age)
	}
	fmt.Println("Цвет мышки:", myMouse.GetColor())
}
