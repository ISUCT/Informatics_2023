package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/airplanestruct"
	"isuct.ru/informatics2022/lab4"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//Лабораторная 2
	fmt.Println("Коротков Данила Максимович")
	//Лабораторная 3,4 (9 Вариант)
	a := 1.1
	b := 0.09
	x_begin := 1.2
	x_end := 2.2
	x_delta := 0.2
	lab4.Solve(a, b, x_begin, x_end, x_delta)
	//Лабораторная 5 (9 Вариант)
	airplane, err := airplanestruct.NewAirplane(3000, "Russian airplane", "Russia")
	checkForError(err)
	aeroplane, err := airplanestruct.NewAirplane(853, "Some Airbus", "France")
	checkForError(err)
	lastairplane, err := airplanestruct.NewAirplane(917, "Boeing-747", "USA")
	checkForError(err)

	airplane.SetModel("Il'ushin IL-62")
	aeroplane.SetModel("Airbus A320")

	fmt.Println("Airplane model is", lastairplane.GetModel())
	fmt.Println("Airplane speed is", lastairplane.GetSpeed())
	fmt.Println("Airplane manufacturing country is", lastairplane.GetManufcountry())
}
