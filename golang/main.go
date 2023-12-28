package main

import (
	"fmt"
	"log"

	airplanestruct "isuct.ru/informatics2022/internal/lab5"
	//"isuct.ru/informatics2022/internal/lab4"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*fmt.Println("Кульбакин Аким Александрович")
	fmt.Println("Задача А")
	xn := 1.2
	xk := 2.2
	dx := 0.2
	a := 1.1
	b := 0.09
	x, y := lab4.TaksA(xn, xk, dx, a, b)
	for i := range y {
		fmt.Println("при x =", x[i], "y =", y[i])
	}
	fmt.Println("Задача B")
	xarray := []float64{1.21, 1.76, 2.53, 3.48, 4.52}
	x, y = lab4.TaskB(xarray, a, b)
	for i := range y {
		fmt.Println("при x =", x[i], "y =", y[i])*/

	//lab 5
	airplane, err := airplanestruct.NewAirplane(3000, "IL-62", "Russia")
	checkForError(err)
	aeroplane, err := airplanestruct.NewAirplane(853, "Some Airbus", "France")
	checkForError(err)
	lastairplane, err := airplanestruct.NewAirplane(968, "Boeing-747", "USA")
	checkForError(err)

	airplane.SetModel("Il'ushin IL-62")
	aeroplane.SetModel("Airbus A320")

	fmt.Println("Airplane model is", lastairplane.GetModel())
	fmt.Println("Airplane speed is", lastairplane.GetSpeed())
	fmt.Println("Airplane manufacturing country is", lastairplane.GetManufcountry())
}
