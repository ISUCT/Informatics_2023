package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal/catStruct"
	"isuct.ru/information2022/internal/lab4"
)

func printResult(xL []float64, yL [float64]) {
	for i := range yL {
		fnt.Printf("x = %.2f, y = %f\n", xL[i], yL[i])
	}
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Махотина Валерия Олеговна")
	fmt.Println("---")

	a1 := Task1(0.1, 0.5, 0.15, 1.37, 0.25)
	fmt.Print(a1)
	fmt.Println("---")

	a2 := Task2(0.1, 0.5, []float64{0.2, 0.3, 0.44, 0.6, 0.56})
	fmt.Println(a2) 

	var city = lab5.NewCity(400, "Russia")

	var err error = city.SetPopulation(400)
	checkForError(err)

	city.SetCountry("Russia")

	fmt.Printf("City's population is %d\n", city.GetPopulation())
	fmt.Printf("This city is in %s\n", city.GetCountry())
	fmt.GetTheCity()
}