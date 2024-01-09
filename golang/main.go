package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab4"
	"isuct.ru/informatics2022/lab5"
)

func main() {
	fmt.Println("Спиридонова Полина Алексеевна")

	a := 7.2
	b := 1.3
	x_start := 1.56
	x_end := 4.71
	x_delta := 0.63

	fmt.Println("Решение задачи под буквой А")
	answers := lab4.TaskA(a, b, x_start, x_end, x_delta)
	lab4.Output(answers)

	fmt.Println("Решение задачи под буквой B")
	answers = lab4.TaskB(a, b, []float64{2.4, 2.8, 3.9, 4.7, 3.16})
	lab4.Output(answers)

	fmt.Printf("Структура\n")
	foxy, err := lab5.NewFox("Felix", 5, "red")
	if err != nil {
		log.Fatal(err)
	}
	err = foxy.SetAge(5)
	if err != nil {
		log.Fatal(err)
	}

	err = foxy.SetColor("red")
	if err != nil {
		log.Fatal(err)
	}
	image := foxy.GetView()
	fmt.Printf(image)
	foxy.DisplayFox()
}
