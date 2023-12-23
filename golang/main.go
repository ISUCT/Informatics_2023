package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/lab4"
	"isuct.ru/informatics2022/lab5"
)

func main() {
	fmt.Println("Голосова Анастасия Сергеевна")

	a := 1.35
	b := 0.98
	x_start := 1.14
	x_end := 4.24
	x_delta := 0.62

	lab4.Solution(a, b, x_start, x_end, x_delta)
	fmt.Printf("Structure\n")

	foxy, err := lab5.NewFox("Alisa", "female", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = foxy.SetAge(12)
	if err != nil {
		log.Fatal(err)
	}

	err = foxy.SetGender("female")
	if err != nil {
		log.Fatal(err)
	}

	image := foxy.GetView()

	fmt.Printf(image)
	foxy.DisplayFox()
}
