package main

import (
	"fmt"
	"log"
	"os"

	lab4 "isuct.ru/informatics2022/internal/lab4"
	"isuct.ru/informatics2022/internal/lab4/lab5"
)

func main() {
	myslicefortask2 := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	fmt.Println("Кузьмин Владислав Иванович")
	results := lab4.Primer1(0.26, 0.66, 0.08)
	fmt.Println(results)
	lab4.Primer2(myslicefortask2)
	fmt.Println(myslicefortask2)

	mouse, err := lab5.NewMouse(0.02, 1, 10.5)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Вес мышки:", mouse.GetWeight())
	fmt.Println("Возраст мышки:", mouse.GetAge())
	fmt.Println("Длина мышки:", mouse.GetLength())
}
