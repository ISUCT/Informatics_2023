package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab4"
	"isuct.ru/informatics2022/lab5"
)

func ErrorTest(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Голосова Анастасия Сергеевна")

	a := 1.35
	b := 0.98
	x_start := 1.14
	x_end := 4.24
	x_delta := 0.62

	lab4.Solution(a, b, x_start, x_end, x_delta)
	fmt.Printf("Structure\n")

	foxy, err_a, err_g := lab5.NewFox("Alisa", "female", 0)
	ErrorTest(err_a)
	ErrorTest(err_g)

	err_a = foxy.SetAge(12)
	ErrorTest(err_a)

	err_g = foxy.SetGender("female")
	ErrorTest(err_g)

	image := foxy.GetView()

	fmt.Printf(image)
	foxy.DisplayFox()
}
