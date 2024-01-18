package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal"
	"isuct.ru/informatics2022/internal/lab5"
)

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const a float64 = 2.0
const b float64 = 0.95
const xStart float64 = 1.25
const xEnd float64 = 2.75
const xStep float64 = 0.3

func main() {
	fmt.Println("Быстров Владислав Евегньевич")

	fmt.Println("Задача А:")
	stepsResults := internal.CalculateFuncForSteps(a, b, xStart, xEnd, xStep)
	fmt.Println(stepsResults)

	fmt.Println("Задача Б:")
	xesResults := internal.CalculateFuncForXes(a, b, []float64{2.2, 3.78, 4.51, 6.58, 1.2})
	fmt.Println(xesResults)

	fmt.Println("Lab 5")

	dog, err := lab5.TheDog(16, "Oleg", "male")
	checkForError(err)

	fmt.Printf("Sobake %d let\n", dog.GetAge())
	fmt.Printf("Imya %s\n", dog.GetName())
	fmt.Printf("Sobaka - %s", dog.GetGender())
}
