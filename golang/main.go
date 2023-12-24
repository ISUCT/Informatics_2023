package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/lab5"
)

func main() {
	// Лабораторная 2
	fmt.Println("Байгулова Ева Валерьевна")

	// Лабораторная 3,4 - Вариант 3
	fmt.Println("Значение функции №3 при данных задачи А =", internal.TaskA(1.25, 2.75, 0.3))
	fmt.Println("Значение функции №3 при данных задачи B =", internal.TaskB([]float64{2.2, 3.78, 4.51, 6.58, 1.2}))

	// Лабораторная 5 - Вариант 3
	dog, err := lab5.NewDog("Hachikō", "Akita Inu", 11)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dog's name is", dog.GetName())
		fmt.Println("Dog's age is", dog.GetAge())
		fmt.Println("Dog's breed is", dog.GetBreed())
	}
}
