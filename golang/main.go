package main

import (
	"fmt"
	"log"
	"isuct.ru/informatics2022/internal/lab4"
	structcat "isuct.ru/informatics2022/internal/structCat"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//Лабораторная работа 2
	fmt.Println("Chalunin Artem Romanovich 1/279")

	//Лабораторная работа номер 4 - варинант 23
	//Решение под А
	fmt.Println("Лабараторная работа номер 3, 4")
	xBegin := 1.5
	xEnd := 3.5
	xDelta := 0.4
	a := 4.1
	b := 2.7
	fmt.Println("Решения под A:")
	fmt.Println(lab4.Task_A(xBegin, xEnd, xDelta, a, b))

	//Решение под Б
	fmt.Println("Решения под B:")
	fmt.Println(lab4.Task_B([]float64{1.9, 2.15, 2.34, 2.74, 3.16}, a, b))

	//Лабораторная работа номер 5 - вариант 1(варианта 23 не было в списке, перепрыгнул на +1)
	fmt.Println("Лаборатноная работа номер 5")
	
	cat, err := structcat.NewCat(10, "Male", "British", "Skaya")
	checkError(err)
	fmt.Println(*cat)
	cat2, err := structcat.NewCat(18, "Female", "Tiger", "Melmon")
	fmt.Printf("Your cat is %d\n", cat.GetAge())
	fmt.Println(cat2)
	checkError(err)
	cat2.AssigneName("Funtik")
	fmt.Println(cat2)
}
